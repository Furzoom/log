package level_test

import (
	"testing"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/level"
	"github.com/furzoom/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestHandler_level(t *testing.T) {
	h := memory.New()

	ctx := &log.Logger{
		Handler: level.New(h, log.ErrorLevel),
		Level:   log.InfoLevel,
	}

	ctx.Info("hello")
	ctx.Warn("world")
	ctx.Error("boom")

	assert.Len(t, h.Entries, 1)
	assert.Equal(t, "boom", h.Entries[0].Message)
}
