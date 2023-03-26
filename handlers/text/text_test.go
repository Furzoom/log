package text

import (
	"os"

	"github.com/furzoom/log"
)

func Example() {
	h := New(os.Stdout)

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
