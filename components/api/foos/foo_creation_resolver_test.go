package foos

import (
	"fmt"
	"testing"
)

type mockCreator struct{}

func (m *mockCreator) saveFoo(name string) (id int64, err error) {
	return 1, nil
}

func (m *mockCreator) saveBar(fooID, value int) (id int64, err error) {
	if fooID != 1 {
		return 0, fmt.Errorf("No foo")
	}
	return 1, nil
}

func TestCreateFoo(t *testing.T) {
	m := &mockCreator{}
	resolver := &fooCreationResolver{creator: m}
	valueInterface, err := resolver.createFoo("test")
	if err != nil {
		t.Fail()
	}
	value, ok := valueInterface.(createReturnValue)
	if !ok {
		t.Fail()
	}
	if value.ID != 1 {
		t.Fail()
	}
}

func TestCreateBar(t *testing.T) {
	m := &mockCreator{}
	resolver := &fooCreationResolver{creator: m}
	valueInterface, err := resolver.createBar(1, 1)
	if err != nil {
		t.Fail()
	}
	value, ok := valueInterface.(createReturnValue)
	if !ok {
		t.Fail()
	}
	if value.ID != 1 {
		t.Fail()
	}
}
func TestCreateBarForNonExistentFoo(t *testing.T) {
	m := &mockCreator{}
	resolver := &fooCreationResolver{creator: m}
	_, err := resolver.createBar(2, 1)
	if err == nil {
		t.Fail()
	}
}
