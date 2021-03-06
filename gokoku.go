package gokoku

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Tmpl is a struct for scaffolding
type Tmpl struct {
	IncludeVCSDir, ExcludeDotDir bool
	Suffix                       string
}

var defaultGokoku = &Tmpl{}

// Logger is replaceable logger
var Logger *log.Logger

func logf(format string, v ...interface{}) {
	if Logger == nil {
		log.Printf(format, v...)
		return
	}
	Logger.Printf(format, v...)
}

// Scaffold directory from http.FileSystem
func Scaffold(hfs fs.FS, root, dst string, data interface{}) error {
	return defaultGokoku.Scaffold(hfs, root, dst, data)
}

// Scaffold directory from http.FileSystem
func (tpl *Tmpl) Scaffold(
	hfs fs.FS,
	root, dst string,
	data interface{}) error {
	return fs.WalkDir(hfs, root, func(path string, fi fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			fname := fi.Name()
			if !tpl.IncludeVCSDir {
				switch fname {
				case ".git", ".bzr", ".fossil", ".hg", ".svn":
					return filepath.SkipDir
				}
			}
			if tpl.ExcludeDotDir && strings.HasPrefix(fname, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		dstPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		dstPath = filepath.Join(dst, dstPath)
		buf := &bytes.Buffer{}
		if err := template.Must(template.New(dstPath).
			Option("missingkey=error").Parse(dstPath)).
			Execute(buf, data); err != nil {
			return fmt.Errorf("failed to scaffold while resolving dst Path %q: %w",
				dstPath, err)
		}
		dstPath = buf.String()
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return fmt.Errorf("failed to scaffold while MkdirAll of %q: %w",
				dstPath, err)
		}
		isTmpl := strings.HasSuffix(dstPath, tpl.Suffix)
		if isTmpl {
			dstPath = strings.TrimSuffix(dstPath, tpl.Suffix)
		}
		err = func() (rerr error) {
			logf("Writing %s\n", dstPath)
			targetF, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer func() {
				e := targetF.Close()
				if rerr == nil && e != nil {
					rerr = e
				}
			}()
			datum, err := fs.ReadFile(hfs, path)
			if err != nil {
				return err
			}
			if isTmpl {
				return template.Must(template.New(dstPath+".tmpl").
					Option("missingkey=error").
					Parse(string(datum))).
					Execute(targetF, data)
			}
			_, err = targetF.Write(datum)
			return err
		}()
		if err != nil {
			return fmt.Errorf("failed to scaffold while templating %q: %w",
				dstPath, err)
		}
		return nil
	})
}
