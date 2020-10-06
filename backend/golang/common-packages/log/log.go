package tblog

import (
	//"encoding/json"
	log "github.com/sirupsen/logrus"
	_ "github.com/tb/task-logger/backend/golang/common-packages/log/hooks/caller"
	_ "github.com/tb/task-logger/backend/golang/common-packages/log/hooks/file"
	_ "github.com/tb/task-logger/backend/golang/common-packages/log/hooks/mail"
	_ "github.com/tb/task-logger/backend/golang/common-packages/log/hooks/trace"

	"github.com/gogap/logrus_mate"
	_ "github.com/gogap/logrus_mate/hooks/graylog"
	_ "github.com/gogap/logrus_mate/hooks/syslog"
	"io"
	//"io/ioutil"
	"os"
)

var (
	newMate *logrus_mate.LogrusMate
)

type FileIOConfig struct {
	FilePath string `json:"filePath"`
}

func InitLogger() error {

	//Register file writer
	logrus_mate.RegisterWriter("fileio", NewFileIOWriter)

	//Get the config files location from environment variable
	filePath := os.Getenv("TASKLOGGER_CONF_FILE")

	if mateConf, err := logrus_mate.LoadLogrusMateConfig(filePath + "/mate.conf"); err != nil {
		log.Error(err)
		return err
	} else {
		if newMate, err = logrus_mate.NewLogrusMate(mateConf); err != nil {
			log.Error(err)
			return err
		} else {
			newMate.Logger("tb").Info("I am test log in new tb log")
		}
	}
	return nil
}

func GetDefaultLogger() *log.Logger {
	logger := newMate.Logger("tb")
	return logger
}

//Responsible for writing logs to the file
func NewFileIOWriter(options logrus_mate.Options) (writer io.Writer, err error) {
	conf := FileIOConfig{}

	if err = options.ToObject(&conf); err != nil {
		return
	}
	writer, err = os.OpenFile(conf.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	return
}
