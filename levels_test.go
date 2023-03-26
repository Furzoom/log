package log

import (
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
		{"unknown", invalidError, -1},
	}

	for _, testCase := range testCases {
		level, err := ParseLevel(testCase.level)
		assert.Equal(t, testCase.expected, level)
		assert.Equal(t, testCase.err, err)
	}
}
