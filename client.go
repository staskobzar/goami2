package goami2

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"
)

type eventKey string

const (
	anyEventKey  = "any-event-c27ffb3c3a56490"
	loginTimeout = 3 * time.Second
)

// Client structure
type Client struct {
	conn   net.Conn
	err    chan error
	cancel context.CancelFunc
	pool   *pool
}

// NewClient creates new AMI client and returns client or error
func NewClient(conn net.Conn) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:   conn,
		err:    make(chan error, 1),
		cancel: cancel,
		pool:   newPool(),
	}

	ch, err := connPipeline(ctx, conn, client.err)
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
				client.pool.dispatch(msg)
			}
		}
	}()
	return client, nil
}

// Close client and stop all routines
func (c *Client) Close() {
	c.cancel()
	close(c.err)
	c.conn.Close()
}

// Error returns channel with connection errors
func (c *Client) Error() <-chan error {
	return c.err
}

// AnyEvent provides channel for any AMI events received
func (c *Client) AnyEvent() (chan *AMIMsg, error) {
	ch, err := c.pool.add(anyEventKey)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// OnEvent provides channel for specific event.
// Returns error if listener for event already created.
func (c *Client) OnEvent(event string) (chan *AMIMsg, error) {
	event = strings.ToLower(event)

	ch, err := c.pool.add(event)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// Action send to Asterisk MI.
// ActionID will be generated.
func (c *Client) Action(actionName string, fields map[string]string) (chan *AMIMsg, error) {
	action := NewAction()
	action.New(actionName)
	actionID := action.ActionID()

	if actionID == "" {
		return nil, errors.New("failed to set response channel for empty ActionID")
	}
	ch, err := c.pool.add(actionID)
	if err != nil {
		return nil, err
	}
	for header, val := range fields {
		action.Field(header, val)
	}
	err = c.send(action.Message())
	if err != nil {
		c.err <- err
		return nil, err
	}
	return ch, nil
}

// Login action. Blocking and waits response.
func (c *Client) Login(user, pass string) error {
	action := NewAction()
	login := action.Login(user, pass)
	ch, err := c.pool.add(action.ActionID())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), loginTimeout)
	defer cancel()
	if err := c.send(login); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return errors.New("login timeout")
	case resp := <-ch:
		if resp.IsSuccess() {
			return nil
		}
		return errors.New(resp.Message())
	}
}

func (c *Client) send(action []byte) error {
	_, err := c.conn.Write(action)
	if err != nil {
		return err
	}
	return nil
}

func connPipeline(ctx context.Context, conn net.Conn, cherr chan<- error) (<-chan *AMIMsg, error) {
	var err error

	err = readPrompt(ctx, conn)
	if err != nil {
		return nil, err
	}

	chanOut, err := newReader(ctx, conn, cherr)
	if err != nil {
		return nil, err
	}
	return chanOut, nil
}

// pool of events and responses with dispatcher
type pool struct {
	m *sync.Map
}

type message struct {
	ch     chan *AMIMsg
	isList bool
}

func newPool() *pool {
	p := &pool{m: &sync.Map{}}
	return p
}

func (p *pool) dispatch(msg *AMIMsg) {
	if msg == nil {
		return
	}
	// dispatch message to Action listener
	go func(m *AMIMsg) {
		if actionID, exists := m.ActionID(); exists {
			p.dispatchAction(actionID, m)
		}
	}(msg)

	// dispatch to Event listeners
	go func(m *AMIMsg) {
		p.dispatchEvent(m)
	}(msg)
}

func (p *pool) dispatchAction(actionID string, msg *AMIMsg) {
	resp := p.get(actionID)
	if resp == nil {
		return
	}

	resp.ch <- msg // dispatch message

	if msg.IsEventListStart() {
		resp.isList = true
		p.update(actionID, resp)
	}

	if !resp.isList || msg.IsEventListEnd() {
		close(resp.ch)
		p.delete(actionID)
	}
}

func (p *pool) dispatchEvent(msg *AMIMsg) {
	eventName, ok := msg.Event()
	if !ok {
		return
	}

	if event := p.get(strings.ToLower(eventName)); event != nil {
		event.ch <- msg // dispatch event to listener
	}

	if anyEvent := p.get(anyEventKey); anyEvent != nil {
		anyEvent.ch <- msg // dispatch event to AnyEvent listener
	}
}

func (p *pool) add(actionID string) (chan *AMIMsg, error) {
	ch := make(chan *AMIMsg)
	s := &message{
		ch:     ch,
		isList: false,
	}
	if _, ok := p.m.LoadOrStore(eventKey(actionID), s); ok {
		return nil, errors.New("Action ID is already in pool: " + actionID)
	}
	return ch, nil
}

func (p *pool) get(key string) *message {
	val, exists := p.m.Load(eventKey(key))
	if !exists {
		return nil
	}
	return val.(*message)
}

func (p *pool) update(actionID string, msg *message) {
	p.delete(actionID)
	p.m.Store(eventKey(actionID), msg)
}

func (p *pool) delete(actionID string) {
	p.m.Delete(eventKey(actionID))
}
