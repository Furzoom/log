// Package discard implements a no-op handler useful for benchmarks
// and test.
package discard

import "github.com/furzoom/log"

// Handler implementation
type Handler struct{}

// New handler.
func New() *Handler {
	return &Handler{}
}

// HandleLog implements log.Handler
func (h *Handler) HandleLog(e *log.Entry) error {
	return nil
}
