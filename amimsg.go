package goami2

import (
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"
)

type msgType int
type fields map[string]string

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
	mu sync.Mutex
	f  atomic.Value // fields
	v  atomic.Value // Variable fields
	cv atomic.Value // ChanVariable fields
	t  msgType
}

func keyForm(key *string) {
	*key = strings.TrimSpace(*key)
	*key = strings.ToLower(*key)
}

func valForm(val *string) {
	*val = strings.TrimSpace(*val)
}

// NewAMIMsg create new AMIMsg
func NewAMIMsg(text string) *AMIMsg {
	msg := &AMIMsg{t: Unknown}
	msg.f.Store(make(fields))
	msg.v.Store(make(fields))
	msg.cv.Store(make(fields))
	for _, str := range strings.Split(text, "\r\n") {
		msg.addField(str)
	}
	return msg
}

func (m *AMIMsg) addField(str string) {
	split := strings.SplitN(str, ":", 2)
	if len(split) < 2 {
		return
	}
	// key := strings.ToLower(split[0])
	key, val := split[0], split[1]
	keyForm(&key)
	valForm(&val)

	m.setType(key)

	switch key {
	case "variable":
		m.storeVariable(val)
	case "chanvariable":
		m.storeChanVariable(val)
	default:
		m.storeField(key, val)
	}

	return
}

func (m *AMIMsg) storeField(key, val string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f := m.f.Load().(fields)
	f[key] = val
	m.f.Store(f)
}

func (m *AMIMsg) storeVariable(fieldValue string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	split := strings.SplitN(fieldValue, "=", 2)
	name, val := split[0], split[1]
	data := m.v.Load().(fields)
	data[name] = val
	m.v.Store(data)
}

func (m *AMIMsg) storeChanVariable(fieldValue string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	split := strings.SplitN(fieldValue, "=", 2)
	name, val := split[0], split[1]
	data := m.cv.Load().(fields)
	data[name] = val
	m.cv.Store(data)
}

func (m *AMIMsg) setType(key string) {
	switch key {
	case "event":
		m.t = Event
	case "response":
		m.t = Response
	}
}

// Field gets AMI Message field value
func (m *AMIMsg) Field(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	keyForm(&key)
	f := m.f.Load().(fields)
	if val, ok := f[key]; ok {
		return val
	}
	return ""
}

// Variable search variable value in Variable fields
func (m *AMIMsg) Variable(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f := m.v.Load().(fields)
	if val, ok := f[key]; ok {
		return val, true
	}
	return "", false
}

// ChanVariable search variable value in ChanVariable fields
func (m *AMIMsg) ChanVariable(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f := m.cv.Load().(fields)
	if val, ok := f[key]; ok {
		return val, true
	}
	return "", false
}

// Var search variable value in Variable and ChanVariable fields
func (m *AMIMsg) Var(key string) (string, bool) {
	if val, ok := m.Variable(key); ok {
		return val, true
	}
	if val, ok := m.ChanVariable(key); ok {
		return val, true
	}
	return "", false
}

// AddField push new field to AMI package.
// If field already set then override the value
func (m *AMIMsg) AddField(key string, value string) {
	keyForm(&key)
	valForm(&value)
	m.storeField(key, value)
}

// DelField deletes field from the AMIMsg
// Returns true when field is found and deleted.
func (m *AMIMsg) DelField(key string) (ret bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	k := strings.TrimSpace(key)
	k = strings.ToLower(k)
	f := m.f.Load().(fields)
	if _, ok := f[k]; ok {
		delete(f, k)
		ret = true
	}
	m.f.Store(f)
	return
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
	data := make(map[string]interface{})
	f := m.f.Load().(fields)
	for k, v := range f {
		data[k] = v
	}

	// add variables
	variables := m.v.Load().(fields)
	if len(variables) > 0 {
		data["variable"] = variables
	}

	// ChanVariable
	chanvars := m.cv.Load().(fields)
	if len(variables) > 0 {
		data["chanvariable"] = chanvars
	}

	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}
