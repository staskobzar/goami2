# goami2: Simple Asterisk Manager Interface (AMI) library

[![Build Status](https://travis-ci.org/staskobzar/goami2.svg?branch=master)](https://travis-ci.org/staskobzar/goami2)
[![Go Report Card](https://goreportcard.com/badge/github.com/staskobzar/goami2)](https://goreportcard.com/report/github.com/staskobzar/goami2)
[![codecov](https://codecov.io/gh/staskobzar/goami2/branch/master/graph/badge.svg)](https://codecov.io/gh/staskobzar/goami2)

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

# API
TODO: Improve docs
```
package goami2 // import "github.com/staskobzar/goami2"
```
## Client
```
type Client struct {
	// Has unexported fields.
}
    Client structure

func NewClient(conn net.Conn) (*Client, error)
    NewClient creates new AMI client and returns client or error

func (c *Client) Action(action string, fields map[string]string) (chan *AMIMsg, error)
    Action send to Asterisk MI.

func (c *Client) AnyEvent() (chan *AMIMsg, error)
    AnyEvent provides channel for any AMI events received

func (c *Client) Close()
    Close client and stop all routines

func (c *Client) Login(user, pass string) error
    Login action. Blocking and waits response.

func (c *Client) OnEvent(event string) (chan *AMIMsg, error)
    OnEvent provides channel for specific event. Returns error if listener for
    event already created.

func (c *Client) Error() <-chan error
    Error returns channel with connection errors

```

## AMIMsg
```
type AMIMsg struct {
	// Has unexported fields.
}
    AMIMsg structure of AMI message

func NewAMIMsg(text string) *AMIMsg
    NewAMIMsg create new AMIMsg

func (m *AMIMsg) ActionID() (string, bool)
    ActionID gets AMI Message ActionID value or false if not exists

func (m *AMIMsg) Event() (string, bool)
    Event gets Event field value from AMI Message

func (m *AMIMsg) Field(key string) string
    Field gets AMI Message field value

func (m *AMIMsg) IsEvent() bool
    IsEvent returns True is AMI Message is Event

func (m *AMIMsg) IsEventList() bool
    IsEventList returns True if AMI Message is EventList

func (m *AMIMsg) IsEventListEnd() bool
    IsEventListEnd returns True if AMI Message indicates that EventList ends

func (m *AMIMsg) IsEventListStart() bool
    IsEventListStart returns True if AMI Message indicates that EventList starts

func (m *AMIMsg) IsResponse() bool
    IsResponse returns True is AMI Message is Response

func (m *AMIMsg) IsSuccess() bool
    IsSuccess returns True if AMI Message is Response and is "Success"

func (m *AMIMsg) Message() string
    Message returns value of AMI Message
```

## Action
```
type Action struct {
	bytes.Buffer
	// Has unexported fields.
}
    Action structure

func NewAction() *Action
    NewAction creats action

func (a *Action) ActionID() string
    ActionID returns AMI ActionID value

func (a *Action) Field(header, value string)
    Field gets Action message

func (a *Action) Login(user, password string) []byte
    Login gets Login AMI Action as []bytes

func (a *Action) Message() []byte
    Message convert Action to []bytes

func (a *Action) New(action string)
    New resets buffer and set new Action
```
