package logfmt_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/logfmt"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0).UTC()
	}
}

func TestHandler_logfmt(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(logfmt.New(&buf))
	log.WithField("user", "Furzoom").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	expected := `timestamp=1970-01-01T00:00:00Z level=info message=hello id=123 user=Furzoom
timestamp=1970-01-01T00:00:00Z level=info message=world
timestamp=1970-01-01T00:00:00Z level=error message=boom
`

	assert.Equal(t, expected, buf.String())
}

type discard struct{}

func (d *discard) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func BenchmarkHandler_logfmt(b *testing.B) {
	d := &discard{}
	log.SetHandler(logfmt.New(d))
	ctx := log.WithField("user", "Furzoom").WithField("id", "123")

	for i := 0; i < b.N; i++ {
		ctx.Info("hello")
	}
}

func Example() {
	h := logfmt.New(os.Stdout)

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
	// output: timestamp=0001-01-01T00:00:00Z level=info message=upload name=Furzoom
}
