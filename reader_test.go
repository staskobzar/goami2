package goami2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestReaderReadEvent(t *testing.T) {
	inStream := []byte(
		"Event: Newchannel\r\n" +
			"Channel: PJSIP/misspiggy-00000001\r\n" +
			"Exten: 31337\r\n" +
			"Context: inbound\r\n\r\n",
	)

	rConn, wConn := net.Pipe()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		wConn.Write(inStream)
	}()

	eventChan, err := newReader(ctx, rConn)
	if err != nil {
		t.Errorf("Failed init Reader: %s", err)
		return
	}

	n, expectedN := 0, 4
	for str := range eventChan {
		if str == "\r\n" {
			break
		}
		n++
	}
	assert.Equal(t, expectedN, n)
}

func TestReaderEarlyCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := newReader(ctx, nil)
	assert.NotNil(t, err)
}

func TestReaderContextClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch, _ := newReader(ctx, nil)
	cancel()
	_, ok := <-ch
	assert.False(t, ok)
}

func TestReaderBuildMessage(t *testing.T) {
	inStream := []string{
		"Event: Newchannel\r\n",
		"Channel: PJSIP/misspiggy-00000001\r\n",
		"Exten: 31337\r\n",
		"Context: inbound\r\n",
		"\r\n",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chIn := make(chan string)
	chOut, err := msgBuilder(ctx, chIn)

	assert.Nil(t, err)

	for _, s := range inStream {
		chIn <- s
	}

	amiMsg := <-chOut
	assert.Equal(t, amiMsg.Field("event"), "Newchannel")
	assert.Equal(t, amiMsg.Type(), Event)
}

func TestReaderPrompt(t *testing.T) {
	rConn, wConn := net.Pipe()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n"))
	}()

	// prompt is OK
	err := readPrompt(ctx, rConn)
	assert.Nil(t, err)

	// prompt Not OK
	go func() {
		wConn.Write([]byte("Not good prompt\r\n"))
	}()
	err = readPrompt(ctx, rConn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Not good prompt")

	// early cancel
	cancel()
	err = readPrompt(ctx, rConn)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}

func TestReaderPromptTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	r, w := net.Pipe()
	defer w.Close()
	err := readPrompt(ctx, r)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}
