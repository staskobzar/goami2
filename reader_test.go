package goami2

import (
	"context"
	"net"
	"testing"
)

func TestReaderReadEvent(t *testing.T) {
	inStream := []byte(
		"Event: Newchannel\r\n" +
			"Channel: PJSIP/misspiggy-00000001\r\n" +
			"Exten: 31337\r\n" +
			"Context: inbound\r\n\r\n",
	)

	rConn, wConn := net.Pipe()
	defer rConn.Close()
	defer wConn.Close()
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

	n := 0
	for reader := range eventChan {
		if reader.str == "\r\n" {
			break
		}
		n++
	}
	if n != 4 {
		t.Errorf("Invalid number of packets read. Expected %d, got %d", 4, n)
	}
}

func TestReaderEarlyCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := newReader(ctx, nil)
	if err == nil {
		t.Errorf("Failed early Reader cancel.")
	}
}

func TestReaderContextClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch, _ := newReader(ctx, nil)
	cancel()
	_, ok := <-ch
	if ok {
		t.Errorf("Failed to cancel reader")
	}
}
