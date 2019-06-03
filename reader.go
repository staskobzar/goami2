package goami2

import (
	"bufio"
	"context"
	"fmt"
	"net"
)

func NewReader(ctx context.Context, conn net.Conn) (<-chan interface{}, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ch := make(chan interface{})
	go func() {
		reader := bufio.NewReader(conn)
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closed context")
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					// TODO: handle reader error
				}
				ch <- line
			}
		}
	}()
	return ch, nil
}
