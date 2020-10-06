package conf

import (
	"testing"
)

func Test_conf(t *testing.T) {

	err := LoadConfigFile()

	if err != nil {
		t.Fatal(err)
	}

}
