package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	conf, err := ReadConfig("../config.conf")
	if err != nil {
		t.Fatalf("read config error %v", err)

	}

	fmt.Println(conf)
}
