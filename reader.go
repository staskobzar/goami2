package goami2

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"log"
	"net"
	"regexp"
	"time"
)

func scanAMIMsg(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if ix := bytes.Index(data, []byte("\r\n\r\n")); ix >= 0 {
		return ix + 4, data[:ix+2], nil
	}

	return 0, nil, nil
}

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
		scanner := bufio.NewScanner(conn)
		scanner.Split(scanAMIMsg)
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					log.Fatalf("Failed to read from connection: %v", err)
				}
				pack := scanner.Text()
				ch <- pack
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
		for {
			select {
			case <-ctx.Done():
				return
			case str := <-chIn:
				chOut <- NewAMIMsg(str)
			}
		}
	}()
	return chOut, nil
}
