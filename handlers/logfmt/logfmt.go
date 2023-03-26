// Package logfmt implements a "logfmt" format handler.
package logfmt

import (
	"io"
	"sync"

	"github.com/furzoom/log"
	"github.com/go-logfmt/logfmt"
)

// Handler implementation
type Handler struct {
	mu  sync.Mutex
	enc *logfmt.Encoder
}

// New handler.
func New(w io.Writer) *Handler {
	return &Handler{
		enc: logfmt.NewEncoder(w),
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if err := h.enc.EncodeKeyval("level", e.Level.String()); err != nil {
		panic(err)
	}

	if err := h.enc.EncodeKeyval("message", e.Message); err != nil {
		panic(err)
	}

	for k, v := range e.Fields {
		if err := h.enc.EncodeKeyval(k, v); err != nil {
			panic(err)
		}
	}

	if err := h.enc.EndRecord(); err != nil {
		panic(err)
	}

	return nil
}
