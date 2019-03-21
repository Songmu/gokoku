package gokoku_test

import (
	"testing"

	"github.com/Songmu/gokoku"
)

func TestLoadPlugin(t *testing.T) {
	_, err := gokoku.LoadPlugin("testdata/plugin.so")
	if err != nil {
		t.Errorf("err occurred: %s", err)
	}
}
