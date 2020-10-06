package impl

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

type TraceHook struct {
}

func NewTraceHook() (*TraceHook, error) {

	return &TraceHook{}, nil
}

func (hook *TraceHook) Fire(entry *logrus.Entry) error {
	entry.Data["trace"] = hook.caller()
	return nil
}

func (hook *TraceHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func (hook *TraceHook) caller() string {
	stack := make([]byte, 2048)
	size := runtime.Stack(stack, false)
	trace := string(stack[:size])

	return trace

}
