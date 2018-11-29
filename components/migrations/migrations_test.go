package migrations

import (
	"testing"
)

func TestLoadSources(t *testing.T) {
	sourceDriver, err := newSourceDriver()
	if err != nil {
		t.Error(err)
	}
	_, err = sourceDriver.First()
	if err != nil {
		t.Error(err)
	}
}
