package goami2

import (
	"bufio"
	"context"
	"github.com/stretchr/testify/assert"
	"net"
	"strings"
	"testing"
	"time"
)

func TestScannerSplitFunc(t *testing.T) {
	input := "Event: FullyBooted\r\nPrivilege: system,all\r\n\r\n" +
		"Status: Fully Booted\r\n\r\n" +
		"Response: Success\r\nActionID: foo-bar1\r\n" +
		"EventList: start\r\nMessage: Queue summary will follow\r\n\r\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(scanAMIMsg)

	count := 0
	for scanner.Scan() {
		count++
	}
	assert.Equal(t, count, 3)
}

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

	msg := <-eventChan
	assert.True(t, msg.IsEvent())
	event, ok := msg.Event()
	assert.True(t, ok)
	assert.Equal(t, event, "Newchannel")
	assert.Equal(t, msg.Field("Exten"), "31337")
	assert.Equal(t, msg.Field("Context"), "inbound")
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
