// Package text implements a development-friendly textual handler.
package text

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/furzoom/log"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr)

// start time.
var start = time.Now()

// Colors.
const (
	none   = 0
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

const (
	dataTimeLayout = "2006-01-02 15:04:05.999"
)

// Colors mapping.
var Colors = [...]int{
	log.DebugLevel: gray,
	log.InfoLevel:  blue,
	log.WarnLevel:  yellow,
	log.ErrorLevel: red,
	log.FatalLevel: red,
}

// Strings mapping.
var Strings = [...]string{
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
	color := Colors[e.Level]
	level := Strings[e.Level]
	names := e.Fields.Names()

	h.mu.Lock()
	defer h.mu.Unlock()

	ts := time.Since(start) / time.Second
	_, _ = fmt.Fprintf(h.Writer, "%s \033[%dm%6s\033[%dm[%04d] %-25s",
		e.Timestamp.Format(dataTimeLayout), color, level, none, ts, e.Message)

	for _, name := range names {
		_, _ = fmt.Fprintf(h.Writer, " %s=%v", name, e.Fields.Get(name))
	}

	_, _ = fmt.Fprintln(h.Writer)

	return nil

}
