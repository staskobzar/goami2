package goami2

import (
	"bufio"
	"net"
	"net/textproto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func srvPrompt(prompt string) net.Conn {
	r, w := net.Pipe()
	go func() {
		defer w.Close()
		if prompt == "timeout" {
			<-time.After(4 * time.Second)
		}
		w.Write([]byte(prompt))
	}()
	return r
}

func srvLogin(response string) net.Conn {
	cln, srv := net.Pipe()

	m := newMessage(textproto.MIMEHeader{"Response": []string{response}})
	m.AddField("Message", "Login response")
	go func() {
		defer srv.Close()
		scanner := bufio.NewScanner(srv)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				if response == "timeout" {
					<-time.After(4 * time.Second)
				}
				srv.Write(m.Bytes())
				break
			}
		}
	}()
	return cln
}

func amiFakeSrv(login string) net.Conn {
	cln, srv := net.Pipe()
	m := newMessage(textproto.MIMEHeader{"Response": []string{login}})
	m.AddField("Message", "Login response")
	go func() {
		defer srv.Close()
		// write prompt
		srv.Write([]byte("Asterisk Call Manager/2.10.4\n"))
		scanner := bufio.NewScanner(srv)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				srv.Write(m.Bytes())
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	return cln
}

func TestClient_readPrompt_invalid_prompt(t *testing.T) {
	conn := srvPrompt("foo\n")
	defer conn.Close()

	c, ctx := newClient(conn)
	err := c.readPrompt(ctx, NetTOUT)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorInvalidPrompt, err)
}

func TestClient_readPrompt_connect_timeout(t *testing.T) {
	conn := srvPrompt("timeout")
	defer conn.Close()

	c, ctx := newClient(conn)
	err := c.readPrompt(ctx, time.Millisecond*10)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorConnTOUT, err)
}

func TestClient_readPrompt_network_fail(t *testing.T) {
	conn := srvPrompt("Asterisk Call Manager/2.10.4\n")
	c, ctx := newClient(conn)
	conn.Close()
	err := c.readPrompt(ctx, NetTOUT)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Network error")
}

func TestClient_readPrompt(t *testing.T) {
	conn := srvPrompt("Asterisk Call Manager/2.10.4\n")
	defer conn.Close()

	c, ctx := newClient(conn)
	err := c.readPrompt(ctx, NetTOUT)
	assert.Nil(t, err)
}

func TestClient_login_success(t *testing.T) {
	conn := srvLogin("Success")
	defer conn.Close()

	c, ctx := newClient(conn)
	err := c.login(ctx, NetTOUT, "admin", "pa55w0rd")
	assert.Nil(t, err)
}

func TestClient_login_failed(t *testing.T) {
	conn := srvLogin("Error")
	defer conn.Close()

	c, ctx := newClient(conn)
	err := c.login(ctx, NetTOUT, "admin", "pa55w0rd")
	assert.NotNil(t, err)
	assert.Equal(t, ErrorLogin, err)
}

func TestClient_login_timeout(t *testing.T) {
	conn := srvLogin("timeout")
	defer conn.Close()

	c, ctx := newClient(conn)
	tout := 10 * time.Millisecond
	err := c.login(ctx, tout, "admin", "pa55w0rd")
	assert.NotNil(t, err)
	assert.Equal(t, ErrorConnTOUT, err)
}

func TestClient_NewClient_fail_login(t *testing.T) {
	conn := amiFakeSrv("Failed")
	_, err := NewClient(conn, "admin", "pass")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Login failed")
}

func TestClient_AnyEvent(t *testing.T) {
	conn := amiFakeSrv("Success")
	defer conn.Close()
	client, err := NewClient(conn, "admin", "pass")
	assert.Nil(t, err)
	ch := client.AllMessages()
	assert.NotNil(t, ch)
}

func TestClient_Action(t *testing.T) {
	conn := amiFakeSrv("Success")
	defer conn.Close()
	client, err := NewClient(conn, "admin", "pass")
	assert.Nil(t, err)
	ok := client.Action(NewAction("CoreStatus"))
	assert.True(t, ok)
}

func TestClient_Close(t *testing.T) {
	conn := amiFakeSrv("Success")
	client, err := NewClient(conn, "admin", "pass")
	assert.Nil(t, err)
	client.AllMessages()

	assert.Equal(t, 1, client.subs.lenMsgChans())

	client.Close()
	assert.Equal(t, 0, client.subs.lenMsgChans())
	_, err = conn.Write([]byte("foo"))
	assert.NotNil(t, err) // connection is already closed
}

func TestClient_fails_subscribe_on_closed(t *testing.T) {
	conn := amiFakeSrv("Success")
	client, err := NewClient(conn, "admin", "pass")
	assert.Nil(t, err)

	client.Close()
	ch := client.OnEvent("NewExten")
	assert.Nil(t, ch)
	ch = client.AllMessages()
	assert.Nil(t, ch)
	ok := client.Action(NewAction("CoreStatus"))
	assert.False(t, ok)
}

func TestClient_Err_channel(t *testing.T) {
	conn := amiFakeSrv("Success")
	client, err := NewClient(conn, "admin", "pass")
	defer client.Close()
	assert.Nil(t, err)

	conn.Close()
	err = <-client.Err()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Network error")
}
