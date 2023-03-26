// Package text implements a development-friendly textual handler.
package text

import (
	"fmt"
	"io"
	"sync"

	"github.com/furzoom/log"
)

const (
	none   = 0
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

var colors = [...]int{
	log.DebugLevel: gray,
	log.InfoLevel:  blue,
	log.WarnLevel:  yellow,
	log.ErrorLevel: red,
	log.FatalLevel: red,
}

var strings = [...]string{
	log.DebugLevel: "DEBUG",
	log.InfoLevel:  "INFO",
	log.WarnLevel:  "WARN",
	log.ErrorLevel: "ERROR",
	log.FatalLevel: "FATAL",
}

// Handler implementation.
type Handler struct {
	mu     sync.Mutex
	Writer io.Writer
}

// New handler.
func New(w io.Writer) *Handler {
	return &Handler{
		Writer: w,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	color := colors[e.Level]
	level := strings[e.Level]

	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Fprintf(h.Writer, "\033[%dm%6s\033[%dm %-25s", color, level, none, e.Message)

	for k, v := range e.Fields {
		fmt.Fprintf(h.Writer, " %s=%v", k, v)
	}

	fmt.Fprintln(h.Writer)

	return nil

}
