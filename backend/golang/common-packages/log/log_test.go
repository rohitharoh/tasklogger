package tblog

import (
	"testing"
)

func Test_Default_Logger(t *testing.T) {
	err := InitLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger := GetDefaultLogger()

	if logger == nil {
		t.Error("Could not get default logger")
	}

}
