package impl

import (
	"github.com/sirupsen/logrus"
	"os"
)

type FileHook struct {
	FileName string
}

func NewFileHook(fileName string) (*FileHook, error) {

	return &FileHook{FileName: fileName}, nil
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	writer, err := os.OpenFile(hook.FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logrus.Println(err.Error())
		return err
	}
	defer writer.Close()
	msg, err := entry.String()
	if err != nil {
		logrus.Println(err.Error())
		return err
	}
	writer.WriteString(msg)
	return nil
}

func (hook *FileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
