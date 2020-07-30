package goami2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAMIMsgEvent(t *testing.T) {
	inStr := "Event: Newchannel\r\n" +
		"Channel: PJSIP/misspiggy-00000001\r\n" +
		"Exten: 31337\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Context: inbound\r\n"

	msg := NewAMIMsg(inStr)
	assert.True(t, msg.IsEvent())
	assert.Empty(t, msg.Field("Foo"))
	assert.Equal(t, msg.Field("Context"), "inbound")
	actionid, ok := msg.ActionID()
	assert.True(t, ok)
	assert.Equal(t, actionid, "SDY4-12837-123878782")
}

func TestAMIMsgResponse(t *testing.T) {
	inStr := "Response: Success\r\n" +
		"EventList: start\r\n" +
		"Message: Channels will follow\r\n"
	resp := NewAMIMsg(inStr)
	assert.True(t, resp.IsResponse())
	assert.True(t, resp.IsSuccess())
	assert.True(t, resp.IsEventList())
	assert.Equal(t, resp.Message(), "Channels will follow")
}

func TestAMIMsgEventNotFound(t *testing.T) {
	inStr := "Response: Success\r\n" +
		"EventList: start\r\n" +
		"Message: Channels will follow\r\n"
	resp := NewAMIMsg(inStr)
	_, ok := resp.Event()
	assert.False(t, ok)
}

func TestAMIMsgToJson(t *testing.T) {
	inStr := "Event: Newchannel\r\n" +
		"Channel: PJSIP/misspiggy-00000001\r\n" +
		"Exten: 31337\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Context: inbound\r\n"

	msg := NewAMIMsg(inStr)

	json := msg.JSON()
	matchStr := `{"event":"Newchannel","channel":"PJSIP/misspiggy-00000001",` +
		`"exten":"31337","actionid":"SDY4-12837-123878782",` +
		`"context":"inbound"}`
	assert.JSONEq(t, matchStr, json)

	// empty message return empty string
	msg = NewAMIMsg("")
	json = msg.JSON()
	assert.Equal(t, "", "")
}
