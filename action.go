package goami2

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

const crlf = "\r\n"

// Action structure
type Action struct {
	bytes.Buffer
	aid string
}

// NewAction creats action
func NewAction() *Action {
	return &Action{}
}

// New resets buffer and set new Action
func (a *Action) New(action string) {
	a.aid = ""
	a.Reset()
	a.WriteString("Action: " + action + crlf)
	a.writeActionID()
}

func (a *Action) writeActionID() {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return
	}
	a.aid = fmt.Sprintf("%x", b)
	a.WriteString(fmt.Sprintf("ActionID: %x\r\n", b))
}

// ActionID returns AMI ActionID value
func (a *Action) ActionID() string {
	return a.aid
}

// Field gets Action message
func (a *Action) Field(header, value string) {
	a.WriteString(header + ": " + value + crlf)
}

// Message convert Action to []bytes
func (a *Action) Message() []byte {
	a.WriteString(crlf)
	return a.Bytes()
}

// Login gets Login AMI Action as []bytes
func (a *Action) Login(user, password string) []byte {
	a.New("Login")
	a.Field("Username", user)
	a.Field("Secret", password)
	return a.Message()
}
