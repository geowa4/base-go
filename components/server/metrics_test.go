package server

import (
	"testing"

	"github.com/spf13/viper"
)

func TestMetricsDefaultPort(t *testing.T) {
	viper.Reset()
	port := getMetricsPort()
	if port != 8001 {
		t.Error("Expected port 8001 but got", port)
	}
}

func TestSpecifiedMetricsPort(t *testing.T) {
	viper.Reset()
	viper.Set("metrics_port", "9999")
	port := getMetricsPort()
	if port != 9999 {
		t.Error("Expected port 9999 but got", port)
	}
}

func TestSpecifiedMetricsPortTooBig(t *testing.T) {
	viper.Reset()
	viper.Set("metrics_port", "65536")
	port := getMetricsPort()
	if port != 8001 {
		t.Error("Expected port 8001 but got", port)
	}
}
