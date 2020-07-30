package goami2

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
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

// ActionFromJSON convert JSON string to action structure
func ActionFromJSON(source string) (*Action, error) {
	var jb interface{}
	action := NewAction()
	err := json.Unmarshal([]byte(source), &jb)
	if err != nil {
		return nil, err
	}

	for k, v := range jb.(map[string]interface{}) {
		action.Field(strings.Title(k), v.(string))
		if strings.EqualFold(k, "actionid") {
			action.aid = v.(string)
		}
	}

	if len(strings.TrimSpace(action.aid)) == 0 {
		action.writeActionID()
	}

	return action, nil
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
