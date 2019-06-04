package goami2

import (
	"bufio"
	"context"
	"net"
)

type readerStream struct {
	str string
	err error
}

func newReader(ctx context.Context, conn net.Conn) (<-chan readerStream, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ch := make(chan readerStream)
	go func() {
		reader := bufio.NewReader(conn)
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, err := reader.ReadString('\n')
				ch <- readerStream{line, err}
			}
		}
	}()
	return ch, nil
}
