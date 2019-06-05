package goami2

import (
	"bufio"
	"context"
	"errors"
	"log"
	"net"
	"regexp"
	"time"
)

// Reads prompt from AMI and blocks. If passed context has no timeout set
// then default timeout will be 3 sec. Returns nil or error.
func readPrompt(ctx context.Context, conn net.Conn) (err error) {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, time.Second*3)
		defer cancel()
	}
	ch := make(chan error)
	go func() {
		defer close(ch)
		reader := bufio.NewReader(conn)
		prompt, err := reader.ReadSlice('\n')
		if err != nil {
			ch <- err
			return
		}
		re := regexp.MustCompile(`Asterisk Call Manager/(\d+\.\d+(?:\.\d+)?)\r\n`)
		if !re.Match(prompt) {
			ch <- errors.New("Invalid AMI prompt: " + string(prompt))
			return
		}
		ch <- nil
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-ch:
	}
	return
}

func newReader(ctx context.Context, conn net.Conn) (<-chan string, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ch := make(chan string)
	go func() {
		reader := bufio.NewReader(conn)
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					log.Fatalf("Failed to read from connection: %v", err)
				}
				ch <- line
			}
		}
	}()
	return ch, nil
}

func msgBuilder(ctx context.Context, chIn <-chan string) (<-chan *AMIMsg, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	chOut := make(chan *AMIMsg)
	go func() {
		defer close(chOut)
		msg := NewAMIMsg()
		for {
			select {
			case <-ctx.Done():
				return
			case str := <-chIn:
				msg.addField(str)
				if msg.isReady() {
					chOut <- msg
					msg = NewAMIMsg()
				}
			}
		}
	}()
	return chOut, nil
}
