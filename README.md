# goami2: Simple Asterisk Manager Interface (AMI) library

[![Build Status](https://travis-ci.org/staskobzar/goami2.svg?branch=master)](https://travis-ci.org/staskobzar/goami2)
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
	client, err := goami2.NewClient(conn)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Login("username", "secret")
	if err != nil {
		log.Fatalln(err)
	}
	ch, err := client.AnyEvent()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	for {
		select {
		case msg := <-ch:
			if event, ok := msg.Event(); ok {
				log.Printf("Got event: %s\n", event)
			}
		case err := <- client.Error():
			client.Close()
			log.Fatalf("Connection error: %s", err)
		}
	}
}
```

# API goami2
```go
    import "github.com/staskobzar/goami2"
```

## Usage

>### type AMIMsg
>
>AMIMsg structure of AMI message
---
#### func  NewAMIMsg

```go
func NewAMIMsg(text string) *AMIMsg
```
NewAMIMsg create new AMIMsg

#### func (*AMIMsg) ActionID

```go
func (m *AMIMsg) ActionID() (string, bool)
```
ActionID gets AMI Message ActionID value or false if not exists

#### func (*AMIMsg) AddField

```go
func (m *AMIMsg) AddField(key string, value string)
```
AddField push new field to AMI package. If field already set then override the
value

#### func (*AMIMsg) ChanVariable

```go
func (m *AMIMsg) ChanVariable(key string) (string, bool)
```
ChanVariable search variable value in ChanVariable fields

#### func (*AMIMsg) DelField

```go
func (m *AMIMsg) DelField(key string) (ret bool)
```
DelField deletes field from the AMIMsg Returns true when field is found and
deleted.

#### func (*AMIMsg) Event

```go
func (m *AMIMsg) Event() (string, bool)
```
Event gets Event field value from AMI Message

#### func (*AMIMsg) Field

```go
func (m *AMIMsg) Field(key string) string
```
Field gets AMI Message field value

#### func (*AMIMsg) IsEvent

```go
func (m *AMIMsg) IsEvent() bool
```
IsEvent returns True is AMI Message is Event

#### func (*AMIMsg) IsEventList

```go
func (m *AMIMsg) IsEventList() bool
```
IsEventList returns True if AMI Message is EventList

#### func (*AMIMsg) IsEventListEnd

```go
func (m *AMIMsg) IsEventListEnd() bool
```
IsEventListEnd returns True if AMI Message indicates that EventList ends

#### func (*AMIMsg) IsEventListStart

```go
func (m *AMIMsg) IsEventListStart() bool
```
IsEventListStart returns True if AMI Message indicates that EventList starts

#### func (*AMIMsg) IsResponse

```go
func (m *AMIMsg) IsResponse() bool
```
IsResponse returns True is AMI Message is Response

#### func (*AMIMsg) IsSuccess

```go
func (m *AMIMsg) IsSuccess() bool
```
IsSuccess returns True if AMI Message is Response and is "Success"

#### func (*AMIMsg) JSON

```go
func (m *AMIMsg) JSON() string
```
JSON returns AMI message as JSON string if error returns empty sting

#### func (*AMIMsg) Message

```go
func (m *AMIMsg) Message() string
```
Message returns value of AMI Message

#### func (*AMIMsg) Var

```go
func (m *AMIMsg) Var(key string) (string, bool)
```
Var search variable value in Variable and ChanVariable fields

#### func (*AMIMsg) Variable

```go
func (m *AMIMsg) Variable(key string) (string, bool)
```
Variable search variable value in Variable fields


>### type Action
>
>Action structure
---
#### func  ActionFromJSON

```go
func ActionFromJSON(source string) (*Action, error)
```
ActionFromJSON convert JSON string to action structure

#### func  NewAction

```go
func NewAction() *Action
```
NewAction creats action

#### func (*Action) ActionID

```go
func (a *Action) ActionID() string
```
ActionID returns AMI ActionID value

#### func (*Action) Field

```go
func (a *Action) Field(header, value string)
```
Field gets Action message

#### func (*Action) Login

```go
func (a *Action) Login(user, password string) []byte
```
Login gets Login AMI Action as []bytes

#### func (*Action) Message

```go
func (a *Action) Message() []byte
```
Message convert Action to []bytes

#### func (*Action) New

```go
func (a *Action) New(action string)
```
New resets buffer and set new Action

>### type Client
>
>Client structure
---
#### func  NewClient

```go
func NewClient(conn net.Conn) (*Client, error)
```
NewClient creates new AMI client and returns client or error

#### func (*Client) Action

```go
func (c *Client) Action(actionName string, fields map[string]string) (chan *AMIMsg, error)
```
Action send to Asterisk MI. ActionID will be generated.

#### func (*Client) AnyEvent

```go
func (c *Client) AnyEvent() (chan *AMIMsg, error)
```
AnyEvent provides channel for any AMI events received

#### func (*Client) Close

```go
func (c *Client) Close()
```
Close client and stop all routines

#### func (*Client) Error

```go
func (c *Client) Error() <-chan error
```
Error returns channel with connection errors

#### func (*Client) Login

```go
func (c *Client) Login(user, pass string) error
```
Login action. Blocking and waits response.

#### func (*Client) OnEvent

```go
func (c *Client) OnEvent(event string) (chan *AMIMsg, error)
```
OnEvent provides channel for specific event. Returns error if listener for event
already created.

#### func (*Client) Send

```go
func (c *Client) Send(msg *Action) (chan *AMIMsg, error)
```
Send AMIMsg to AMI server
