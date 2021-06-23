package goami2

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"strings"
	"sync"
	"time"
)

// NetTOUT timeout for net connections and requests
var NetTOUT = time.Second * 3

// ErrorConnTOUT error on connection timeout
var ErrorConnTOUT = errorNew("Connection timeout")

// ErrorInvalidPrompt invalid prompt received from AMI server
var ErrorInvalidPrompt = errorNew("Invalid prompt on AMI server")

// ErrorNet networking errors
var ErrorNet = errorNew("Network error")

// ErrorLogin AMI server login failed
var ErrorLogin = errorNew("Login failed")

// Client structure
type Client struct {
	mu     sync.RWMutex
	reader *textproto.Reader
	writer *bufio.Writer
	conn   net.Conn
	cancel context.CancelFunc
	subs   *pubsub
	err    chan error
}

// NewClient connects to Asterisk using conn and tries to login using
// username and password. Creates new AMI client and returns client or error
func NewClient(conn net.Conn, username, password string) (*Client, error) {
	client, ctx := newClient(conn)
	if err := client.readPrompt(ctx, NetTOUT); err != nil {
		return nil, err
	}

	if err := client.login(ctx, NetTOUT, username, password); err != nil {
		return nil, err
	}

	client.initReader(ctx)

	return client, nil
}

// AllMessages	subscribes to any AMI message received from Asterisk server
// returns send-only channel or nil
func (c *Client) AllMessages() <-chan *Message {
	return c.subs.subscribe(keyAnyMsg)
}

// OnEvent subscribes by event by name (case insensative) and
// returns send-only channel or nil
func (c *Client) OnEvent(name string) <-chan *Message {
	return c.subs.subscribe(name)
}

// Action sends AMI message to Asterisk server and returns send-only
// response channel on nil
func (c *Client) Action(msg *Message) bool {
	if msg.ActionID() == "" {
		msg.AddActionID()
	}

	if err := c.write(msg.Bytes()); err != nil {
		c.emitError(err)
		return false
	}
	return true
}

// Close client and destroy all subscriptions to events and action responses
func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cancel()
	c.subs.destroy()
	c.conn.Close()
	close(c.err)
	c.err = nil
}

// Err retuns channel for error and signals that client should be probably restarted
func (c *Client) Err() <-chan error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.err
}

func newClient(conn net.Conn) (*Client, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		reader: textproto.NewReader(bufio.NewReader(conn)),
		writer: bufio.NewWriter(conn),
		conn:   conn,
		cancel: cancel,
	}
	return client, ctx
}

func (c *Client) readPrompt(parentCtx context.Context, tout time.Duration) error {
	ctx, cancel := context.WithTimeout(parentCtx, tout)
	defer cancel()

	reader := func() (chan string, chan error) {
		prompt := make(chan string)
		fail := make(chan error)
		go func() {
			defer close(prompt)
			defer close(fail)
			line, err := c.reader.ReadLine()
			if err != nil {
				fail <- err
				return
			}
			prompt <- line
		}()
		return prompt, fail
	}

	prompt, fail := reader()
	select {
	case <-ctx.Done():
		return ErrorConnTOUT
	case err := <-fail:
		return ErrorNet.msg(err.Error())
	case promptLine := <-prompt:
		if !strings.HasPrefix(promptLine, "Asterisk Call Manager") {
			return ErrorInvalidPrompt.msg(promptLine)
		}
	}
	return nil
}

func (c *Client) login(parentCtx context.Context, tout time.Duration, user, pass string) error {
	ctx, cancel := context.WithTimeout(parentCtx, tout)
	defer cancel()

	msg := loginMessage(user, pass)

	reader := func() (chan *Message, chan error) {
		chMsg := make(chan *Message)
		fail := make(chan error)
		go func() {
			defer close(chMsg)
			defer close(fail)
			msg, err := c.read()
			if err != nil {
				fail <- err
				return
			}
			chMsg <- msg
		}()
		return chMsg, fail
	}

	chMsg, fail := reader()
	if err := c.write(msg.Bytes()); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ErrorConnTOUT.msg("failed login")
	case err := <-fail:
		return ErrorNet.msg(err.Error())
	case msg := <-chMsg:
		if !msg.IsSuccess() {
			return ErrorLogin
		}
	}

	return nil
}

func (c *Client) initReader(ctx context.Context) {
	c.subs = newPubsub()
	c.err = make(chan error)

	go func() {
		defer c.subs.disable()
		for {
			select {
			case <-ctx.Done():
				c.emitError(ctx.Err())
				return
			default:
				msg, err := c.read()
				if ctx.Err() != nil {
					c.emitError(ctx.Err())
					return // already closed
				}
				if err != nil {
					if err == ErrorNet { // only network errors break the loop
						c.emitError(err)
						return
					}
				}
				c.publish(msg)
			}
		}
	}()
}

func (c *Client) publish(msg *Message) {
	if msg != nil {
		c.subs.publish(msg)
	}
}

func (c *Client) emitError(err error) {
	go func(err error) {
		c.mu.Lock()
		defer c.mu.Unlock()
		if err == nil || c.err == nil {
			return
		}
		// c.mu.Lock()
		// defer c.mu.Unlock()
		c.err <- err
	}(err)
}

func (c *Client) write(msg []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, err := c.writer.Write(msg); err != nil {
		return err
	}
	if err := c.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (c *Client) read() (*Message, error) {
	headers, err := c.reader.ReadMIMEHeader()
	if err != nil {
		if err.Error() == "EOF" || err.Error() == "io: read/write on closed pipe" {
			return nil, ErrorNet
		}
		return nil, err
	}
	msg := newMessage(headers)
	return msg, nil
}
