package impl

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type CallerHook struct {
}

func NewCallerHook() (*CallerHook, error) {

	return &CallerHook{}, nil
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	env := os.Getenv("NODE_ENV")
	//Adding caller file no and line no
	entry.Data["caller"] = hook.caller()
	//Adding environment e.g. development, staging, production etc
	entry.Data["environment"] = env
	return nil
}

func (hook *CallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *CallerHook) caller() string {
	file, line := getCallerIgnoringLogMulti(2)

	return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")

}

func getCaller(callDepth int, suffixesToIgnore ...string) (file string, line int) {
	// bump by 1 to ignore the getCaller (this) stack frame
	callDepth++
outer:
	for {
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
			break
		}

		for _, s := range suffixesToIgnore {
			if strings.HasSuffix(file, s) {
				callDepth++
				continue outer
			}
		}
		break
	}
	return
}

//Get list of files to be ignored for reporting file name and line no
func getCallerIgnoringLogMulti(callDepth int) (string, int) {
	// the +1 is to ignore this (getCallerIgnoringLogMulti) frame
	return getCaller(callDepth+1, "logrus/hooks.go", "logrus/entry.go", "logrus/logger.go", "logrus/exported.go", "asm_amd64.s")
}
