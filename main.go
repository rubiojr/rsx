package main

import (
	"context"
	"embed"
	"os"
	"reflect"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/risor-io/risor"
	"github.com/risor-io/risor/errz"
	ros "github.com/risor-io/risor/os"
)

//go:embed lib/*.risor
var _rsrLib embed.FS

//go:embed lib/rsx.risor
var _rsxLib string

//go:embed lib/pool.risor
var _rsxPool string

//go:embed main.risor
var _mainRsr string

//go:embed main.go
var _mainGo string

//go:embed generator.go
var _generatorGo string

//go:embed importer.go
var _importerGo string

//go:embed repl.go
var _replGo string

//go:embed version.go
var _versionGo string

//go:embed go.mod
var _goMod string

//go:embed go.sum
var _goSum string

//go:embed .modules
var _rsModules string

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
		TimeFormat:      time.Kitchen,
		Prefix:          "",
	})

	if os.Getenv("RSX_DEBUG") != "" {
		logger.SetLevel(log.DebugLevel)
	}

	ctx := context.Background()
	opts := []risor.Option{
		risor.WithConcurrency(),
		risor.WithListenersAllowed(),
		risor.WithLocalImporter("lib"),
		risor.WithGlobal("_generatorGo", _generatorGo),
		risor.WithGlobal("_goMod", _goMod),
		risor.WithGlobal("_goSum", _goSum),
		risor.WithGlobal("_generatorGo", _generatorGo),
		risor.WithGlobal("_importerGo", _importerGo),
		risor.WithGlobal("_mainGo", _mainGo),
		risor.WithGlobal("_replGo", _replGo),
		risor.WithGlobal("_rsModules", _rsModules),
		risor.WithGlobal("_rsxLib", _rsxLib),
		risor.WithGlobal("_rsxVersion", Version),
		risor.WithGlobal("_rsxPool", _rsxPool),
		risor.WithGlobal("_versionGo", _versionGo),
		risor.WithGlobal("pool", _rsxPool),
		risor.WithGlobal("rsx", _rsxLib),
	}

	for k, v := range globalModules() {
		if reflect.ValueOf(v).IsNil() {
			logger.Warnf("Missing module %s, missing build tag?", k)
			continue
		}
		opts = append(opts, risor.WithGlobal(k, v))
	}

	opts = append(opts, risor.WithImporter(newEmbedImporter()))

	if len(os.Args) > 1 && os.Args[1] == "run" {
		m, err := os.ReadFile("main.risor")
		if err != nil {
			logger.Fatal("error reading main.risor")
		}
		_mainRsr = string(m)
		if len(os.Args) > 2 {
			ros.SetScriptArgs(append([]string{"main.risor"}, os.Args[2:]...))
		}
	} else if len(os.Args) > 1 && os.Args[1] == "eval" {
		if len(os.Args) < 3 {
			logger.Fatal("missing script argument")
		}
		script := os.Args[2]
		m, err := os.ReadFile(script)
		if err != nil {
			logger.Fatal("error reading", script)
		}
		_mainRsr = string(m)
		if len(os.Args) > 2 {
			ros.SetScriptArgs(append([]string{script}, os.Args[3:]...))
		}
	} else if len(os.Args) > 1 && os.Args[1] == "repl" {
		Repl(ctx, opts)
	} else {
		ros.SetScriptArgs(os.Args)
	}

	_, err := risor.Eval(
		ctx,
		_mainRsr,
		opts...,
	)
	if err != nil {
		errMsg := err.Error()
		if friendlyErr, ok := err.(errz.FriendlyError); ok {
			errMsg = friendlyErr.FriendlyErrorMessage()
		}
		logger.Fatal(errMsg)
	}
}
