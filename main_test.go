package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	if sum != 3 {
		t.Error("Expected 3 but got ", sum)
	}
}
