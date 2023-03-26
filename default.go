package log

import (
	"bytes"
	"fmt"
	"log"
)

// handleStdLog outputs to the stdlib log.
func handleStdLog(e *Entry) error {
	names := e.Fields.Names()
	level := levelNames[e.Level]
	var b bytes.Buffer

	_, _ = fmt.Fprintf(&b, "%5s %s |", level, e.Message)

	for _, name := range names {
		_, _ = fmt.Fprintf(&b, " %s=%v", name, e.Fields.Get(name))
	}

	log.Println(b.String())

	return nil
}
