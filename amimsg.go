package goami2

import (
	"strings"
)

type MsgType int

const (
	Unknown MsgType = iota
	Event
	Response
)

type AMIMsg struct {
	f map[string]string
	t MsgType
}

func NewAMIMsg() *AMIMsg {
	return &AMIMsg{make(map[string]string), Unknown}
}

func (m *AMIMsg) addFieldOrEOF(str string) (EOF bool) {
	if str == "\r\n" {
		return true
	}
	split := strings.SplitN(str, ":", 2)
	key := strings.ToLower(split[0])
	val := strings.TrimSpace(split[1])
	m.f[key] = val

	switch key {
	case "event":
		m.t = Event
	case "response":
		m.t = Response
	}
	return
}

func (m *AMIMsg) Field(key string) string {
	if val, ok := m.f[key]; ok {
		return val
	}
	return ""
}

func (m *AMIMsg) Type() MsgType {
	return m.t
}
