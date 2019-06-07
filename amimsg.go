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
	f   map[string]string
	t   MsgType
	EOF bool
}

func NewAMIMsg() *AMIMsg {
	return &AMIMsg{make(map[string]string), Unknown, false}
}

func (m *AMIMsg) addField(str string) {
	m.EOF = str == "\r\n"
	if m.EOF {
		return
	}
	split := strings.SplitN(str, ":", 2)
	if len(split) < 2 {
		return
	}
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
	key = strings.ToLower(key)
	if val, ok := m.f[key]; ok {
		return val
	}
	return ""
}

func (m *AMIMsg) isReady() bool {
	return m.EOF
}

func (m *AMIMsg) Type() MsgType {
	return m.t
}

func (m *AMIMsg) ActionId() (string, bool) {
	id := strings.TrimSpace(m.Field("ActionId"))
	if id == "" {
		return "", false
	} else {
		return id, true
	}
}

func (m *AMIMsg) IsResponse() bool {
	return m.t == Response
}

func (m *AMIMsg) IsEvent() bool {
	return m.t == Event
}

func (m *AMIMsg) IsSuccess() bool {
	return m.IsResponse() && strings.ToLower(m.Field("response")) == "success"
}

func (m *AMIMsg) IsEventList() bool {
	event := m.Field("EventList")
	return strings.TrimSpace(event) != ""
}

func (m *AMIMsg) Event() (string, bool) {
	event := m.Field("event")
	if strings.TrimSpace(event) == "" {
		return "", false
	}
	return event, true
}

func (m *AMIMsg) Message() string {
	return m.Field("message")
}
