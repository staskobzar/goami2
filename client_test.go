package goami2

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func amiResponseLogin(conn net.Conn, response string) {
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
		response += "ActionID: " + actionid + "\r\n\r\n"
		conn.Write([]byte(response))
	}()
}

func TestConnectLoginSuccess(t *testing.T) {
	response := "Response: Success\r\nMessage: Authentication accepted\r\n"
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()
	client, err := NewClient(rConn)
	defer client.Close()
	assert.Nil(t, err)

	amiResponseLogin(wConn, response)

	err = client.Login("admin", "password")
	assert.Nil(t, err)
}

/*
func TestConnectLoginFail(t *testing.T) {
	response := "Response: Error\r\nMessage: Authentication failed\r\n"
	wConn, rConn := net.Pipe()
	go func() { wConn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")) }()

	ch, err := connPipeline(context.TODO(), rConn)
	assert.Nil(t, err)

	amiResponseLogin(wConn, response)

	ok := connLogin("admin", "pass", rConn, ch)
	assert.False(t, ok)
}
*/
