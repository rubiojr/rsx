package main

import (
	"context"
	"embed"
	"flag"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/risor-io/risor"
	"github.com/risor-io/risor/modules/cli"
	"github.com/risor-io/risor/modules/sql"
	"github.com/risor-io/risor/modules/uuid"
	ros "github.com/risor-io/risor/os"
)

//go:embed lib/*.risor
var rsrLib embed.FS

//go:embed lib/rsx.risor
var rsxLib string

//go:embed main.risor
var rsrEntry string

//go:embed main.go
var mainGo string

func main() {
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.Kitchen,
		Prefix:          "",
	})

	if *debug {
		logger.SetLevel(log.DebugLevel)
	}

	ctx := context.Background()
	tempDir, err := os.MkdirTemp("", "risor-lib")
	if err != nil {
		logger.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	evalCode := ""
	err = fs.WalkDir(rsrLib, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			err := os.MkdirAll(filepath.Dir(filepath.Join(tempDir, path)), 0755)
			if err != nil {
				return err
			}

			dst, err := os.Create(filepath.Join(tempDir, path))
			if err != nil {
				return err
			}

			defer dst.Close()
			src, err := rsrLib.Open(path)
			if err != nil {
				return err
			}
			defer src.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}

			base := filepath.Base(path)
			var extension = filepath.Ext(base)
			var name = base[0 : len(base)-len(extension)]
			logger.Debugf("available module %s...", name)
		}
		return nil
	})

	if err != nil {
		logger.Fatal(err)
	}

	evalCode += rsrEntry
	importerOpt := risor.WithLocalImporter(filepath.Join(tempDir, "/lib"))
	ros.SetScriptArgs(os.Args)
	_, err = risor.Eval(
		ctx,
		evalCode,
		importerOpt,
		risor.WithGlobal("cli", cli.Module()),
		risor.WithGlobal("sql", sql.Module()),
		risor.WithGlobal("uuid", uuid.Module()),
		risor.WithGlobal("mainGo", mainGo),
		risor.WithGlobal("rsxLib", rsxLib),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
