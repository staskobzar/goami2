# goami2: Simple Asterisk Manager Interface (AMI) library

[![Build Status](https://github.com/staskobzar/goami2/actions/workflows/ci.yml/badge.svg)](https://github.com/staskobzar/goami2/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/staskobzar/goami2)](https://goreportcard.com/report/github.com/staskobzar/goami2)
[![codecov](https://codecov.io/gh/staskobzar/goami2/branch/master/graph/badge.svg)](https://codecov.io/gh/staskobzar/goami2)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/staskobzar/goami2)
[![GitHub release](https://img.shields.io/github/release/staskobzar/goami2.svg)](https://github.com/staskobzar/goami2/releases)

Usage example:

```go
import "github.com/staskobzar/goami2"
//.....more code

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5038")
	if err != nil {
		log.Fatalln(err)
	}
	client, err := goami2.NewClient(conn, "admin", "secret")
	if err != nil {
		log.Fatalln(err)
	}

	chAll := client.AllMessages()
	chEvt := client.OnEvent("Newchannel")

	defer client.Close()
	for {
		select {
		case msg := <-chAll:
			log.Printf("Got message: %s\n", msg.JSON())
		case msg := <-chEvt:
			log.Printf("Got new channel: %s\n", msg.Field("Channel"))
		case err := <- client.Error():
			client.Close()
			log.Fatalf("Connection error: %s", err)
		}
	}
}
```

## Usage

```go
var ErrorConnTOUT = errorNew("Connection timeout")
```
ErrorConnTOUT error on connection timeout

```go
var ErrorInvalidPrompt = errorNew("Invalid prompt on AMI server")
```
ErrorInvalidPrompt invalid prompt received from AMI server

```go
var ErrorLogin = errorNew("Login failed")
```
ErrorLogin AMI server login failed

```go
var ErrorNet = errorNew("Network error")
```
ErrorNet networking errors

```go
var NetTOUT = time.Second * 3
```
NetTOUT timeout for net connections and requests

#### type Client

```go
type Client struct {
}
```

Client structure

#### func  NewClient

```go
func NewClient(conn net.Conn, username, password string) (*Client, error)
```
NewClient connects to Asterisk using conn and tries to login using username and
password. Creates new AMI client and returns client or error

#### func (*Client) Action

```go
func (c *Client) Action(msg *Message) bool
```
Action sends AMI message to Asterisk server and returns send-only response
channel on nil

#### func (*Client) AllMessages

```go
func (c *Client) AllMessages() <-chan *Message
```
AllMessages subscribes to any AMI message received from Asterisk server returns
send-only channel or nil

#### func (*Client) Close

```go
func (c *Client) Close()
```
Close client and destroy all subscriptions to events and action responses

#### func (*Client) Err

```go
func (c *Client) Err() <-chan error
```
Err retuns channel for error and signals that client should be probably
restarted

#### func (*Client) OnEvent

```go
func (c *Client) OnEvent(name string) <-chan *Message
```
OnEvent subscribes by event by name (case insensative) and returns send-only
channel or nil

#### type Message

```go
type Message struct {
}
```


#### func  FromJSON

```go
func FromJSON(jstr string) (*Message, error)
```
FromJSON creates message from JSON string

#### func  NewAction

```go
func NewAction(name string) *Message
```
NewAction creats action message

#### func (*Message) ActionID

```go
func (m *Message) ActionID() string
```
ActionID get AMI message action id as string

#### func (*Message) AddActionID

```go
func (m *Message) AddActionID()
```
AddActionID generate random action id and insert into message

#### func (*Message) AddField

```go
func (m *Message) AddField(key, value string)
```
AddField set new field from pair key: value

#### func (*Message) Bytes

```go
func (m *Message) Bytes() []byte
```
String return AMI message as byte array

#### func (*Message) DelField

```go
func (m *Message) DelField(key string)
```
DelField deletes fields associated with key

#### func (*Message) Field

```go
func (m *Message) Field(key string) string
```
Field returns first value associated with the given key

#### func (*Message) FieldValues

```go
func (m *Message) FieldValues(key string) []string
```
FieldValues returns all values associated with the given key

#### func (*Message) IsEvent

```go
func (m *Message) IsEvent() bool
```
IsEvent returns true if AMI message is Event

#### func (*Message) IsResponse

```go
func (m *Message) IsResponse() bool
```
IsResponse returns true if AMI message is Response

#### func (*Message) IsSuccess

```go
func (m *Message) IsSuccess() bool
```
IsSuccess returns true if AMI message is Response with value success

#### func (*Message) JSON

```go
func (m *Message) JSON() string
```
JSON returns AMI message as JSON string

#### func (*Message) String

```go
func (m *Message) String() string
```
String return AMI message as string

#### func (*Message) Var

```go
func (m *Message) Var(key string) (string, bool)
```
Var search in AMI message fields Variable and ChanVariable for a value of the
type key=value or just key. If found, returns value as string and true. Variable
name is case sensative.
