package trace

import (
	"github.com/sirupsen/logrus"
	"github.com/tb/task-logger/backend/golang/common-packages/log/hooks/trace/impl"
	"github.com/gogap/logrus_mate"
)

type TraceHookConfig struct {
}

func init() {
	logrus_mate.RegisterHook("trace", NewTraceHook)
}

func NewTraceHook(options logrus_mate.Options) (hook logrus.Hook, err error) {
	conf := TraceHookConfig{}
	if err = options.ToObject(&conf); err != nil {
		return
	}

	hook, err = impl.NewTraceHook()

	return
}
