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
	dataTimeLayout = "2006-01-02 15:04:05.000"
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
	isTTY  bool
}

// New handler.
func New(w io.Writer) *Handler {
	h := &Handler{}
	if f, ok := w.(*os.File); ok {
		fi, err := f.Stat()
		if err == nil && (fi.Mode()&os.ModeCharDevice == os.ModeCharDevice) {
			h.isTTY = true
		}
	}

	h.Writer = w

	return h
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	color := Colors[e.Level]
	level := Strings[e.Level]
	names := e.Fields.Names()
	frame := e.Frame

	h.mu.Lock()
	defer h.mu.Unlock()

	ts := time.Since(start) / time.Second
	if h.isTTY {
		_, _ = fmt.Fprintf(h.Writer, "%s \033[%dm%6s\033[%dm[%04d] %b:%d %n() %-25s",
			e.Timestamp.Format(dataTimeLayout), color, level, none, ts, frame, frame, frame, e.Message)
	} else {
		_, _ = fmt.Fprintf(h.Writer, "%s %6s[%04d] %b:%d %n() %-25s",
			e.Timestamp.Format(dataTimeLayout), level, ts, frame, frame, frame, e.Message)
	}

	for _, name := range names {
		_, _ = fmt.Fprintf(h.Writer, " %s=%v", name, e.Fields.Get(name))
	}

	_, _ = fmt.Fprintln(h.Writer)

	return nil

}
