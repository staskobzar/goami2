package goami2

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/textproto"
	"reflect"
	"strings"
	"sync"
)

type Message struct {
	header textproto.MIMEHeader
	mu     sync.RWMutex
}

// NewAction creats action message
func NewAction(name string) *Message {
	h := make(textproto.MIMEHeader)
	h.Add("Action", name)
	a := newMessage(h)
	return a
}

// FromJSON creates message from JSON string
func FromJSON(jstr string) (*Message, error) {
	var jb interface{}
	msg := newMessage(textproto.MIMEHeader{})

	err := json.Unmarshal([]byte(jstr), &jb)
	if err != nil {
		return msg, err
	}

	for k, v := range jb.(map[string]interface{}) {
		ftype := reflect.ValueOf(v)
		switch ftype.Kind() {
		case reflect.Map:
			for vname, vval := range v.(map[string]interface{}) {
				msg.AddField(k, fmt.Sprintf("%s=%v", vname, vval))
			}
		default:
			msg.AddField(k, fmt.Sprintf("%v", v))
		}
	}

	return msg, nil
}

// String return AMI message as string
func (m *Message) String() string {
	buf := m.toByteBuf()
	return buf.String()
}

// String return AMI message as byte array
func (m *Message) Bytes() []byte {
	buf := m.toByteBuf()
	return buf.Bytes()
}

// Field returns first value associated with the given key
func (m *Message) Field(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.header.Get(key)
}

// Var search in AMI message fields Variable and ChanVariable for a value
// of the type key=value or just key. If found, returns value as string
// and true. Variable name is case sensative.
func (m *Message) Var(key string) (string, bool) {
	vars := append(m.FieldValues("variable"), m.FieldValues("chanvariable")...)
	for _, amivar := range vars {
		k, v := varsSplit(amivar)
		if k == key {
			return v, true
		}
	}
	return "", false
}

// FieldValues returns all values associated with the given key
func (m *Message) FieldValues(key string) []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.header.Values(key)
}

// AddField set new field from pair key: value
func (m *Message) AddField(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.header.Add(key, value)
}

// DelField deletes fields associated with key
func (m *Message) DelField(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.header.Del(key)
}

// IsEvent returns true if AMI message is Event
func (m *Message) IsEvent() bool {
	return m.Field("Event") != ""
}

// IsResponse returns true if AMI message is Response
func (m *Message) IsResponse() bool {
	return m.Field("Response") != ""
}

// IsSuccess returns true if AMI message is Response with value success
func (m *Message) IsSuccess() bool {
	return strings.EqualFold(m.Field("Response"), "success")
}

// AddActionID generate random action id and insert into message
func (m *Message) AddActionID() {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err == nil {
		m.AddField("ActionID", fmt.Sprintf("%x", b))
	}
}

// JSON returns AMI message as JSON string
func (m *Message) JSON() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{})

	for key, val := range m.header {
		field := strings.ToLower(key)
		if len(val) == 1 {
			data[field] = val[0]
		} else {
			data[field] = varsMap(val)
		}
	}

	if b, err := json.Marshal(data); err == nil {
		return string(b)
	}
	return ""
}

func varsMap(values []string) map[string]string {
	res := make(map[string]string)
	for _, value := range values {
		k, v := varsSplit(value)
		res[k] = v
	}
	return res
}

func varsSplit(value string) (string, string) {
	split := strings.SplitN(value, "=", 2)
	key := split[0]
	if len(split) == 1 {
		return key, ""
	}
	return key, split[1]
}

// ActionID get AMI message action id as string
func (m *Message) ActionID() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.header.Get("actionid")
}

func newMessage(header textproto.MIMEHeader) *Message {
	m := &Message{}
	m.header = header
	return m
}

func (m *Message) toByteBuf() bytes.Buffer {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var buf bytes.Buffer
	for key := range m.header {
		for _, val := range m.header.Values(key) {
			buf.WriteString(key)
			buf.WriteString(": ")
			buf.WriteString(val)
			buf.WriteString("\r\n")
		}
	}
	buf.WriteString("\r\n")
	return buf
}

func loginMessage(username, password string) *Message {
	action := NewAction("Login")
	action.AddField("Username", username)
	action.AddField("Secret", password)
	return action
}
