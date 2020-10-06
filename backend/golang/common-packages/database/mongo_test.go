package db

import (
	"github.com/tb/task-logger/backend/golang/common-packages/conf"
	"testing"
)

func Test_Mongo(t *testing.T) {

	mapDbSession, err := InitMongoDbSession()

	if err != nil {
		t.Fatal(err)
	}

	for _, value := range mapDbSession {

		value.Close()
	}

}

func init() {
	conf.LoadConfigFile()
}
