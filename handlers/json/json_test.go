package json_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/json"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0).UTC()
	}
}

func TestHandler_json(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(json.New(&buf))
	log.WithField("user", "Furzoom").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	expected := `{"fields":{"id":"123","user":"Furzoom"},"level":"info","timestamp":"1970-01-01T00:00:00Z","message":"hello"}
{"fields":{},"level":"info","timestamp":"1970-01-01T00:00:00Z","message":"world"}
{"fields":{},"level":"error","timestamp":"1970-01-01T00:00:00Z","message":"boom"}
`

	assert.Equal(t, expected, buf.String())
}
