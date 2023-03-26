// Package memory implements an in-memory handler useful for testing,
// as the entries can be accessed after writes.
package memory

import "github.com/furzoom/log"

// Handler implementation.
type Handler struct {
	Entries []*log.Entry
}

// New Handlers.
func New() *Handler {
	return &Handler{}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	h.Entries = append(h.Entries, e)
	return nil
}
