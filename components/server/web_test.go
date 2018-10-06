package server

import (
	"os"
	"testing"
)

func TestDefaultPort(t *testing.T) {
	port := getAppPort()
	if port != 8000 {
		t.Error("Expected port 8000 but got", port)
	}
}

func TestSpecifiedPort(t *testing.T) {
	os.Setenv("GO_APP_PORT", "9999")
	port := getAppPort()
	if port != 9999 {
		t.Error("Expected port 9999 but got", port)
	}
}
