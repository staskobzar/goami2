package goami2

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Client is a AMI connection management object
type Client struct {
	mu      sync.Mutex
	conn    net.Conn
	recv    chan *Message
	err     chan error
	timeout time.Duration // connection read/write timeout
}

// Action sends AMI action to an Asterisk server
// Returns true on success and false if fails
// This function is deprecated and will be removed
// Use Send or MustSend instead
func (c *Client) Action(action *Message) bool {
	if err := c.MustSend(action.Byte()); err != nil {
		return false
	}
	return true
}

// AllMessages returns a channel that receives any AMI messages
// from the client connection
func (c *Client) AllMessages() <-chan *Message {
	return c.recv
}

// Close client
func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		_ = c.conn.Close()
		c.conn = nil
	}
	if !isClosedChan(c.recv) {
		close(c.recv)
		c.recv = nil
	}
	if !isClosedChan(c.err) {
		close(c.err)
		c.err = nil
	}
}

// Err returns channel of errors of the client
func (c *Client) Err() <-chan error {
	return c.err
}

// MustSend sends a message to nework and returns
// network errors if any write away. May block
// until network timeout
func (c *Client) MustSend(msg []byte) error {
	if err := c.setWTimeout(); err != nil {
		return fmt.Errorf("%w: failed to set net timeout: %q", ErrConn, err)
	}
	if _, err := c.conn.Write(msg); err != nil {
		return fmt.Errorf("%w: failed send message: %q", ErrConn, err)
	}
	return nil
}

// Send AMI message as a bytes array
// None-blocking method. Any network errors
// will be send back via Client.Err() channel
func (c *Client) Send(msg []byte) {
	go func(message []byte) {
		_ = c.MustSend(message)
	}(msg)
}

// creates Client with default values
func makeClient(conn net.Conn) *Client {
	return &Client{
		conn:    conn,
		recv:    make(chan *Message, 12),
		err:     make(chan error, 1),
		timeout: netTimeout,
	}
}

// main consumer loop that reads from connection
func (c *Client) loop(ctx context.Context) {
	buf := make([]byte, bufSize)
	_ = c.conn.SetReadDeadline(time.Time{})
	for {
		select {
		case <-ctx.Done():
			c.emitErr(ErrEOF)
			return
		default:
			msg, err := c.read(buf)
			if err != nil {
				if errors.Is(err, ErrConn) {
					c.emitErr(fmt.Errorf("%w: conn read error: %s", ErrEOF, err))
					return
				}
				c.emitErr(err)
				continue
			}

			c.emitMsg(msg)
		}
	}
}

func (c *Client) read(buf []byte) (*Message, error) {
	packet, err := readPack(c.conn, buf)
	if err != nil {
		return nil, fmt.Errorf("%w: failed read: %s", ErrConn, err)
	}
	msg, err := Parse(packet)
	if err != nil {
		return nil, fmt.Errorf("%w: failed parse AMI message: %s", ErrAMI, err)
	}
	return msg, nil
}

func readPack(conn net.Conn, buf []byte) (string, error) {
	var packet string
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("%w: failed read: %s", ErrConn, err)
		}
		packet += string(buf[:n])
		if len(packet) > len(buf) {
			return "", fmt.Errorf("%w: too long input: %d", ErrAMI, len(packet))
		}

		if strings.HasSuffix(packet, "\r\n\r\n") {
			return packet, nil
		}
	}
}

func (c *Client) emitErr(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case c.err <- err:
	case <-time.After(chanGiveup):
		// failed to send and exit here to avoid blocking
	}
}

func (c *Client) emitMsg(msg *Message) {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case c.recv <- msg:
	case <-time.After(chanGiveup):
		// failed to send and exit here to avoid blocking
	}
}

func (c *Client) login(username, password string) error {
	// make sure connection is not blocking
	if err := c.setRTimeout(); err != nil {
		return fmt.Errorf("%w: failed setup read timeout for login: %s", ErrConn, err)
	}

	if err := c.setWTimeout(); err != nil {
		return fmt.Errorf("%w: failed setup read timeout for login: %s", ErrConn, err)
	}

	// read prompt
	buf := make([]byte, bufSize)
	n, err := c.conn.Read(buf)
	if err != nil {
		return fmt.Errorf("%w: failed read prompt: %s", ErrConn, err)
	}
	if promptPrefix != string(buf[:len(promptPrefix)]) {
		return fmt.Errorf("%w: unexpected prompt: %q", ErrAMI, buf[:n])
	}

	// send login
	login := NewAction("Login")
	login.AddField("Username", username)
	login.AddField("Secret", password)
	if _, err := c.conn.Write(login.Byte()); err != nil {
		return fmt.Errorf("%w: failed write login: %q", ErrConn, err)
	}

	// read login response
	msg, err := c.read(buf)
	if err != nil {
		return fmt.Errorf("%w: failed to read login response: %s", ErrAMI, err)
	}

	if !msg.IsSuccess() {
		return fmt.Errorf("%w: failed login: %q", ErrAMI, msg.Field("Message"))
	}

	return nil
}

func (c *Client) setRTimeout() error {
	return c.conn.SetReadDeadline(time.Now().Add(c.timeout))
}

func (c *Client) setWTimeout() error {
	return c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
}

func isClosedChan[T any](c <-chan T) bool {
	if c == nil {
		return true
	}
	select {
	case <-c:
		return true
	default:
		return false
	}
}
