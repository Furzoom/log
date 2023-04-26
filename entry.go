package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// Assert interface compliance.
var _ Interface = (*Entry)(nil)

// Now returns the current time.
var Now = time.Now

// Entry represents a single log entry.
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"fields"`
	Level     Level     `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	start     time.Time
	fields    []Fields
	Frame     Frame `json:"-"`
	depth     int
}

// NewEntry returns a new entry for `log`.
func NewEntry(log *Logger) *Entry {
	return &Entry{
		Logger: log,
	}
}

// WithFields returns a new Entry with `fields` set.
func (e *Entry) WithFields(fields Fielder) *Entry {
	var f []Fields
	f = append(f, e.fields...)
	f = append(f, fields.Fields())

	return &Entry{
		Logger: e.Logger,
		fields: f,
	}
}

// WithField returns a new Entry with the `key` and `value` set.
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(Fields{key: value})
}

// WithError returns a new Entry with the "error" set to `err`.
//
// The given error may implement .Fielder, if it does the method
// will add all its `.Fields()` into the returned entry.
func (e *Entry) WithError(err error) *Entry {
	if err == nil {
		return e
	}

	ctx := e.WithField("error", err.Error())

	if f, ok := err.(Fielder); ok {
		ctx = ctx.WithFields(f.Fields())
	}

	return ctx
}

// WithCaller includes the "filename" and "line" fields.
func (e *Entry) WithCaller() *Entry {
	if _, file, line, ok := runtime.Caller(1); ok {
		return e.WithFields(Fields{
			"filename": file,
			"line":     line,
		})
	}
	return e
}

// Debug level message.
func (e *Entry) Debug(msg string) {
	e.log(DebugLevel, msg)
}

// Info level message.
func (e *Entry) Info(msg string) {
	e.log(InfoLevel, msg)
}

// Warn level message.
func (e *Entry) Warn(msg string) {
	e.log(WarnLevel, msg)
}

// Error level message.
func (e *Entry) Error(msg string) {
	e.log(ErrorLevel, msg)
}

// Fatal level message, followed by an exit.
func (e *Entry) Fatal(msg string) {
	e.log(FatalLevel, msg)
	os.Exit(1)
}

// Debugf level formatted message.
func (e *Entry) Debugf(msg string, v ...interface{}) {
	e.depth += 1
	e.Debug(fmt.Sprintf(msg, v...))
}

// Infof level formatted message.
func (e *Entry) Infof(msg string, v ...interface{}) {
	e.depth += 1
	e.Info(fmt.Sprintf(msg, v...))
}

// Warnf level formatted message.
func (e *Entry) Warnf(msg string, v ...interface{}) {
	e.depth += 1
	e.Warn(fmt.Sprintf(msg, v...))
}

// Errorf level formatted message.
func (e *Entry) Errorf(msg string, v ...interface{}) {
	e.depth += 1
	e.Error(fmt.Sprintf(msg, v...))
}

// Fatalf level formatted message, followed by an exit.
func (e *Entry) Fatalf(msg string, v ...interface{}) {
	e.depth += 1
	e.Fatal(fmt.Sprintf(msg, v...))
}

func (e *Entry) log(level Level, msg string) {
	e.Frame = NewFrame(e.depth)
	e.Logger.log(level, e, msg)
	if level == FatalLevel {
		os.Exit(1)
	}
}

// Trace returns a new Entry with a stop method to fire off
// a corresponding completion log, useful with defer.
func (e *Entry) Trace(msg string) *Entry {
	e.Info(msg)
	v := e.WithFields(e.Fields)
	v.Message = msg
	v.start = time.Now()
	return v
}

// Stop should be used with Trace, to fire off the completion
// message. When an `err` is passed the "error" field is set,
// and the log level is error.
func (e *Entry) Stop(err *error) {
	duration := time.Since(e.start).Milliseconds()
	if err == nil || *err == nil {
		e.WithField("duration", duration).Info(e.Message)
	} else {
		e.WithField("duration", duration).WithError(*err).Error(e.Message)
	}
}

// mergeFields returns the fields list collapsed into a single map.
func (e *Entry) mergeFields() Fields {
	f := Fields{}

	for _, fields := range e.fields {
		for k, v := range fields {
			f[k] = v
		}
	}

	return f
}

// finalize returns a copy of the Entry with fields merged.
func (e *Entry) finalize(level Level, msg string, depth int) *Entry {
	return &Entry{
		Logger:    e.Logger,
		Fields:    e.mergeFields(),
		Level:     level,
		Message:   msg,
		Timestamp: Now(),
		Frame:     e.Frame,
		depth:     depth,
	}
}
