package goami2

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

const CRLF = "\r\n"

type Action struct {
	bytes.Buffer
	aid string
}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) New(action string) {
	a.aid = ""
	a.Reset()
	a.WriteString("Action: " + action + CRLF)
	a.writeActionId()
}

func (a *Action) writeActionId() {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return
	}
	a.aid = fmt.Sprintf("%x", b)
	a.WriteString(fmt.Sprintf("ActionID: %x\r\n", b))
}

func (a *Action) ActionId() string {
	return a.aid
}

func (a *Action) Field(header, value string) {
	a.WriteString(header + ": " + value + CRLF)
}

func (a *Action) Message() []byte {
	a.WriteString(CRLF)
	return a.Bytes()
}

func (a *Action) Login(user, password string) []byte {
	a.New("Login")
	a.Field("Username", user)
	a.Field("Secret", password)
	return a.Message()
}
