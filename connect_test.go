package goami2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestConnectPipeline(t *testing.T) {
	inStream := []string{
		"Event: FullyBooted\r\n",
		"Privilege: system,all\r\n",
		"Status: System is Fully Booted\r\n",
		"\r\n",
		"Event: Newchannel\r\n",
		"Channel: PJSIP/misspiggy-00000001\r\n",
		"Exten: 31337\r\n",
		"Context: inbound\r\n",
		"\r\n",
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
	assert.Equal(t, ev1.Field("event"), "FullyBooted")
	assert.Equal(t, ev1.Field("status"), "System is Fully Booted")
	assert.True(t, ev2.IsEvent())
	assert.Equal(t, ev2.Field("event"), "Newchannel")
	assert.Equal(t, ev2.Field("exten"), "31337")
}
