package goami2

import (
	"strings"
	"sync"
)

const (
	keyAnyEvent = "anyevent"
)

type pubChan chan Message
type subChans map[string][]pubChan

type pubsub struct {
	mu   sync.RWMutex
	subs subChans
	off  bool
}

func newPubsub() *pubsub {
	p := &pubsub{}
	p.subs = make(subChans)
	return p
}

func (ps *pubsub) disable() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.off = true
}

func (ps *pubsub) subAnyEvent() pubChan {
	return ps.subscribe(keyAnyEvent)
}

// subscribe to event by name or by action id as key
func (ps *pubsub) subByKey(key string) pubChan {
	return ps.subscribe(strings.ToLower(key))
}

func (ps *pubsub) destroy() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.off = true
	for key, subs := range ps.subs {
		for _, ch := range subs {
			close(ch)
		}
		delete(ps.subs, key)
	}
}

func (ps *pubsub) subscribe(key string) pubChan {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.off {
		return nil
	}

	ch := make(pubChan)
	ps.subs[key] = append(ps.subs[key], ch)

	return ch
}

func (ps *pubsub) publish(msg Message) bool {
	if ps.off {
		return false
	}

	if msg.IsEvent() {
		ps.pubAnyEvent(msg)
		ps.pubEvent(msg)
	}
	ps.pubByActionID(msg)
	return true
}

func (ps *pubsub) pubEvent(msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	key := strings.ToLower(msg.Field("event"))
	if subs, ok := ps.subs[key]; ok {
		pubToChannels(subs, msg)
	}
}

func (ps *pubsub) pubAnyEvent(msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if subs, ok := ps.subs[keyAnyEvent]; ok {
		pubToChannels(subs, msg)
	}
}

// TODO: delete subscription after publish
// TODO: lists sending control
func (ps *pubsub) pubByActionID(msg Message) {
	actionId := msg.ActionID()
	if actionId == "" {
		return
	}
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if subs, ok := ps.subs[actionId]; ok {
		pubToChannels(subs, msg)
	}
}

func pubToChannels(subs []pubChan, msg Message) {
	for _, ch := range subs {
		go func(sub pubChan, m Message) {
			sub <- m
		}(ch, msg)
	}
}
