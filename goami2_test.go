package goami2

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientContext(t *testing.T) {

	t.Run("fail on login", func(t *testing.T) {
		connClient, connSrv := net.Pipe()
		connSrvSess(connSrv, []string{"invalid message\r\n\r\n"})

		_, err := NewClient(connClient, "admin", "pa55w0rd")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "AMI proto: invalid input")
	})

	t.Run("successfull login", func(t *testing.T) {
		connClient, connSrv := net.Pipe()
		connSrvSess(connSrv,
			[]string{
				"Response: Success\r\nMessage: Authentication accepted\r\n\r\n",
				"Event: DeviceStateChange\r\nPrivilege: call,all\r\n" +
					"Device: SIP/okon.ferry.sip.com\r\nState: INVALID\r\n\r\n",
			})
		cl, err := NewClient(connClient, "admin", "pa55w0rd")
		assert.Nil(t, err)
		msg := <-cl.AllMessages()
		assert.Equal(t, "DeviceStateChange", msg.Field("Event"))
		assert.Equal(t, "call,all", msg.Field("privilege"))
		assert.Equal(t, "SIP/okon.ferry.sip.com", msg.Field("Device"))
		assert.Equal(t, "INVALID", msg.Field("state"))
	})

	t.Run("terminate on conn close", func(t *testing.T) {
		connClient, connSrv := net.Pipe()
		connSrvSess(connSrv,
			[]string{"Response: Success\r\nMessage: Authentication accepted\r\n\r\n"})
		cl, err := NewClient(connClient, "admin", "pa55w0rd")
		assert.Nil(t, err)
		_ = connSrv.Close()
		err = <-cl.Err()
		assert.ErrorIs(t, err, ErrEOF)
	})
}
