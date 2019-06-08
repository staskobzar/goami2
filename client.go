package goami2

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"
)

var anyEventKey string = "any-event"

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
}

func (p *pool) responseDispatch(msg *AMIMsg) bool {
	actionID, ok := msg.ActionId()
	if !ok {
		return false
	}
	if respChan, ok := p.actionConsumer[actionID]; ok {
		respChan.ch <- msg
		if msg.IsEventListStart() {
			respChan.isList = true
		}
		if !respChan.isList || (respChan.isList && msg.IsEventListEnd()) {
			close(respChan.ch)
			delete(p.actionConsumer, actionID)
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
	if ch, ok := p.eventConsumer[event]; ok {
		ch <- msg
	}
	if ch, ok := p.eventConsumer[anyEventKey]; ok {
		ch <- msg
	}
}

func (p *pool) addAction(actionId string) chan *AMIMsg {
	if _, ok := p.actionConsumer[actionId]; ok {
		return nil
	}
	chMsg := make(chan *AMIMsg)
	p.actionConsumer[actionId] = newActionDispatchConsumer(chMsg)
	return chMsg
}

type Client struct {
	p      *pool
	cancel context.CancelFunc
	m      sync.Mutex
	action *Action
	conn   net.Conn
}

// NewClient creates new AMI client and returns client or error
func NewClient(conn net.Conn) (*Client, error) {
	p := &pool{
		make(map[string]*actionDispatchConsumer),
		make(map[string]chan *AMIMsg),
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

func (c *Client) Close() {
	c.cancel()
}

func (c *Client) AnyEvent() (chan *AMIMsg, error) {
	if _, ok := c.p.eventConsumer[anyEventKey]; ok {
		return nil, errors.New("any event consumer channel already exists")
	}
	chMsg := make(chan *AMIMsg)
	c.p.eventConsumer[anyEventKey] = chMsg
	return chMsg, nil
}

func (c *Client) OnEvent(event string) (chan *AMIMsg, error) {
	event = strings.ToLower(strings.TrimSpace(event))
	if _, ok := c.p.eventConsumer[event]; ok {
		return nil, errors.New("event '" + event + "' consumer channel already exists")
	}
	chMsg := make(chan *AMIMsg)
	c.p.eventConsumer[event] = chMsg
	return chMsg, nil
}

func (c *Client) Action(action string, fields map[string]string) (chan *AMIMsg, error) {
	c.action.New(action)
	actionID := c.action.ActionId()

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

func (c *Client) Login(user, pass string) error {
	login := c.action.Login(user, pass)
	ch := c.p.addAction(c.action.ActionId())
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
		c.m.Lock()
		defer c.m.Unlock()
		if c.p.responseDispatch(msg) {
			return
		}
		c.p.eventDispatch(msg)
	}()
}
