package goami2

import (
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockSrv struct {
	addr string
	user string
	pass string
	ln   net.Listener
}

func e2eServer(messages []string) *mockSrv {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	srv := &mockSrv{addr: ln.Addr().String(), user: "admin", pass: "pa55wd", ln: ln}

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		//prompt
		if _, err = conn.Write([]byte("Asterisk Call Manager/2.10.4\r\n")); err != nil {
			panic(err)
		}

		// login
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
		if !strings.Contains(string(buf), "name: admin") &&
			!strings.Contains(string(buf), "ecret: pa55wd") {
			msg := dummyResponse("Failed", "Auth failed")
			conn.Write(msg.Bytes())
			return
		}

		resp := dummyResponse("Success", "Auth ok")
		conn.Write(resp.Bytes())

		for _, msg := range messages {
			conn.Write([]byte(msg))
		}
	}()

	return srv
}

func TestE2E_AnyEvent(t *testing.T) {
	srv := e2eServer(getAmiFixtureCall())

	conn, err := net.Dial("tcp", srv.addr)
	if err != nil {
		panic(err)
	}
	client, err := NewClient(conn, srv.user, srv.pass)
	defer client.Close()
	assert.Nil(t, err)
	ch := client.AnyEvent()
	assert.NotNil(t, ch)

	want := 0
	for _, m := range getAmiFixtureCall() {
		if strings.HasPrefix(m, "Event: ") {
			want++
		}
	}

	have := 0
	for range ch {
		have++
		if have >= want {
			break
		}
	}
	assert.Equal(t, want, have)
}

func TestE2E_OnEvent(t *testing.T) {
	srv := e2eServer(getAmiFixtureCall())

	conn, err := net.Dial("tcp", srv.addr)
	if err != nil {
		panic(err)
	}
	client, err := NewClient(conn, srv.user, srv.pass)
	defer client.Close()
	assert.Nil(t, err)

	ch := client.OnEvent("newchannel")

	want := 0
	for _, m := range getAmiFixtureCall() {
		if strings.HasPrefix(m, "Event: Newchannel") {
			want++
		}
	}

	have := 0
	for range ch {
		have++
		if have >= want {
			break
		}
	}
	assert.Equal(t, want, have)
}
