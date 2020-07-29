package goami2

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"regexp"
	"time"
)

const (
	NetChunkSize = 256
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

// read AMI from server, convert to AMIMsg and send downstream
func newReader(ctx context.Context,
	conn net.Conn,
	cherr chan<- error) (<-chan *AMIMsg, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	rpipe, chClose := netReader(ctx, cherr, conn)
	ch := netScanner(ctx, rpipe, cherr, chClose)
	return ch, nil
}

// read from network and send text to  scanner
// also controls network errors
func netReader(ctx context.Context,
	cherr chan<- error,
	conn net.Conn) (*io.PipeReader, chan bool) {
	// internal routing of messages for better network errors handle
	rpipe, wpipe := io.Pipe()
	chClose := make(chan bool)
	go func() {
		defer wpipe.Close()
		defer close(chClose)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, NetChunkSize)
				ln, err := conn.Read(buf)
				if ctx.Err() != nil {
					return // already closed
				}
				if err != nil {
					cherr <- err
					return
				}
				if ln > 0 {
					wpipe.Write(buf[:ln])
				}
			}
		}
	}()
	return rpipe, chClose
}

// scan AMI packets and convert to AMIMsg structure
func netScanner(ctx context.Context,
	rpipe *io.PipeReader,
	cherr chan<- error,
	chClose chan bool) chan *AMIMsg {
	ch := make(chan *AMIMsg)
	go func() {
		scanner := bufio.NewScanner(rpipe)
		scanner.Split(scanAMIMsg)
		defer close(ch)
		defer rpipe.Close()
		for {
			select {
			case <-chClose:
				return
			default:
				scanner.Scan()
				if ctx.Err() != nil {
					return // already closed
				}
				if err := scanner.Err(); err != nil {
					cherr <- err
					continue
				}
				pack := scanner.Text()
				if len(pack) > 0 {
					ch <- NewAMIMsg(pack)
				}
			}
		}
	}()
	return ch
}
