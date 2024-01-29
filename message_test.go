package goami2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageNewMessageFields(t *testing.T) {
	msg := NewMessage()
	msg.AddField("Event", "NewState")
	msg.AddField("Channel", "PJSIP/mpgy-0001")
	msg.AddField("Exten", "31337")
	msg.AddField("Foo", "bar")
	msg.AddField("Actionid", "SDY4-12837-123878782")
	msg.AddField("Context", "inbound")
	msg.AddField("Bar", "last")

	assert.Equal(t, 7, msg.Len())

	msg.DelField("foo")
	assert.Equal(t, 6, msg.Len())
	msg.DelField("ActionID")
	assert.Equal(t, 5, msg.Len())
	msg.DelField("Bar")
	assert.Equal(t, 4, msg.Len())

	want := "Event: NewState\r\nChannel: PJSIP/mpgy-0001\r\n" +
		"Exten: 31337\r\nContext: inbound\r\n\r\n"
	assert.Equal(t, want, msg.String())
}

func TestMessageFieldValues(t *testing.T) {
	input := "Event: Newstate\r\n" +
		"Channel: SIP/9170-12-0000003c\r\n" +
		"CallerIDNum: 9170\r\n" +
		"CallerIDName: Bud Heller\r\n" +
		"Exten: 9898\r\n" +
		"ChanVariable: realm=\r\n" +
		"ChanVariable: SIPDOMAIN=okon.sip.com\r\n" +
		"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n"

	msg, err := Parse(input)
	assert.Nil(t, err)

	hdrs := msg.FieldValues("ChanVariable")
	assert.Equal(t, 3, len(hdrs))
	assert.Equal(t, "realm=", hdrs[0])
	assert.Equal(t, "SIPCALLID=b5ser03agv7huo7faai5", hdrs[2])
}

func TestMessageAction(t *testing.T) {
	msg := NewAction("Status")
	assert.Equal(t, "Action: Status\r\n\r\n", msg.String())

	msg.AddField("ActionID", "foo@bar")
	assert.Equal(t, "foo@bar", msg.ActionID())

	msg.AddActionID()
	assert.NotEmpty(t, msg.ActionID())
	assert.NotEqual(t, "foo@bar", msg.ActionID())
}

func TestMessageSetField(t *testing.T) {
	msg := NewAction("Status")
	msg.AddField("Foo", "111")

	assert.Equal(t, "Action: Status\r\nFoo: 111\r\n\r\n", msg.String())

	msg.SetField("Foo", "222")
	assert.Equal(t, "Action: Status\r\nFoo: 222\r\n\r\n", msg.String())

	msg.SetField("Bar", "333")
	assert.Equal(t, "Action: Status\r\nFoo: 222\r\nBar: 333\r\n\r\n", msg.String())
}

func TestMessageVar(t *testing.T) {
	m := NewMessage()
	m.AddField("Event", "Newchannel")
	m.AddField("Exten", "31337")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern")
	m.AddField("ChanVariable", "realm=sip.com")
	m.AddField("Chanvariable", "account=123")
	m.AddField("ParkeeChanVariable", "parkctx=acc700")
	m.AddField("ParkeeChanVariable", "parksrv=pbx001.phone.com")
	m.AddField("OrigTransfererChanVariable", "OTRCV=atcv1000")
	m.AddField("SecondTransfererChanVariable", "STRCV=atcv2000")
	m.AddField("TransfereeChanVariable", "TECV=atcv3000")
	m.AddField("TransferTargetChanVariable", "TTCV=atcv4000")

	tests := []struct {
		key    string
		exists bool
		want   string
	}{
		{"DIR", true, "inbound"},
		{"realm", true, "sip.com"},
		{"account", true, "123"},
		{"extern", true, ""},
		{"parkctx", true, "acc700"},
		{"parksrv", true, "pbx001.phone.com"},
		{"OTRCV", true, "atcv1000"},
		{"STRcv", true, "atcv2000"},
		{"TECV", true, "atcv3000"},
		{"ttcv", true, "atcv4000"},
		{"FOO", false, ""},
	}

	for _, tc := range tests {
		val, ok := m.Var(tc.key)
		assert.Equal(t, tc.exists, ok, tc.key)
		assert.Equal(t, tc.want, val, tc.key)
	}
}

func TestMessageJSON(t *testing.T) {
	m := NewMessage()
	m.AddField("Event", "Newchannel")
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

func TestMessageJSONMultiValues(t *testing.T) {
	m := NewMessage()
	m.AddField("Event", "Newchannel")
	m.AddField("Channel", "PJSIP/misspiggy-00000001")
	m.AddField("Exten", "31337")
	m.AddField("Actionid", "SDY4-12837-123878782")
	m.AddField("Context", "inbound")
	m.AddField("Variable", "DIR=inbound")
	m.AddField("Variable", "extern=true")
	m.AddField("Variable", "FOO")
	m.AddField("ChanVariable", "realm=sip.pbx.com")
	m.AddField("ChanVariable", "account=123")

	str := m.JSON()
	want := `{"actionid":"SDY4-12837-123878782","channel":"PJSIP/misspiggy-00000001",` +
		`"context":"inbound","event":"Newchannel","exten":"31337",` +
		`"chanvariable":{"account":"123","realm":"sip.pbx.com"},` +
		`"variable":{"DIR":"inbound","extern":"true","FOO":""}}`

	assert.JSONEq(t, want, str)
}

func TestMessageToJSONEmpty(t *testing.T) {
	m := NewMessage()
	jstr := m.JSON()
	assert.Equal(t, "{}", jstr)
}

func TestMessageFromJSON(t *testing.T) {
	t.Run("successfully created message", func(t *testing.T) {
		jstr := `{"action":"ConfbridgeKick","conference":"Sales",` +
			`"channel":"Local/Sales-65f4a00b-001"}`
		m, err := FromJSON(jstr)
		assert.Nil(t, err)

		assert.Equal(t, "ConfbridgeKick", m.Field("Action"))
		assert.Equal(t, "Sales", m.Field("Conference"))
		assert.Equal(t, "Local/Sales-65f4a00b-001", m.Field("Channel"))
	})

	t.Run("fail parse from json", func(t *testing.T) {
		tests := []struct {
			input string
			want  error
		}{
			{"", ErrAMI},
			{"foo", ErrAMI},
			{`{"action":"Foo"`, ErrAMI},
		}

		for _, tc := range tests {
			_, err := FromJSON(tc.input)
			assert.ErrorIs(t, err, tc.want, tc.input)
		}
	})
}

func TestMessageFromJSONMultilevel(t *testing.T) {
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
