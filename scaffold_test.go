package gokoku

import (
	"net/http"
	"testing"
)

func TestScaffold(t *testing.T) {
	Scaffold(http.Dir("testdata/assets"), ".", ".tmp/scaf2", nil)

	if false {
		t.Errorf("something went wrong")
	}
}
