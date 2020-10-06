package file

import (
	"github.com/sirupsen/logrus"
	"github.com/tb/task-logger/backend/golang/common-packages/log/hooks/file/impl"
	"github.com/gogap/logrus_mate"
)

type FileHookConfig struct {
	FileName string `json:"fileName"`
}

func init() {
	logrus_mate.RegisterHook("file", NewFileHook)
}

func NewFileHook(options logrus_mate.Options) (hook logrus.Hook, err error) {
	conf := FileHookConfig{}
	if err = options.ToObject(&conf); err != nil {
		return
	}

	hook, err = impl.NewFileHook(conf.FileName)

	return
}
