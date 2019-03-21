package main

import (
	"net/http"

	// import statik
	_ "github.com/Songmu/gokoku/example/statik"
	"github.com/rakyll/statik/fs"
)

func Fs() (http.FileSystem, error) {
	return fs.New()
}
