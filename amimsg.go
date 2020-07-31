package goami2

import (
	"encoding/json"
	"strings"
	"sync"
)

type msgType int

const (
	// Unknown AMI Message type
	Unknown msgType = iota
	// Event AMI Message type
	Event
	// Response AMI Message type
	Response
)

// AMIMsg structure of AMI message
type AMIMsg struct {
	f map[string]string
	t msgType
	m sync.Mutex
}

// NewAMIMsg create new AMIMsg
func NewAMIMsg(text string) *AMIMsg {
	msg := &AMIMsg{f: make(map[string]string), t: Unknown}
	for _, str := range strings.Split(text, "\r\n") {
		msg.addField(str)
	}
	return msg
}

func (m *AMIMsg) addField(str string) {
	m.m.Lock()
	defer m.m.Unlock()
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

// Field gets AMI Message field value
func (m *AMIMsg) Field(key string) string {
	key = strings.ToLower(key)
	if val, ok := m.f[key]; ok {
		return val
	}
	return ""
}

// AddField push new field to AMI package.
// If field already set then override the value
func (m *AMIMsg) AddField(key string, value string) {
	m.m.Lock()
	defer m.m.Unlock()
	k := strings.TrimSpace(key)
	k = strings.ToLower(k)
	m.f[k] = value
}

// ActionID gets AMI Message ActionID value or false if not exists
func (m *AMIMsg) ActionID() (string, bool) {
	id := strings.TrimSpace(m.Field("ActionId"))
	if id == "" {
		return "", false
	}
	return id, true
}

// IsResponse returns True is AMI Message is Response
func (m *AMIMsg) IsResponse() bool {
	return m.t == Response
}

// IsEvent returns True is AMI Message is Event
func (m *AMIMsg) IsEvent() bool {
	return m.t == Event
}

// IsSuccess returns True if AMI Message is Response and is "Success"
func (m *AMIMsg) IsSuccess() bool {
	return m.IsResponse() && strings.ToLower(m.Field("response")) == "success"
}

// IsEventList returns True if AMI Message is EventList
func (m *AMIMsg) IsEventList() bool {
	event := m.Field("EventList")
	return strings.TrimSpace(event) != ""
}

// IsEventListStart returns True if AMI Message indicates that EventList starts
func (m *AMIMsg) IsEventListStart() bool {
	if !m.IsSuccess() {
		return false
	}
	if !m.IsEventList() {
		return false
	}
	eventList := m.Field("EventList")
	return strings.ToLower(eventList) == "start"
}

// IsEventListEnd returns True if AMI Message indicates that EventList ends
func (m *AMIMsg) IsEventListEnd() bool {
	eventList := m.Field("EventList")
	return strings.ToLower(eventList) == "complete"
}

// Event gets Event field value from AMI Message
func (m *AMIMsg) Event() (string, bool) {
	event := m.Field("event")
	if strings.TrimSpace(event) == "" {
		return "", false
	}
	return event, true
}

// Message returns value of AMI Message
func (m *AMIMsg) Message() string {
	return m.Field("message")
}

// JSON returns AMI message as JSON string
// if error returns empty sting
func (m *AMIMsg) JSON() string {
	b, err := json.Marshal(m.f)
	if err != nil {
		return ""
	}
	return string(b)
}
