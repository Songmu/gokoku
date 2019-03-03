package gokoku

import (
	"net/http"
	"testing"
)

func TestScaffold(t *testing.T) {
	data := struct {
		Author, PackagePath        string
		GitHubHost, Owner, Package string
		Year                       int
	}{
		Author:      "Songmu",
		PackagePath: "github.com/Songmu/gokoku",
		GitHubHost:  "github.com",
		Owner:       "Songmu",
		Package:     "gokoku",
		Year:        2019,
	}

	err := Scaffold(http.Dir("testdata/basic"), ".", ".tmp/scaf22", data)
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
}
