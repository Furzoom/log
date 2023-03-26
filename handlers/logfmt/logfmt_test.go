package logfmt

import (
	"os"

	"github.com/furzoom/log"
)

func Example() {
	h := New(os.Stdout)

	e := &log.Entry{
		Message: "upload",
		Fields: log.Fields{
			"name": "Furzoom",
		},
		Level: log.InfoLevel,
	}

	if err := h.HandleLog(e); err != nil {
		panic(err)
	}
	// output: level=info message=upload name=Furzoom\n
}
