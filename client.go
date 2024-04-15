package goami2

import (
	"bufio"
	"context"
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
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

// MustSend sends a message to nework and returns
// network errors if any write away. May block
// until network timeout
func (c *Client) MustSend(msg []byte) error {
	if c.conn == nil {
		return fmt.Errorf("%w: closed connection: failed to send message", ErrConn)
	}
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
	chPack, errConn := consume(c.conn)
	defer c.Close()
	for {
		select {
		case pack := <-chPack:
			msg, err := Parse(pack)
			if err != nil {
				c.emitErr(err)
				continue
			}
			c.emitMsg(msg)
		case <-ctx.Done():
			c.emitErr(ErrEOF)
			return
		case err := <-errConn:
			c.emitErr(err)
			return
		}
	}
}

// comsume all AMI data from network and split by AMI terminating \r\n\r\n.
// When found send to main loop to parse or send error and stop on network close
func consume(conn net.Conn) (chan string, chan error) {
	_ = conn.SetReadDeadline(time.Time{}) // assure no dealine for reading
	pack, chErr := make(chan string), make(chan error)
	go func(chPack chan string, chErr chan error, conn net.Conn) {
		defer close(chPack)
		defer close(chErr)
		reader := bufio.NewReader(conn)
		buf := &strings.Builder{}
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				chErr <- fmt.Errorf("%w: failed read: %s", ErrEOF, err)
				return
			}
			_, _ = buf.WriteString(line)
			if line == "\r\n" { // end of packet
				chPack <- buf.String()
				buf.Reset()
			}
		}
	}(pack, chErr, conn)
	return pack, chErr
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
	buf := make([]byte, 1024) // long enough for prompt
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
	n, err = c.conn.Read(buf)
	if err != nil {
		return fmt.Errorf("%w: failed to read login response: %s", ErrAMI, err)
	}

	msg, err := Parse(string(buf[:n]))
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
