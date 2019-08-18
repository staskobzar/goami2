package goami2

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectPipeline(t *testing.T) {
	inStream := []string{
		"Event: FullyBooted\r\nPrivilege: system,all\r\n",
		"Status: System is Fully Booted\r\n\r\n",
		"Event: Newchannel\r\nChannel: PJSIP/misspiggy-00000001\r\n",
		"Exten: 31337\r\nContext: inbound\r\n\r\n",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()

	ch, err := connPipeline(ctx, rConn)
	assert.Nil(t, err)

	go func() {
		for _, s := range inStream {
			wConn.Write([]byte(s))
		}
	}()
	ev1 := <-ch
	ev2 := <-ch
	assert.True(t, ev1.IsEvent())
	assert.Equal(t, ev1.Field("Event"), "FullyBooted")
	assert.Equal(t, ev1.Field("status"), "System is Fully Booted")
	assert.True(t, ev2.IsEvent())
	assert.Equal(t, ev2.Field("event"), "Newchannel")
	assert.Equal(t, ev2.Field("Exten"), "31337")
}

func BenchmarkConnectPipelineEvent(b *testing.B) {
	inStream := []string{
		"Event: RTCPSent\r\n",
		"Privilege: reporting,all\r\n",
		"Channel: SIP/router01-00008bc7\r\n",
		"ChannelState: 6\r\n",
		"ChannelStateDesc: Up\r\n",
		"CallerIDNum: 4504717783\r\n",
		"CallerIDName: HVE:ANIMONDE TERREB\r\n",
		"ConnectedLineNum: 223\r\n",
		"ConnectedLineName: HVE Chirurgie\r\n",
		"Language: fr\r\n",
		"AccountCode:\r\n",
		"Context: default\r\n",
		"Exten: MYQUEUE_203\r\n",
		"Priority: 11\r\n",
		"Uniqueid: 1555773835.102725\r\n",
		"Linkedid: 1555773835.102725\r\n",
		"To: 199.182.134.111:10131\r\n",
		"From: 199.182.134.120:15405\r\n",
		"SSRC: 0x7fb143e7\r\n",
		"PT: 200(SR)\r\n",
		"ReportCount: 1\r\n",
		"SentNTP: 1555774000.629136\r\n",
		"SentRTP: 1883844000\r\n",
		"SentPackets: 8069\r\n",
		"SentOctets: 1291040\r\n",
		"Report0SourceSSRC: 0x74816fee\r\n",
		"Report0FractionLost: 0\r\n",
		"Report0CumulativeLost: 0\r\n",
		"Report0HighestSequence: 65424\r\n",
		"Report0SequenceNumberCycles: 0\r\n",
		"Report0IAJitter: 0\r\n",
		"Report0LSR: 3098583250\r\n",
		"Report0DLSR: 0.0000\r\n",
		"\r\n",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()

	ch, _ := connPipeline(ctx, rConn)

	go func() {
		for i := 0; i < b.N; i++ {
			for _, s := range inStream {
				wConn.Write([]byte(s))
			}

		}
	}()
	N := 0
	for msg := range ch {
		N++
		msg.IsEvent()
		if N >= b.N {
			break
		}
	}
}
