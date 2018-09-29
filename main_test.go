package main

import "testing"

func TestDefaultPort(t *testing.T) {
	addr := getAddr("")
	if addr != ":8000" {
		t.Error("Expected port 8000 but got", addr)
	}
}

func TestSpecifiedPort(t *testing.T) {
	addr := getAddr("9999")
	if addr != ":9999" {
		t.Error("Expected port 9999 but got", addr)
	}
}
