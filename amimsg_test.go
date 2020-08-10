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

func TestAMIMsgAddField(t *testing.T) {
	inStr := "Event: Newchannel\r\n" +
		"Channel: PJSIP/misspiggy-00000001\r\n" +
		"Exten: 31337\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Context: inbound\r\n"

	msg := NewAMIMsg(inStr)
	msg.AddField("AMISrvID", "5")
	msg.AddField("AMIHost", "10.0.5.5:5038")
	assert.Equal(t, "5", msg.Field("AMISrvID"))
	assert.Equal(t, "10.0.5.5:5038", msg.Field("AMIHost"))
}

func TestAMIMsgDelField(t *testing.T) {
	inStr := "Event: Newchannel\r\n" +
		"Channel: PJSIP/misspiggy-00000001\r\n" +
		"Exten: 31337\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Context: inbound\r\n"

	msg := NewAMIMsg(inStr)
	assert.Equal(t, "inbound", msg.Field("Context"))
	assert.True(t, msg.DelField("context"))
	assert.False(t, msg.DelField("context"))
	assert.Equal(t, "", msg.Field("Context"))
}

func TestAMIMsgVariableField(t *testing.T) {
	inStr := "Action: Originate\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Channel: PJSIP/kermit-00000002\r\n" +
		"Context: outbound\r\n" +
		"Exten: s\r\n" +
		"Priority: 1\r\n" +
		"CallerID: \"Kermit the Frog\" <123-4567>\r\n" +
		"Account: FrogLegs\r\n" +
		"Variable: MY_VAR=frogs jump\r\n" +
		"Variable: HIDE_FROM_CHEF=true\r\n"
	msg := NewAMIMsg(inStr)
	v, ok := msg.Variable("foo")
	assert.False(t, ok)
	v, ok = msg.Variable("MY_VAR")
	assert.True(t, ok)
	assert.Equal(t, "frogs jump", v)
	v, ok = msg.Variable("HIDE_FROM_CHEF")
	assert.True(t, ok)
	assert.Equal(t, "true", v)
}

func TestAMIMsgChanVariableField(t *testing.T) {
	inStr := "Event: QueueJoin\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Channel: PJSIP/kermit-00000002\r\n" +
		"Context: outbound\r\n" +
		"Exten: s\r\n" +
		"Priority: 1\r\n" +
		"CallerID: \"Kermit the Frog\" <123-4567>\r\n" +
		"Account: FrogLegs\r\n" +
		"ChanVariable: realm=sip.pbx.com\r\n" +
		"ChanVariable: account=123\r\n"
	msg := NewAMIMsg(inStr)
	v, ok := msg.ChanVariable("foo")
	assert.False(t, ok)
	v, ok = msg.ChanVariable("realm")
	assert.True(t, ok)
	assert.Equal(t, "sip.pbx.com", v)
	v, ok = msg.ChanVariable("account")
	assert.True(t, ok)
	assert.Equal(t, "123", v)
}

func TestAMIMsgVarField(t *testing.T) {
	inStr := "Event: QueueJoin\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Channel: PJSIP/kermit-00000002\r\n" +
		"Context: outbound\r\n" +
		"Exten: s\r\n" +
		"Priority: 1\r\n" +
		"CallerID: \"Kermit the Frog\" <123-4567>\r\n" +
		"Account: FrogLegs\r\n" +
		"Variable: MY_VAR=frogs jump\r\n" +
		"ChanVariable: account=123\r\n"
	msg := NewAMIMsg(inStr)
	v, ok := msg.Var("foo")
	assert.False(t, ok)
	v, ok = msg.Var("account")
	assert.True(t, ok)
	assert.Equal(t, "123", v)
	v, ok = msg.Var("MY_VAR")
	assert.True(t, ok)
	assert.Equal(t, "frogs jump", v)
}

func TestAMIMsgVarFieldWhereVarHasNoValue(t *testing.T) {
	inStr := "Event: QueueJoin\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Channel: PJSIP/kermit-00000002\r\n" +
		"Context: outbound\r\n" +
		"Exten: s\r\n" +
		"Priority: 1\r\n" +
		"CallerID: \"Kermit the Frog\" <123-4567>\r\n" +
		"Account: FrogLegs\r\n" +
		"Variable: FOO\r\n" +
		"Variable: account=\r\n" +
		"ChanVariable: realm=\r\n" +
		"ChanVariable: SIPURI\r\n"
	msg := NewAMIMsg(inStr)
	v, ok := msg.Var("foo")
	assert.False(t, ok)

	v, ok = msg.Var("FOO")
	assert.True(t, ok)
	assert.Empty(t, v)

	v, ok = msg.Var("account")
	assert.True(t, ok)
	assert.Empty(t, v)

	v, ok = msg.Var("realm")
	assert.True(t, ok)
	assert.Empty(t, v)

	v, ok = msg.Var("SIPURI")
	assert.True(t, ok)
	assert.Empty(t, v)
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

func TestAMIMsgToJsonWithVariables(t *testing.T) {
	inStr := "Event: Newchannel\r\n" +
		"Channel: PJSIP/misspiggy-00000001\r\n" +
		"Exten: 31337\r\n" +
		"ActionId: SDY4-12837-123878782\r\n" +
		"Context: inbound\r\n" +
		"Variable: DIR=inbound\r\n" +
		"Variable: extern=true\r\n" +
		"ChanVariable: realm=sip.pbx.com\r\n" +
		"ChanVariable: account=123\r\n"

	msg := NewAMIMsg(inStr)

	json := msg.JSON()

	matchStr := `{"actionid":"SDY4-12837-123878782","channel":"PJSIP/misspiggy-00000001",` +
		`"context":"inbound","event":"Newchannel","exten":"31337",` +
		`"chanvariable":{"account":"123","realm":"sip.pbx.com"},` +
		`"variable":{"DIR":"inbound","extern":"true"}}`
	assert.JSONEq(t, matchStr, json)
}
