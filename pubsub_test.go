package goami2

import (
	"fmt"
	"net/textproto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test helper to generate dummy AMI event as Message
func dummyEvent(name string) *Message {
	m := newMessage(textproto.MIMEHeader{"Event": []string{name}})
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")

	return m
}

// test helper to generate dummy AMI response as Message
func dummyResponse(state, msg string) *Message {
	if msg == "" {
		msg = "Dummy response message"
	}
	m := newMessage(textproto.MIMEHeader{"Response": []string{state}})
	m.AddField("Message", msg)

	return m
}

// test helper to avoid long blocking channels
func recvOrTimeout(ch pubChan) *Message {
	select {
	case m := <-ch:
		return m
	case <-time.After(100 * time.Millisecond):
		fmt.Println("===> recvOrTimeout: Timeout receiving from chan")
		return nil
	}
}

func TestPubsub_newPubsub(t *testing.T) {
	ps := newPubsub()
	assert.IsType(t, ps, &pubsub{})
	assert.False(t, ps.off)
	assert.Empty(t, ps.msg)
}

func TestPubsub_subAnyMessage(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()

	ch := ps.subscribe(keyAnyMsg)
	assert.Equal(t, 1, ps.lenMsgChans())
	chMatch, ok := ps.msg[keyAnyMsg]
	assert.True(t, ok)
	assert.Equal(t, chMatch, ch)
}

func TestPubsub_subEvent(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()

	ch := ps.subscribe("AgentLogoff")
	assert.Equal(t, 1, ps.lenMsgChans())
	chMatch, ok := ps.msg["agentlogoff"]
	assert.True(t, ok)
	assert.Equal(t, chMatch, ch)
}

func TestPubsub_subEvent_disabled(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()
	ps.disable()

	ch1 := ps.subscribe(keyAnyMsg)
	assert.Nil(t, ch1)
	ch2 := ps.subscribe("AgentLogoff")
	assert.Nil(t, ch2)
}

func TestPubsub_publish_anyEvent(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()
	ch := ps.subscribe(keyAnyMsg)

	m := dummyEvent("Newchannel")
	ps.publish(m)

	msg := recvOrTimeout(ch)
	assert.Equal(t, m, msg)
}

func TestPubsub_publish_any_message(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()
	ch := ps.subscribe(keyAnyMsg)

	ps.publish(dummyResponse("Success", ""))

	msg := recvOrTimeout(ch)
	assert.True(t, msg.IsResponse())
	assert.Equal(t, "Success", msg.Field("Response"))

	ps.publish(dummyResponse("Failed", ""))
	msg = recvOrTimeout(ch)
	assert.True(t, msg.IsResponse())
	assert.Equal(t, "Failed", msg.Field("Response"))

	ps.publish(dummyEvent("Newchannel"))
	msg = recvOrTimeout(ch)
	assert.False(t, msg.IsResponse())
	assert.Equal(t, "Newchannel", msg.Field("Event"))
}

func TestPubsub_publish_event(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()
	ch := ps.subscribe("NewExten")

	ps.publish(dummyEvent("Newexten"))

	msg := recvOrTimeout(ch)
	assert.Equal(t, "Newexten", msg.Field("event"))
}

func TestPubsub_publish_event_filtered(t *testing.T) {
	ps := newPubsub()
	defer ps.destroy()
	ch := ps.subscribe("NewExten")

	ps.publish(dummyEvent("NewCallerid"))
	ps.publish(dummyEvent("NewConnectedLine"))
	ps.publish(dummyResponse("Success", ""))
	ps.publish(dummyEvent("Newexten"))

	msg := recvOrTimeout(ch)
	assert.Equal(t, "Newexten", msg.Field("event"))
}

func TestPubsub_publish_on_disabled(t *testing.T) {
	ps := newPubsub()
	ps.disable()
	assert.False(t, ps.publish(dummyEvent("NewChannel")))
}

func TestPubsub_destroy(t *testing.T) {
	ps := newPubsub()
	ps.subscribe("NewExten")
	ps.subscribe("QueueMember")
	ps.subscribe("NewChannel")

	assert.Equal(t, 3, ps.lenMsgChans())

	ps.subscribe("NewChannel")
	assert.Equal(t, 3, ps.lenMsgChans())

	assert.False(t, ps.off)

	ps.destroy()
	assert.True(t, ps.off)
	assert.Equal(t, 0, len(ps.msg))
}

func TestPubsub_publish_subscribe_after_destroy(t *testing.T) {
	ps := newPubsub()
	ch := ps.subscribe(keyAnyMsg)
	ps.destroy()

	assert.Nil(t, ps.subscribe("foo"))
	assert.False(t, ps.publish(dummyEvent("QueueParams")))
	_, ok := <-ch
	assert.False(t, ok)
}
