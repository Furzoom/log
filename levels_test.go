package log

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	testCases := []struct {
		level    string
		err      error
		expected Level
	}{
		{"debug", nil, DebugLevel},
		{"info", nil, InfoLevel},
		{"warn", nil, WarnLevel},
		{"error", nil, ErrorLevel},
		{"fatal", nil, FatalLevel},
		{"DEBUG", nil, DebugLevel},
		{"Warning", nil, WarnLevel},
		{"unknown", ErrInvalidLevel, -1},
	}

	for _, testCase := range testCases {
		level, err := ParseLevel(testCase.level)
		assert.Equal(t, testCase.expected, level)
		assert.Equal(t, testCase.err, err)
	}
}

func TestLevel_MarshalJSON(t *testing.T) {
	e := Entry{
		Level:   InfoLevel,
		Message: "hello",
		Fields:  Fields{},
	}

	expected := `{"fields":{},"level":"info","timestamp":"0001-01-01T00:00:00Z","message":"hello"}`
	b, err := json.Marshal(e)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(b))
}

func TestLevel_UnmarshalJSON(t *testing.T) {
	s := `{"fields":{},"level":"info","timestamp":"0001-01-01T00:00:00Z","message":"hello"}`
	e := new(Entry)

	err := json.Unmarshal([]byte(s), e)
	assert.NoError(t, err)
	assert.Equal(t, InfoLevel, e.Level)
}
