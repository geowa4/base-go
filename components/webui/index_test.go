package webui

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	respRec := httptest.NewRecorder()
	mux := NewStaticMux()
	mux.ServeHTTP(respRec, httptest.NewRequest("GET", "/", nil))
	if respRec.Code != 200 {
		t.Error("Expected 200 but got", respRec.Code)
	}
	body := respRec.Body.String()
	if !strings.Contains(body, "Hello world!") {
		t.Error("Expected greeting but got", body)
	}
}
