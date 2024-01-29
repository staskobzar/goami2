// Package goami2 library for Asterisk Manager Interface
package goami2

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"
)

const (
	bufSize      = 1 << 12
	promptPrefix = "Asterisk Call Manager/"
	netTimeout   = 2 * time.Second       // default timeout for network read/write
	chanGiveup   = 10 * time.Millisecond // timeout to giveup sending to a channel
)

var (
	Error   = errors.New("goami2")
	ErrConn = fmt.Errorf("%w: net conn", Error)
	ErrAMI  = fmt.Errorf("%w: AMI proto", Error)
	ErrEOF  = fmt.Errorf("%w: terminated", Error)
)

// NewClient creates client. It is using NewClientWithContext in the background
// with a bogus context. For better context control use NewClientWithContext function.
func NewClient(conn net.Conn, username, password string) (*Client, error) {
	return NewClientWithContext(context.Background(), conn, username, password)
}

// NewClientWithContext creates client with provided connection net.Conn and login into
// AMI server. It returns error if fials to login. Runs internal connection loop and
// provides AMI messages via AllMessages and error via Err methods.
func NewClientWithContext(ctx context.Context, conn net.Conn, username, password string) (*Client, error) {
	cl := makeClient(conn)
	if err := cl.login(username, password); err != nil {
		return nil, err
	}

	go cl.loop(ctx)

	return cl, nil
}
