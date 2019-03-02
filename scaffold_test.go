package gokoku

import (
	"net/http"
	"testing"
)

func TestScaffold(t *testing.T) {
	err := Scaffold(http.Dir("testdata/basic"), ".", ".tmp/scaf2", nil)
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}

	if false {
		t.Errorf("something went wrong")
	}
}
