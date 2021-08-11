package goami2

import (
	"fmt"
	"net/textproto"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_NewAction(t *testing.T) {
	a := NewAction("Agents")
	assert.Equal(t, "Action: Agents\r\n\r\n", a.String())
}

func TestMessage_AddField(t *testing.T) {
	a := NewAction("SendText")
	a.AddField("message", "hello world")
	want := []string{"Action: SendText\r\n",
		"Message: hello world\r\n"}
	assert.Contains(t, a.String(), want[0])
	assert.Contains(t, a.String(), want[1])
	assert.True(t, strings.HasSuffix(a.String(), "\r\n\r\n"))
}

func TestMessage_Bytes(t *testing.T) {
	a := NewAction("Agents")
	assert.Equal(t, []byte("Action: Agents\r\n\r\n"), a.Bytes())
}

func TestMessage_ActionID(t *testing.T) {
	a := NewAction("CoreStatus")
	a.AddField("actionid", "foo@123")
	id := a.ActionID()
	want := []string{"Action: CoreStatus\r\n", "Actionid: foo@123\r\n"}
	assert.Contains(t, a.String(), want[0])
	assert.Contains(t, a.String(), want[1])
	assert.Equal(t, "foo@123", id)
	assert.True(t, strings.HasSuffix(a.String(), "\r\n\r\n"))
}

func TestMessage_AddActionID(t *testing.T) {
	a := NewAction("CoreStatus")
	a.AddActionID()
	id := a.ActionID()

	want := []string{fmt.Sprintf("Actionid: %s\r\n", id),
		"Action: CoreStatus\r\n"}
	assert.True(t, len(id) > 0)
	assert.Contains(t, a.String(), want[0])
	assert.Contains(t, a.String(), want[1])
	assert.True(t, strings.HasSuffix(a.String(), "\r\n\r\n"))
}

func TestMessage_Field(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Exten", "31337")
	m.AddField("Actionid", "SDY4-12837-123878782")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern=true")
	m.AddField("ChanVariable", "realm=sip.pbx.com")
	m.AddField("ChanVariable", "account=123")

	assert.Equal(t, "31337", m.Field("exten"))
	assert.Equal(t, "inbound", m.Field("Context"))
	assert.Equal(t, "", m.Field("Foo"))
	assert.Equal(t, "DIR=inbound", m.Field("variable"))
	assert.Equal(t, "realm=sip.pbx.com", m.Field("chanvariable"))
}

func TestMessage_DelField(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Exten", "31337")
	m.AddField("Actionid", "SDY4-12837-123878782")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern=true")
	m.AddField("ChanVariable", "realm=sip.pbx.com")
	m.AddField("ChanVariable", "account=123")

	assert.Equal(t, "31337", m.Field("exten"))
	m.DelField("exten")
	assert.Equal(t, "", m.Field("exten"))
	assert.ElementsMatch(t, []string{"realm=sip.pbx.com", "account=123"},
		m.FieldValues("ChanVariable"))
	m.DelField("chanvariable")
	assert.Empty(t, m.FieldValues("chanvariable"))
}

func TestMessage_FieldValues(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern=1")
	m.AddField("Set", "realm=sip.com")
	m.AddField("Set", "account=123")

	assert.Equal(t, []string{"31337"}, m.FieldValues("exten"))
	assert.Equal(t, []string{"inbound"}, m.FieldValues("Context"))
	assert.ElementsMatch(t, []string{"DIR=inbound", "extern=1"},
		m.FieldValues("variable"))
	assert.ElementsMatch(t, []string{"realm=sip.com", "account=123"},
		m.FieldValues("Set"))
	assert.Nil(t, m.FieldValues("Foo"))
}

func TestMessage_IsEvent(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")

	assert.True(t, m.IsEvent())
	assert.False(t, m.IsResponse())
}

func TestMessage_IsResponse(t *testing.T) {
	header := textproto.MIMEHeader{"Response": []string{"Success"}}
	m := newMessage(header)
	m.AddField("Message", "All is good")

	assert.False(t, m.IsEvent())
	assert.True(t, m.IsResponse())
}

func TestMessage_IsSuccess(t *testing.T) {
	header := textproto.MIMEHeader{"Response": []string{"Success"}}
	m := newMessage(header)
	m.AddField("Message", "All is good")

	assert.True(t, m.IsSuccess())
}

func TestMessage_IsSuccess_false(t *testing.T) {
	header := textproto.MIMEHeader{"Response": []string{"Failed"}}
	m := newMessage(header)
	m.AddField("Message", "All failed")

	assert.False(t, m.IsSuccess())
}

func TestMessage_Var(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern")
	m.AddField("ChanVariable", "realm=sip.com")
	m.AddField("Chanvariable", "account=123")
	m.AddField("ParkeeChanVariable", "parkctx=acc700")
	m.AddField("ParkeeChanVariable", "parksrv=pbx001.phone.com")

	v, ok := m.Var("DIR")
	assert.True(t, ok)
	assert.Equal(t, "inbound", v)

	v, ok = m.Var("realm")
	assert.True(t, ok)
	assert.Equal(t, "sip.com", v)

	v, ok = m.Var("account")
	assert.True(t, ok)
	assert.Equal(t, "123", v)

	v, ok = m.Var("extern")
	assert.True(t, ok)
	assert.Empty(t, v)

	v, ok = m.Var("parkctx")
	assert.True(t, ok)
	assert.Equal(t, "acc700", v)

	v, ok = m.Var("parksrv")
	assert.True(t, ok)
	assert.Equal(t, "pbx001.phone.com", v)

	v, ok = m.Var("foo")
	assert.False(t, ok)
	assert.Empty(t, v)
}

func TestMessage_newMessage(t *testing.T) {
	header := textproto.MIMEHeader{
		"Event":     []string{"DeviceStateChange"},
		"Device":    []string{"SIP/9170-12"},
		"Privilege": []string{"call,all"},
		"State":     []string{"INUSE"},
	}
	m := newMessage(header)
	assert.Contains(t, m.String(), "Event: DeviceStateChange\r\n")
	assert.Contains(t, m.String(), "State: INUSE\r\n")
	assert.True(t, strings.HasSuffix(m.String(), "\r\n\r\n"))
}

func TestMessage_newMessage_Chanvariable(t *testing.T) {
	header := textproto.MIMEHeader{
		"Event":  []string{"DeviceStateChange"},
		"Device": []string{"SIP/9170-12"},
		"Chanvariable": []string{
			"realm=okon.ferry.xyz",
			"SIPDOMAIN=",
			"SIPCALLID=3aff7b@10.2.3.18:5060",
		},
		"State": []string{"INUSE"},
	}
	m := newMessage(header)
	assert.Contains(t, m.String(), "Chanvariable: realm=okon.ferry.xyz\r\n")
	assert.Contains(t, m.String(), "Chanvariable: SIPDOMAIN=\r\n")
	assert.Contains(t, m.String(), "Chanvariable: SIPCALLID=3aff7b@10.2.3.18:5060\r\n")
	assert.True(t, strings.HasSuffix(m.String(), "\r\n\r\n"))
}

func TestMessage_ToJSON(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Channel", "PJSIP/misspiggy-00000001")
	m.AddField("Exten", "31337")
	m.AddField("Actionid", "SDY4-12837-123878782")
	m.AddField("Context", "inbound")

	jstr := m.JSON()
	want := `{"event":"Newchannel","channel":"PJSIP/misspiggy-00000001",` +
		`"exten":"31337","actionid":"SDY4-12837-123878782",` +
		`"context":"inbound"}`
	assert.JSONEq(t, want, jstr)
}

func TestMessage_ToJSON_Variables(t *testing.T) {
	header := textproto.MIMEHeader{"Event": []string{"Newchannel"}}
	m := newMessage(header)
	m.AddField("Channel", "PJSIP/misspiggy-00000001")
	m.AddField("Exten", "31337")
	m.AddField("Actionid", "SDY4-12837-123878782")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern=true")
	m.AddField("Variable", "FOO")
	m.AddField("ChanVariable", "realm=sip.pbx.com")
	m.AddField("ChanVariable", "account=123")

	jstr := m.JSON()
	want := `{"actionid":"SDY4-12837-123878782","channel":"PJSIP/misspiggy-00000001",` +
		`"context":"inbound","event":"Newchannel","exten":"31337",` +
		`"chanvariable":{"account":"123","realm":"sip.pbx.com"},` +
		`"variable":{"DIR":"inbound","extern":"true","FOO":""}}`
	assert.JSONEq(t, want, jstr)
}

func TestMessage_ToJSON_Empty(t *testing.T) {
	m := newMessage(textproto.MIMEHeader{})
	jstr := m.JSON()
	assert.Equal(t, "{}", jstr)
}

func TestMessage_FromJSON(t *testing.T) {
	jstr := `{"action":"ConfbridgeKick","conference":"Sales",` +
		`"channel":"Local/Sales-65f4a00b-001"}`
	m, err := FromJSON(jstr)
	assert.Nil(t, err)

	assert.Equal(t, "ConfbridgeKick", m.Field("Action"))
	assert.Equal(t, "Sales", m.Field("Conference"))
	assert.Equal(t, "Local/Sales-65f4a00b-001", m.Field("Channel"))
}

func TestMessage_FromJSON_Multilevel(t *testing.T) {
	jstr := `{"actionid":"SDY4-12837-123878782","channel":"PJSIP/misspiggy-00000001",` +
		`"context":"inbound","event":"Newchannel","exten":"31337",` +
		`"chanvariable":{"account":"123","realm":"sip.pbx.com"},` +
		`"variable":{"DIR":"inbound","extern":"true","FOO":""}}`
	m, err := FromJSON(jstr)
	assert.Nil(t, err)

	assert.Equal(t, "Newchannel", m.Field("Event"))
	assert.Equal(t, "SDY4-12837-123878782", m.ActionID())
	assert.Equal(t, "PJSIP/misspiggy-00000001", m.Field("channel"))
	assert.Equal(t, "31337", m.Field("exten"))
	assert.ElementsMatch(t, []string{"account=123", "realm=sip.pbx.com"},
		m.FieldValues("chanvariable"))
	assert.ElementsMatch(t, []string{"DIR=inbound", "extern=true", "FOO="},
		m.FieldValues("variable"))
}

func TestMessage_loginMessage(t *testing.T) {
	msg := loginMessage("james", "passwd")
	assert.Equal(t, "Login", msg.Field("action"))
	assert.Equal(t, "james", msg.Field("Username"))
	assert.Equal(t, "passwd", msg.Field("Secret"))
}
