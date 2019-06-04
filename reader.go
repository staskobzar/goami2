package goami2

import (
	"bufio"
	"context"
	"log"
	"net"
)

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
				if msg.addFieldOrEOF(str) {
					chOut <- msg
					msg = NewAMIMsg()
				}
			}
		}
	}()
	return chOut, nil
}
