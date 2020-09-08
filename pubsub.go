package goami2

import (
	"strings"
	"sync"
)

const (
	keyAnyMsg = "anymessage"
)

type pubChan chan *Message

type msgChans map[string][]pubChan

type pubsub struct {
	mu  sync.RWMutex
	msg msgChans
	off bool
}

func newPubsub() *pubsub {
	p := &pubsub{}
	p.msg = make(msgChans)
	return p
}

func (ps *pubsub) disable() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.off = true
}

func (ps *pubsub) lenMsgChans() int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return len(ps.msg)
}

func (ps *pubsub) destroy() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.off = true
	for key, subs := range ps.msg {
		for _, ch := range subs {
			close(ch)
		}
		delete(ps.msg, key)
	}
}

// subscribe to event by name or by action id as key
func (ps *pubsub) subscribe(key string) pubChan {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.off {
		return nil
	}
	key = strings.ToLower(key)

	ch := make(pubChan)
	ps.msg[key] = append(ps.msg[key], ch)

	return ch
}

func (ps *pubsub) publish(msg *Message) bool {
	if ps.off {
		return false
	}

	ps.pubAnyMsg(msg)

	if msg.IsEvent() {
		name := strings.ToLower(msg.Field("event"))
		ps.pubEvent(name, msg)
	}
	return true
}

func (ps *pubsub) pubEvent(name string, msg *Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if subs, ok := ps.msg[name]; ok {
		pubToChannels(subs, msg)
	}
}

func (ps *pubsub) pubAnyMsg(msg *Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if subs, ok := ps.msg[keyAnyMsg]; ok {
		pubToChannels(subs, msg)
	}
}

func pubToChannels(subs []pubChan, msg *Message) {
	for _, ch := range subs {
		go func(sub pubChan, m *Message) {
			sub <- m
		}(ch, msg)
	}
}
