package foos

import "testing"

func TestConvertingMapToSlice(t *testing.T) {
	fda := &fooDataAccessor{}
	foosByID := make(map[int]*foo)
	foosByID[1] = &foo{ID: 1}
	foosByID[2] = &foo{ID: 2}
	foosByID[3] = &foo{ID: 3}
	foos := fda.mapToSlice(foosByID)
	if len(foos) != 3 {
		t.Fail()
	}
}
