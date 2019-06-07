package goami2

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
)

var anyEventKey string = "any-event"

type pool struct {
	actionIDConsumer map[string]chan *AMIMsg
	eventConsumer    map[string]chan *AMIMsg
}

func (p *pool) responseDispatch(msg *AMIMsg) bool {
	actionID, ok := msg.ActionId()
	if !ok {
		return false
	}
	respChan, ok := p.actionIDConsumer[actionID]
	if ok {
		respChan <- msg
		return true
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
	if _, ok := p.actionIDConsumer[actionId]; ok {
		return nil
	}
	chMsg := make(chan *AMIMsg)
	p.actionIDConsumer[actionId] = chMsg
	return chMsg
}

type Client struct {
	p      *pool
	cancel context.CancelFunc
	m      sync.RWMutex
	action *Action
	conn   net.Conn
}

// NewClient creates new AMI client and returns client or error
func NewClient(conn net.Conn) (*Client, error) {
	p := &pool{
		make(map[string]chan *AMIMsg),
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

func (c *Client) Action(action string, fields map[string]string) (chan *AMIMsg, error) {
	c.action.New(action)
	actionId := c.action.ActionId()

	if actionId == "" {
		return nil, errors.New("failed to set response channel for empty ActionID")
	}
	ch := c.p.addAction(actionId)
	if ch == nil {
		return nil, errors.New("response channel for ActionID '" + actionId + "' already exists")
	}
	for header, val := range fields {
		c.action.Field(header, val)
	}
	c.conn.Write(c.action.Message())
	return ch, nil
}

func (c *Client) Login(user, pass string) (chan *AMIMsg, error) {
	f := make(map[string]string)
	f["Username"] = user
	f["Secret"] = pass
	return c.Action("Login", f)
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
