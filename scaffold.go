package gokoku

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rakyll/statik/fs"
	"golang.org/x/xerrors"
)

func Scaffold(hfs http.FileSystem, root, dst string, data interface{}) error {
	return fs.Walk(hfs, root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			if fi.Name() == ".git" {
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
		if err := template.Must(template.New(dstPath).Parse(dstPath)).Execute(buf, data); err != nil {
			return xerrors.Errorf("failed to scaffold while resolving dst Path %q: %w",
				dstPath, err)
		}
		dstPath = buf.String()
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return xerrors.Errorf("failed to scaffold while MkdirAll of %q: %w",
				dstPath, err)
		}
		err = func() (rerr error) {
			log.Printf("Writing %s\n", dstPath)
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
			return template.Must(template.New(dstPath+".tmpl").Parse(string(datum))).
				Execute(targetF, data)
		}()
		if err != nil {
			return xerrors.Errorf("failed to scaffold while templating %q: %w",
				dstPath, err)
		}
		return nil
	})
}
