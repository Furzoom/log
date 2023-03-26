package log_test

import (
	"errors"
	"testing"

	"github.com/furzoom/log"
	"github.com/furzoom/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

type Pet struct {
	Name string
	Age  int
}

func (p *Pet) Fields() log.Fields {
	return log.Fields{
		"name": p.Name,
		"age":  p.Age,
	}
}

func TestInfo(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	log.Infof("logged in %s", "Container")

	e := h.Entries[0]

	assert.Equal(t, e.Message, "logged in Container")
	assert.Equal(t, e.Level, log.InfoLevel)
}

func TestFielder(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	pet := &Pet{"Container", 3}
	log.WithFields(pet).Info("add pet")

	e := h.Entries[0]
	assert.Equal(t, log.Fields{"name": "Container", "age": 3}, e.Fields)
	assert.Equal(t, e.Level, log.InfoLevel)
	assert.Equal(t, e.Message, "add pet")
}

// Unstructured logging is supported, but not recommended since it is hard to query.
func Example_unstructured() {
	log.Infof("%s logged in", "Furzoom")
}

// Structured logging is supported with fields, and is recommended over the
// formatted message variants.
func Example_structured() {
	log.WithField("name", "Furzoom").Info("logged in")
}

// Errors are passed to WithError(), populating the "error" field.
func Example_errors() {
	err := errors.New("boom")
	log.WithError(err).Error("upload failed")
}

// Multiple fields can be set, via chaining, or WithFields().
func Example_multipleFields() {
	log.WithFields(log.Fields{
		"user": "Furzoom",
		"file": "sloth.png",
		"type": "image/png",
	}).Info("upload")
}

// Trace can be used to simplify logging of start and completion
// events, for example an upload with may fail.
func Example_trace() {
	defer log.Trace("upload").Stop(errors.New("unauthorized"))
}
