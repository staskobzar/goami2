package goami2

import (
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dummyEvent(name string) Message {
	m := newMessage(textproto.MIMEHeader{"Event": []string{name}})
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")

	return m
}

func dummyResponse(state, msg string) Message {
	if msg == "" {
		msg = "Dummy response message"
	}
	m := newMessage(textproto.MIMEHeader{"Response": []string{state}})
	m.AddField("Message", msg)

	return m
}

func TestPubsub_newPubsub(t *testing.T) {
	ps := newPubsub()
	assert.IsType(t, ps, &pubsub{})
	assert.False(t, ps.off)
	assert.Empty(t, ps.subs)
}

func TestPubsub_subAnyEvent(t *testing.T) {
	ps := newPubsub()
	ch := ps.subAnyEvent()
	assert.Equal(t, 1, len(ps.subs))
	assert.Equal(t, 1, len(ps.subs[keyAnyEvent]))
	chMatch, ok := ps.subs[keyAnyEvent]
	assert.True(t, ok)
	assert.Equal(t, chMatch[0], ch)
}

func TestPubsub_subEvent(t *testing.T) {
	ps := newPubsub()
	ch := ps.subByKey("AgentLogoff")
	assert.Equal(t, 1, len(ps.subs))
	assert.Equal(t, 1, len(ps.subs["agentlogoff"]))
	chMatch, ok := ps.subs["agentlogoff"]
	assert.True(t, ok)
	assert.Equal(t, chMatch[0], ch)
}

func TestPubsub_subscribe_disabled(t *testing.T) {
	ps := newPubsub()
	ps.disable()

	ch1 := ps.subAnyEvent()
	assert.Nil(t, ch1)
	ch2 := ps.subByKey("AgentLogoff")
	assert.Nil(t, ch2)
}

func TestPubsub_publish_anyEvent(t *testing.T) {
	ps := newPubsub()
	ch := ps.subAnyEvent()
	defer close(ch)

	m := dummyEvent("Newchannel")
	ps.publish(m)

	msg := <-ch
	assert.Equal(t, m, msg)
}

func TestPubsub_publish_anyEvent_filtered(t *testing.T) {
	ps := newPubsub()
	ch := ps.subAnyEvent()
	defer close(ch)

	ps.publish(dummyResponse("Success", ""))
	ps.publish(dummyResponse("Failed", ""))
	ps.publish(dummyEvent("Newchannel"))

	msg := <-ch
	assert.Equal(t, "Newchannel", msg.Field("event"))
}

func TestPubsub_publish_event(t *testing.T) {
	ps := newPubsub()
	ch := ps.subByKey("NewExten")
	defer close(ch)

	ps.publish(dummyEvent("Newexten"))

	msg := <-ch
	assert.Equal(t, "Newexten", msg.Field("event"))
}

func TestPubsub_publish_event_filtered(t *testing.T) {
	ps := newPubsub()
	ch := ps.subByKey("NewExten")
	defer close(ch)

	ps.publish(dummyEvent("NewCallerid"))
	ps.publish(dummyEvent("NewConnectedLine"))
	ps.publish(dummyResponse("Success", ""))
	ps.publish(dummyEvent("Newexten"))

	msg := <-ch
	assert.Equal(t, "Newexten", msg.Field("event"))
}

func TestPubsub_publish_action_response(t *testing.T) {
	ps := newPubsub()

	m := dummyResponse("Success", "Filtered message")
	m.AddActionID()
	actionId := m.ActionID()

	ch := ps.subByKey(actionId)
	defer close(ch)

	ps.publish(dummyResponse("Failed", ""))
	ps.publish(dummyEvent("Newexten"))
	ps.publish(m)

	msg := <-ch
	assert.True(t, msg.IsResponse())
	assert.Equal(t, "Success", msg.Field("Response"))
	assert.Equal(t, "Filtered message", msg.Field("Message"))
}

func TestPubsub_publish_on_disabled(t *testing.T) {
	ps := newPubsub()
	ps.disable()
	assert.False(t, ps.publish(dummyEvent("NewChannel")))
}

func TestPubsub_destroy(t *testing.T) {
	ps := newPubsub()
	ps.subByKey("NewExten")
	ps.subByKey("QueueMember")
	ps.subByKey("NewChannel")

	assert.Equal(t, 3, len(ps.subs))

	ps.subByKey("NewChannel")
	assert.Equal(t, 3, len(ps.subs))
	chN := 0
	for _, c := range ps.subs {
		chN += len(c)
	}
	assert.Equal(t, 4, chN)
	assert.False(t, ps.off)

	ps.destroy()
	assert.True(t, ps.off)
	assert.Equal(t, 0, len(ps.subs))
}

func TestPubsub_publish_subscribe_after_destroy(t *testing.T) {
	ps := newPubsub()
	ch := ps.subAnyEvent()
	ps.destroy()

	assert.Nil(t, ps.subscribe("foo"))
	assert.False(t, ps.publish(dummyEvent("QueueParams")))
	_, ok := <-ch
	assert.False(t, ok)
}
