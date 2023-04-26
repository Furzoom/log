package text_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/text"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0)
	}
}

func TestHandler_text(t *testing.T) {
	var buf bytes.Buffer

	log.SetHandler(text.New(&buf))
	log.WithField("user", "Furzoom").WithField("id", "123").Info("hello")
	log.WithField("user", "Furzoom").Info("world")
	log.WithField("user", "Furzoom").Error("boom")

	expected := `1970-01-01 08:00:00.000   INFO[0000] text_test.go:24 TestHandler_text() hello                     id=123 user=Furzoom
1970-01-01 08:00:00.000   INFO[0000] text_test.go:25 TestHandler_text() world                     user=Furzoom
1970-01-01 08:00:00.000  ERROR[0000] text_test.go:26 TestHandler_text() boom                      user=Furzoom
`

	assert.Equal(t, expected, buf.String())
}

func Example() {
	h := text.New(os.Stdout)

	testCases := []struct {
		level  log.Level
		msg    string
		fields log.Fields
	}{
		{log.DebugLevel, "debug message", log.Fields{"name": "furzoom"}},
		{log.InfoLevel, "info message", log.Fields{"user": "furzoom", "age": 17}},
		{log.WarnLevel, "warn message", log.Fields{}},
		{log.ErrorLevel, "error messaeg", log.Fields{"context": "stuff"}},
	}

	for _, testCase := range testCases {
		e := &log.Entry{
			Level:   testCase.level,
			Message: testCase.msg,
			Fields:  testCase.fields,
		}

		if err := h.HandleLog(e); err != nil {
			panic(err)
		}
	}
}
