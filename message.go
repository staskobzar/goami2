package goami2

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

// Message represents AMI message object
type Message struct {
	h []Header
}

// Header of AMI Message
type Header struct {
	Name  string
	Value string
}

// NewMessage creates new Message
func NewMessage() *Message {
	return &Message{
		// big enough capacity to have best performance and alloc
		h: make([]Header, 0, 32),
	}
}

// NewAction creats action message
func NewAction(name string) *Message {
	msg := NewMessage()
	msg.AddField("Action", name)
	return msg
}

// Field return field value.  Case insensitive field search by name
func (m *Message) Field(key string) string {
	for _, hdr := range m.Headers() {
		if strings.EqualFold(hdr.Name, key) {
			return hdr.Value
		}
	}
	return ""
}

// FieldValues list of values for multiple headers with the same name
func (m *Message) FieldValues(key string) []string {
	hdrs := make([]string, 0)
	for _, h := range m.Headers() {
		if strings.EqualFold(h.Name, key) {
			hdrs = append(hdrs, h.Value)
		}
	}
	return hdrs
}

// Headers list of the Message
func (m *Message) Headers() []Header {
	return m.h
}

// ActionID get AMI message action id as string
func (m *Message) ActionID() string {
	return m.Field("ActionID")
}

// AddActionID create random ID and add ActionID field to the message
func (m *Message) AddActionID() {
	buf := make([]byte, 12)
	_, _ = rand.Read(buf)
	m.SetField("ActionID", fmt.Sprintf("%x", buf))
}

// AddField add field with key and name
func (m *Message) AddField(key, value string) {
	m.h = append(m.h, Header{Name: key, Value: value})
}

// DelField removes field from message by name
func (m *Message) DelField(key string) {
	for i, hdr := range m.Headers() {
		if strings.EqualFold(hdr.Name, key) {
			m.h = slices.Delete(m.h, i, i+1)
			return
		}
	}
}

// SetField get first available field by name and updates
// its value. If field not in list then append new field
func (m *Message) SetField(name, value string) {
	id := -1

	for i, hdr := range m.Headers() {
		if strings.EqualFold(hdr.Name, name) {
			id = i
			break
		}
	}

	if id == -1 {
		m.AddField(name, value)
	} else {
		m.h[id].Value = value
	}
}

// IsEvent returns true if message is event
func (m *Message) IsEvent() bool {
	return len(m.Field("Event")) > 0
}

// IsResponse returns true if message is response
func (m *Message) IsResponse() bool {
	return !m.IsEvent()
}

// IsSuccess returns true if successfully response
func (m *Message) IsSuccess() bool {
	return strings.EqualFold(m.Field("Response"), "success")
}

// Len returns number of headers in the message
func (m *Message) Len() int {
	return len(m.h)
}

func (m *Message) JSON() string {
	push := func(m map[string]string, val string) map[string]string {
		k, v := varsplit(val)
		m[k] = v
		return m
	}
	data := make(map[string]any)
	for _, h := range m.Headers() {
		key := strings.ToLower(h.Name)
		v, ok := data[key]
		if !ok {
			data[key] = h.Value
		}

		switch t := v.(type) {
		case string:
			m := push(map[string]string{}, t)
			data[key] = push(m, h.Value)
		case map[string]string:
			data[key] = push(t, h.Value)
		}
	}

	if res, err := json.Marshal(data); err == nil {
		return string(res)
	}
	return "{}"
}

func FromJSON(input string) (*Message, error) {
	var jsonObj any
	if err := json.Unmarshal([]byte(input), &jsonObj); err != nil {
		return nil, fmt.Errorf("%w: failed to unmarshal from json: %s", ErrAMI, err)
	}
	msg := NewMessage()

	for k, v := range jsonObj.(map[string]any) {
		switch t := v.(type) {
		case map[string]any:
			for nam, val := range t {
				msg.AddField(k, fmt.Sprintf("%v=%v", nam, val))
			}
		default:
			msg.AddField(k, fmt.Sprintf("%v", t))
		}
	}

	return msg, nil
}

// String AMI message as string
func (m *Message) String() string {
	var buf strings.Builder
	for _, h := range m.Headers() {
		buf.WriteString(h.Name)
		buf.WriteString(": ")
		buf.WriteString(h.Value)
		buf.WriteString("\r\n")
	}
	buf.WriteString("\r\n")

	return buf.String()
}

// Byte AMI message as byte array
func (m *Message) Byte() []byte {
	return []byte(m.String())
}

// Var search in AMI message variable fields like Variable or ChanVariable etc
// for a value of the type key=value or just key. If found, returns value as string
// and true. Variable name is case insensative.
func (m *Message) Var(key string) (string, bool) {
	for _, h := range m.Headers() {
		if !strings.Contains(strings.ToUpper(h.Name), "VARIABLE") {
			continue
		}

		k, v := varsplit(h.Value)
		if strings.EqualFold(k, key) {
			return v, true
		}
	}
	return "", false
}

func varsplit(input string) (string, string) {
	v := strings.SplitN(input, "=", 2)
	if len(v) == 2 {
		return v[0], v[1]
	}
	return v[0], ""
}
