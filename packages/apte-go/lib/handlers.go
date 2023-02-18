package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type BaseHandlerContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type ReceiveHandlerContext struct {
	*BaseHandlerContext
}

func (ctx *ReceiveHandlerContext) GetJSONData(base any) (data map[string]any, err error) {
	if base == nil {
		err = errors.New("nil base provided")
	}

	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, base)

	return
}

type StreamHandlerContext struct {
	*BaseHandlerContext
}

func (ctx *StreamHandlerContext) Dispatch(messageType string, data []byte) {
	if data == nil {
		data = []byte{}
	}

	ctx.Writer.Write([]byte(fmt.Sprintf("event: %s\n", messageType)))
	ctx.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", data)))

	if flusher, ok := ctx.Writer.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (ctx *StreamHandlerContext) Continuous(callback func() bool) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for {
			select {
			// On each iteration, check if the request has completed.
			case <-ctx.Request.Context().Done():
				close(done)
				return
			default:
				// Once the callback returns false, complete.
				if !callback() {
					close(done)
					return
				}
			}
		}
	}()
	return done
}

type ReceiveHandler func(ctx *ReceiveHandlerContext)

type StreamHandler func(ctx *StreamHandlerContext)
