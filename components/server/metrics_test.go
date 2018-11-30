package server

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/justinas/alice"

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

func TestMetricsMux(t *testing.T) {
	respRec := httptest.NewRecorder()
	mux := newMetricsMux(alice.New())
	mux.ServeHTTP(respRec, httptest.NewRequest("GET", "/", nil))
	if respRec.Code != 200 {
		t.Error("Expected 200 but got", respRec.Code)
	}
	contentType, ok := respRec.HeaderMap["Content-Type"]
	if !ok {
		t.Fail()
	}
	if !strings.Contains(contentType[0], "application/json") {
		t.Errorf("Incorrect Content-Type: %v", contentType)
	}
}
