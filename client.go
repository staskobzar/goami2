package goami2

import (
	"context"
	"net"
	"time"
)

type eventKey string

const (
	anyEventKey  = eventKey("any-event-c27ffb3c3a56490")
	loginTimeout = 3 * time.Second
)

// Client structure
type Client struct {
	conn   net.Conn
	cancel context.CancelFunc
}

// NewClient creates new AMI client and returns client or error
func NewClient(conn net.Conn) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:   conn,
		cancel: cancel,
	}

	ch, err := connPipeline(ctx, conn)
	if err != nil {
		return nil, err
	}

	go func() {
		defer conn.Close()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-ch:
				_ = msg
				// dispatch message
			}
		}
	}()
	return client, nil
}

// Close client and stop all routines
func (c *Client) Close() {
	c.cancel()
}

// AnyEvent provides channel for any AMI events received
func (c *Client) AnyEvent() (chan *AMIMsg, error) {
	chMsg := make(chan *AMIMsg)
	return chMsg, nil
}

// OnEvent provides channel for specific event.
// Returns error if listener for event already created.
func (c *Client) OnEvent(event string) (chan *AMIMsg, error) {
	chMsg := make(chan *AMIMsg)
	return chMsg, nil
}

// Action send to Asterisk MI.
func (c *Client) Action(action string, fields map[string]string) (chan *AMIMsg, error) {
	return nil, nil
}

// Login action. Blocking and waits response.
func (c *Client) Login(user, pass string) error {
	return nil
}
