package goami2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAMIMsgEvent(t *testing.T) {
	inStream := []string{
		"Event: Newchannel",
		"Channel: PJSIP/misspiggy-00000001",
		"Exten: 31337",
		"ActionId: SDY4-12837-123878782",
		"Context: inbound",
	}
	msg := NewAMIMsg()
	for _, s := range inStream {
		msg.addFieldOrEOF(s)
	}
	assert.Equal(t, msg.Type(), Event)
	assert.Empty(t, msg.Field("Foo"))
	assert.Equal(t, msg.Field("Context"), "inbound")
	assert.Equal(t, msg.ActionId(), "SDY4-12837-123878782")
}

func TestAMIMsgResponse(t *testing.T) {
	inStream := []string{
		"Response: Success",
		"EventList: start",
		"Message: Channels will follow",
	}
	resp := NewAMIMsg()
	for _, s := range inStream {
		resp.addFieldOrEOF(s)
	}
	assert.Equal(t, resp.Type(), Response)
	assert.True(t, resp.IsResponse())
	assert.True(t, resp.IsSuccess())
	assert.True(t, resp.IsEventList())
	assert.Equal(t, resp.Message(), "Channels will follow")
}
