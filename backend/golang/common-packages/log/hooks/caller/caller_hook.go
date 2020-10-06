package caller

import (
	"github.com/sirupsen/logrus"
	"github.com/tb/task-logger/backend/golang/common-packages/log/hooks/caller/impl"
	"github.com/gogap/logrus_mate"
)

type CallerHookConfig struct {
}

func init() {
	logrus_mate.RegisterHook("caller", NewCallerHook)
}

func NewCallerHook(options logrus_mate.Options) (hook logrus.Hook, err error) {
	conf := CallerHookConfig{}
	if err = options.ToObject(&conf); err != nil {
		return
	}

	hook, err = impl.NewCallerHook()

	return
}
