# goami2: Simple Asterisk Manager Interface (AMI) library

[![Build Status](https://github.com/staskobzar/goami2/actions/workflows/ci.yml/badge.svg)](https://github.com/staskobzar/goami2/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/staskobzar/goami2)](https://goreportcard.com/report/github.com/staskobzar/goami2)
[![codecov](https://codecov.io/gh/staskobzar/goami2/branch/master/graph/badge.svg)](https://codecov.io/gh/staskobzar/goami2)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/staskobzar/goami2)
[![GitHub release](https://img.shields.io/github/release/staskobzar/goami2.svg)](https://github.com/staskobzar/goami2/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/staskobzar/goami2.svg)](https://pkg.go.dev/github.com/staskobzar/goami2)

Go library that provides interface to Asteris AMI. The library uses given ```net.Conn``` (tcp or tls), login
and starts reading AMI messages from connection and parse them into ```*Message``` object.
Provides simple interface to send messages to AMI.
It uses ```re2c``` to create very fast protocol parser. Run "make bench" for benchmark tests.

Usage example:

```go

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/staskobzar/goami2"
)

func main() {
	conn, err := net.Dial("tcp", "asterisk.host:5038")
	if err != nil {
		log.Fatalln(err)
	}

	// login and create client
	client, err := goami2.NewClient(conn, "admin", "pa55w0rd")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("client connected and logged in")

	// create AMI Action message and send to the asterisk.host
	msg := goami2.NewAction("CoreStatus")
	msg.AddActionID()
	client.Send(msg.Byte())

	defer client.Close() // close connections and all channels
	for {
		select {
		case msg := <-client.AllMessages():
			log.Printf("Got message: %s\n", msg.JSON())
		case err := <-client.Err():
			// terminate on network closed
			if errors.Is(err, goami2.ErrEOF) {
				log.Fatalf("Connection error: %s", err)
			}
			log.Printf("WARNING: AMI client error: %s", err)
		}
	}
}
```

## Docs
See documentation at [https://pkg.go.dev/github.com/staskobzar/goami2](https://pkg.go.dev/github.com/staskobzar/goami2)
