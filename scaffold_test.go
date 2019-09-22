package gokoku

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func tempd(t *testing.T) string {
	tempd, err := ioutil.TempDir("", "gokokutest-")
	if err != nil {
		t.Fatal(err)
	}
	return tempd
}

var testdata = struct {
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

func TestScaffold(t *testing.T) {
	tdir := tempd(t)
	defer os.RemoveAll(tdir)

	err := Scaffold(http.Dir("testdata/basic"), ".", tdir, testdata)
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}

	err = dirDiff(tdir, "./testdata/basic-expect")
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
}

func TestGokoku_Scaffold(t *testing.T) {
	tdir := tempd(t)
	defer os.RemoveAll(tdir)

	tpl := &Tmpl{Suffix: ".tmpl"}
	err := tpl.Scaffold(http.Dir("testdata/basic-suffix"), ".", tdir, testdata)
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}

	err = dirDiff(tdir, "./testdata/basic-expect")
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
}

func dirDiff(dirA, dirB string) error {
	errStr := ""
	err := filepath.Walk(dirA, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dirA, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dirB, rel)
		if fi, err := os.Stat(target); err != nil || fi.IsDir() {
			errStr += fmt.Sprintf("file %q is missing in %q\n", rel, dirB)
			return nil
		}
		got, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		expect, err := ioutil.ReadFile(target)
		if err != nil {
			return err
		}
		if !reflect.DeepEqual(expect, got) {
			errStr += fmt.Sprintf("%q and %q have differ\n", path, target)
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = filepath.Walk(dirB, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dirB, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dirB, rel)
		if fi, err := os.Stat(target); err != nil || fi.IsDir() {
			errStr += fmt.Sprintf("file %q is missing in %q\n", rel, dirA)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
