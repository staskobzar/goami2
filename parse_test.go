package goami2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const rawPack = "Event: Newstate\r\n" +
	"Privilege: call,all\r\n" +
	"Channel: SIP/9170-12-0000003c\r\n" +
	"ChannelState: 4\r\n" +
	"ChannelStateDesc: Ring\r\n" +
	"CallerIDNum: 9170\r\n" +
	"CallerIDName: Bud Heller\r\n" +
	"ConnectedLineNum: <unknown>\r\n" +
	"ConnectedLineName: <unknown>\r\n" +
	"Language: fr\r\n" +
	"AccountCode: \r\n" +
	"Context: default\r\n" +
	"Exten: 9898\r\n" +
	"Priority: 1\r\n" +
	"Address-IP: 10.0.0.1\r\n" +
	"Uniqueid: 1598887681.60\r\n" +
	"Linkedid: 1598887681.60\r\n" +
	"ChanVariable: realm=\r\n" +
	"ChanVariable: SIPDOMAIN=okon.sip.com\r\n" +
	"ChanVariable: SIPCALLID=b5ser03agv7huo7faai5\r\n\r\n"

func TestParseEvent(t *testing.T) {
	input := "Event: DeviceStateChange\r\n" +
		"Privilege: call,all\r\n" +
		"Device: SIP/9170-12\r\n" +
		"State: INUSE\r\n\r\n"

	msg, err := Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, 4, msg.Len())

	t.Run("bool event or response", func(t *testing.T) {
		assert.True(t, msg.IsEvent())
		assert.False(t, msg.IsResponse())
		assert.False(t, msg.IsSuccess())
	})

	t.Run("fields values", func(t *testing.T) {
		assert.Equal(t, "DeviceStateChange", msg.Field("event"))
		assert.Equal(t, "call,all", msg.Field("Privilege"))
		assert.Equal(t, "SIP/9170-12", msg.Field("device"))
		assert.Equal(t, "INUSE", msg.Field("State"))
	})

	t.Run("Message to string match input", func(t *testing.T) {
		assert.Equal(t, input, msg.String())
	})

	t.Run("Message to bytes match input", func(t *testing.T) {
		assert.Equal(t, []byte(input), msg.Byte())
	})
}

func TestParseEventOK(t *testing.T) {
	tests := map[string][]string{
		`endpoint call`:       getAmiFixtureCall(),
		`call to queue`:       getAmiFixtureQueues(),
		`queue memeber pause`: getAmiFixtureQueueMembers(),
		`conference bridge`:   getAmiFixtureConfBridge(),
	}

	for name, packets := range tests {
		t.Run(name, func(t *testing.T) {
			for _, input := range packets {
				msg, err := Parse(input)
				assert.Nil(t, err)
				assert.NotNil(t, msg)
			}
		})
	}
}

func TestParseEventFail(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		`empty string`: {
			"", "unexpected end of input",
		},
		`incomplete packet input`: {
			"Event: MusicOnHoldStop\r\nPrivilege: call,all\r\nChannel: SIP/1031-12-0000006b\r\nChannelState: 6\r\nChannelStateDesc: Up",
			"unexpected end of input",
		},
		`invalid string`: {
			"foo bar", "invalid input",
		},
		`invalid header`: {
			"Event: MoH\r\nPriv:all\r\nChan\r\n\r\n", "invalid input",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg, err := Parse(tc.input)
			assert.ErrorIs(t, err, ErrAMI)
			assert.ErrorContains(t, err, tc.want)
			assert.Nil(t, msg)
		})
	}
}

func TestParseResponse(t *testing.T) {
	input := "Response: Success\r\n" +
		"ActionID: 175a511275bae@okon.ferry.sip.com\r\n" +
		"EventList: start\r\n" +
		"Message: Queue status will follow\r\n\r\n"

	msg, err := Parse(input)
	assert.Nil(t, err)
	assert.Equal(t, 4, msg.Len())
	assert.True(t, msg.IsResponse())
	assert.True(t, msg.IsSuccess())
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Parse(rawPack)
	}
}

func BenchmarkMessageToString(b *testing.B) {
	msg, _ := Parse(rawPack)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = msg.String()
	}
}
