package server

import (
	"testing"

	"github.com/spf13/viper"
)

func TestDefaultAppPort(t *testing.T) {
	viper.Reset()
	port := getAppPort()
	if port != 8000 {
		t.Error("Expected port 8000 but got", port)
	}
}

func TestSpecifiedAppPort(t *testing.T) {
	viper.Reset()
	viper.Set("app_port", "9999")
	port := getAppPort()
	if port != 9999 {
		t.Error("Expected port 9999 but got", port)
	}
}

func TestSpecifiedAppPortTooBig(t *testing.T) {
	viper.Reset()
	viper.Set("app_port", "65536")
	port := getAppPort()
	if port != 8000 {
		t.Error("Expected port 8000 but got", port)
	}
}
