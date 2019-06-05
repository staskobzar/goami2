package goami2

import (
	"context"
	"net"
)

func connPipeline(ctx context.Context, conn net.Conn) (<-chan *AMIMsg, error) {
	var err error

	err = readPrompt(ctx, conn)
	if err != nil {
		return nil, err
	}

	chanIn, err := newReader(ctx, conn)
	if err != nil {
		return nil, err
	}
	chanOut, err := msgBuilder(ctx, chanIn)
	if err != nil {
		return nil, err
	}
	return chanOut, nil
}
