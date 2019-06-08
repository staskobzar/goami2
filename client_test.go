package goami2

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"strings"
	"testing"
)

func amiResponseServer(conn net.Conn, response []string) {
	go func() {
		scanner := bufio.NewScanner(conn)
		var actionid string
		for scanner.Scan() {
			s := scanner.Text()
			fmt.Sscanf(s, "ActionID: %s", &actionid)
			if s == "" {
				break
			}
		}
		for _, s := range response {
			s = strings.Replace(s, "$ACTIONID$", actionid, 1)
			conn.Write([]byte(s))
		}
	}()
}

func TestClientLoginSuccess(t *testing.T) {
	response := []string{
		"Response: Success\r\n",
		"ActionID: $ACTIONID$\r\n",
		"Message: Authentication accepted\r\n\r\n",
	}
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)

	amiResponseServer(wConn, response)

	err = client.Login("admin", "password")
	assert.Nil(t, err)
	assert.Zero(t, len(client.p.actionConsumer))
}

func TestClientLoginFailed(t *testing.T) {
	response := []string{
		"Response: Error\r\n",
		"ActionID: $ACTIONID$\r\n",
		"Message: Authentication failed\r\n\r\n",
	}
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)

	amiResponseServer(wConn, response)

	err = client.Login("admin", "password")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Authentication failed")
	assert.Zero(t, len(client.p.actionConsumer))
}

func TestClientEventList(t *testing.T) {
	response := []string{
		"Response: Success\r\n",
		"ActionID: $ACTIONID$\r\n",
		"EventList: start\r\n",
		"Message: Queue summary will follow\r\n",
		"\r\n",
		"Event: QueueSummary\r\n",
		"Queue: Sales\r\n",
		"LoggedIn: 2\r\n",
		"Available: 2\r\n",
		"Callers: 0\r\n",
		"HoldTime: 0\r\n",
		"TalkTime: 0\r\n",
		"LongestHoldTime: 0\r\n",
		"ActionID: $ACTIONID$\r\n",
		"\r\n",
		"Event: QueueSummary\r\n",
		"Queue: Reception_12\r\n",
		"LoggedIn: 2\r\n",
		"Available: 2\r\n",
		"Callers: 0\r\n",
		"HoldTime: 0\r\n",
		"TalkTime: 0\r\n",
		"LongestHoldTime: 0\r\n",
		"ActionID: $ACTIONID$\r\n",
		"\r\n",
		"Event: QueueSummaryComplete\r\n",
		"ActionID: foo-bar1-OTHER-RESPONSE-MIX-TEST\r\n",
		"EventList: Complete\r\n",
		"ListItems: 8\r\n",
		"\r\n",
		"Event: QueueSummaryComplete\r\n",
		"ActionID: $ACTIONID$\r\n",
		"EventList: Complete\r\n",
		"ListItems: 2\r\n",
		"\r\n",
	}
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)

	amiResponseServer(wConn, response)
	ch, err := client.Action("QueueSummary", nil)
	assert.Nil(t, err)
	N := 0
	EN := 0
	for msg := range ch {
		N++
		if msg.IsEvent() {
			EN++
		}
	}
	assert.Zero(t, len(client.p.actionConsumer))
	assert.Equal(t, N, 4)
	assert.Equal(t, EN, 3)
}

func TestClientAnyEventConsumer(t *testing.T) {
	events := []string{
		"Event: Hold\r\n" +
			"Privilege: call,all\r\n" +
			"Channel: SIP/monvet.modulis.clusterpbx.ca-00008bc8\r\n" +
			"ConnectedLineName: HVE:ANIMONDE TERREB\r\n" +
			"Context: dial-simple\r\n" +
			"Exten: 221\r\n" +
			"Priority: 1\r\n" +
			"\r\n",
		"Event: MusicOnHoldStart\r\n" +
			"Privilege: call,all\r\n" +
			"Channel: Local/221-1259@from-queue-00008258;2\r\n" +
			"ChannelState: 6\r\n" +
			"Exten: 221\r\n" +
			"Priority: 2\r\n" +
			"Class: default@monvet.modulis.clusterpbx.ca\r\n" +
			"\r\n",
		"Event: RTCPSent\r\n" +
			"Privilege: reporting,all\r\n" +
			"Channel: SIP/cc.humark.clusterpbx.ca-00008bef\r\n" +
			"ChannelState: 6\r\n" +
			"ChannelStateDesc: Up\r\n" +
			"Context: dial-simple\r\n" +
			"Exten: 210\r\n" +
			"\r\n",
	}
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)

	ch, err := client.AnyEvent()
	assert.Nil(t, err)
	assert.Equal(t, len(client.p.eventConsumer), 1)

	go func() {
		for _, s := range events {
			wConn.Write([]byte(s))
		}
	}()

	NE := 0
	for i := 0; i < len(events); i++ {
		msg := <-ch
		if msg.IsEvent() {
			NE++
		}
	}
	assert.Equal(t, NE, 3)
}

func TestClientCustomEventConsumer(t *testing.T) {
	events := []string{
		"Event: Hold\r\n" +
			"Privilege: call,all\r\n" +
			"Channel: SIP/monvet.modulis.clusterpbx.ca-00008bc8\r\n" +
			"ConnectedLineName: HVE:ANIMONDE TERREB\r\n" +
			"Context: dial-simple\r\n" +
			"Exten: 221\r\n" +
			"Priority: 1\r\n" +
			"\r\n",
		"Event: MusicOnHold\r\n" +
			"Privilege: call,all\r\n" +
			"Channel: Local/221-1259@from-queue-00008258;2\r\n" +
			"ChannelState: 6\r\n" +
			"Exten: 221\r\n" +
			"Priority: 2\r\n" +
			"Class: default@monvet.modulis.clusterpbx.ca\r\n" +
			"\r\n",
		"Event: RTCPSent\r\n" +
			"Privilege: reporting,all\r\n" +
			"Channel: SIP/cc.humark.clusterpbx.ca-00008bef\r\n" +
			"ChannelState: 6\r\n" +
			"ChannelStateDesc: Up\r\n" +
			"Context: dial-simple\r\n" +
			"Exten: 210\r\n" +
			"\r\n",
	}
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)

	ch, err := client.OnEvent("MusicOnHold")
	assert.Nil(t, err)
	assert.Equal(t, len(client.p.eventConsumer), 1)

	go func() {
		for _, s := range events {
			wConn.Write([]byte(s))
		}
	}()

	msg := <-ch
	assert.True(t, msg.IsEvent())
	assert.Equal(t, msg.Field("channel"),
		"Local/221-1259@from-queue-00008258;2")
}

func TestClientAddExistingEventChan(t *testing.T) {
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	assert.Nil(t, err)
	_, err = client.OnEvent("MusicOnHold")
	assert.Nil(t, err)
	assert.Equal(t, len(client.p.eventConsumer), 1)

	_, err = client.OnEvent("MusicOnHold")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "already exists")
	assert.Equal(t, len(client.p.eventConsumer), 1)
}
