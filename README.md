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
      log.Printf("Got event: %s\n", msg.Event())
    }
  }
}
```

# API
TODO: Comming soon...
