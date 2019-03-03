package gokoku

import (
	"net/http"
	"plugin"

	"golang.org/x/xerrors"
)

// LoadPlugin is experimental feature to get http.FileSystem from Go plugin
func LoadPlugin(plugPath string) (http.FileSystem, error) {
	plug, err := plugin.Open(plugPath)
	if err != nil {
		return nil, xerrors.Errorf("failed to load plugin %s: %w", plugPath, err)
	}
	symbol, err := plug.Lookup("Fs")
	if err != nil {
		return nil, xerrors.Errorf("Fs func isn't defined in %s: %w", plugPath, err)
	}
	if fn, ok := symbol.(func() (http.FileSystem, error)); ok {
		return fn()
	}
	return nil, xerrors.New("Fs func should be `func() (http.FileSystem, error))`")
}
