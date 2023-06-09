package multi_test

import (
	"testing"
	"time"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/memory"
	"github.com/furzoom/log/handlers/multi"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Now = func() time.Time {
		return time.Unix(0, 0)
	}
}

func TestHandler_multi(t *testing.T) {
	a := memory.New()
	b := memory.New()

	log.SetHandler(multi.New(a, b))
	log.WithField("user", "Furzoom").WithField("id", "123").Info("hello")
	log.Info("world")
	log.Error("boom")

	assert.Len(t, a.Entries, 3)
	assert.Len(t, b.Entries, 3)
}
