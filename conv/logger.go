package conv

import "log"

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type defaultLogger struct {
}

var _ Logger = (*defaultLogger)(nil)

var convLogger = Logger(&defaultLogger{})

func WithLogger(l Logger) {
	convLogger = l
}

func (d *defaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

func (d *defaultLogger) Warn(args ...interface{}) {
	log.Println(args...)
}

func (d *defaultLogger) Error(args ...interface{}) {
	log.Println(args...)
}
