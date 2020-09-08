package goami2

import (
	"strings"
	"sync"
)

const (
	keyAnyMsg = "anymessage"
)

type pubChan chan *Message

type msgChans map[string]pubChan

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
	if ps.off {
		return
	}
	ps.off = true
	for key, ch := range ps.msg {
		close(ch)
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
	if _, ok := ps.msg[key]; !ok {
		ps.msg[key] = ch
	}

	return ch
}

func (ps *pubsub) publish(msg *Message) bool {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if ps.off {
		return false
	}

	if ch, ok := ps.msg[keyAnyMsg]; ok {
		go func(ch pubChan) {
			ch <- msg
		}(ch)
	}

	if name := strings.ToLower(msg.Field("event")); name != "" {
		if ch, ok := ps.msg[name]; ok {
			go func(ch pubChan) {
				ch <- msg
			}(ch)
		}
	}

	return true
}
