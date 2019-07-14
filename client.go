package goami2

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"
)

var anyEventKey = "any-event-c27ffb3c3a56490"

type actionDispatchConsumer struct {
	ch     chan *AMIMsg
	isList bool
}

func newActionDispatchConsumer(ch chan *AMIMsg) *actionDispatchConsumer {
	return &actionDispatchConsumer{
		ch:     ch,
		isList: false,
	}
}

type pool struct {
	actionConsumer map[string]*actionDispatchConsumer
	eventConsumer  map[string]chan *AMIMsg
	mux            sync.Mutex
}

func (p *pool) responseDispatch(msg *AMIMsg) bool {
	p.mux.Lock()
	defer p.mux.Unlock()
	actionID, ok := msg.ActionID()
	if !ok {
		return false
	}
	if respChan, ok := p.matchAction(actionID); ok {
		respChan.ch <- msg
		if msg.IsEventListStart() {
			respChan.isList = true
		}
		if !respChan.isList || (respChan.isList && msg.IsEventListEnd()) {
			close(respChan.ch)
			p.delAction(actionID)
			return true
		}
	}
	return false
}

func (p *pool) eventDispatch(msg *AMIMsg) {
	if !msg.IsEvent() {
		return
	}
	event, ok := msg.Event()
	if !ok {
		return
	}
	event = strings.ToLower(strings.TrimSpace(event))
	if ch, ok := p.matchEvent(event); ok {
		ch <- msg
	}
	if ch, ok := p.matchEvent(anyEventKey); ok {
		ch <- msg
	}
}

func (p *pool) matchEvent(event string) (chan<- *AMIMsg, bool) {
	ch, ok := p.eventConsumer[event]
	return ch, ok
}

func (p *pool) setEventConsumer(event string, chMsg chan *AMIMsg) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.eventConsumer[event] = chMsg
}

func (p *pool) matchAction(action string) (*actionDispatchConsumer, bool) {
	ch, ok := p.actionConsumer[action]
	return ch, ok
}

func (p *pool) addAction(actionID string) chan *AMIMsg {
	p.mux.Lock()
	defer p.mux.Unlock()
	if _, ok := p.matchAction(actionID); ok {
		return nil
	}
	chMsg := make(chan *AMIMsg)
	p.actionConsumer[actionID] = newActionDispatchConsumer(chMsg)
	return chMsg
}

func (p *pool) delAction(actionID string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	delete(p.actionConsumer, actionID)
}

// Client structure
type Client struct {
	p      *pool
	cancel context.CancelFunc
	action *Action
	conn   net.Conn
}

// NewClient creates new AMI client and returns client or error
func NewClient(conn net.Conn) (*Client, error) {
	p := &pool{
		actionConsumer: make(map[string]*actionDispatchConsumer),
		eventConsumer:  make(map[string]chan *AMIMsg),
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{p: p, cancel: cancel, action: NewAction(), conn: conn}
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
			case message := <-ch:
				c.dispatch(message)
			}
		}
	}()

	return c, nil
}

// Close client and stop all routines
func (c *Client) Close() {
	c.cancel()
}

// AnyEvent provides channel for any AMI events received
func (c *Client) AnyEvent() (chan *AMIMsg, error) {
	if _, ok := c.p.matchEvent(anyEventKey); ok {
		return nil, errors.New("any event consumer channel already exists")
	}
	chMsg := make(chan *AMIMsg)
	c.p.setEventConsumer(anyEventKey, chMsg)
	return chMsg, nil
}

// OnEvent provides channel for specific event.
// Returns error if listener for event already created.
func (c *Client) OnEvent(event string) (chan *AMIMsg, error) {
	event = strings.ToLower(strings.TrimSpace(event))
	if _, ok := c.p.matchEvent(event); ok {
		return nil, errors.New("event '" + event + "' consumer channel already exists")
	}
	chMsg := make(chan *AMIMsg)
	c.p.setEventConsumer(event, chMsg)
	return chMsg, nil
}

// Action send to Asterisk MI.
func (c *Client) Action(action string, fields map[string]string) (chan *AMIMsg, error) {
	c.action.New(action)
	actionID := c.action.ActionID()

	if actionID == "" {
		return nil, errors.New("failed to set response channel for empty ActionID")
	}
	ch := c.p.addAction(actionID)
	if ch == nil {
		return nil, errors.New("response channel for ActionID '" + actionID + "' already exists")
	}
	for header, val := range fields {
		c.action.Field(header, val)
	}
	c.conn.Write(c.action.Message())
	return ch, nil
}

// Login action. Blocking and waits response.
func (c *Client) Login(user, pass string) error {
	login := c.action.Login(user, pass)
	ch := c.p.addAction(c.action.ActionID())
	if ch == nil {
		return errors.New("response channel for ActionID already exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c.conn.Write(login)
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

func (c *Client) dispatch(msg *AMIMsg) {
	go func() {
		if c.p.responseDispatch(msg) {
			return
		}
		c.p.eventDispatch(msg)
	}()
}
