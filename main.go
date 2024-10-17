package main

import (
	"context"
	"embed"
	"os"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/risor-io/risor"
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
	//cfg := risor.NewConfig()
	opts := []risor.Option{
		risor.WithConcurrency(),
		risor.WithListenersAllowed(),
		risor.WithLocalImporter("lib"),
		risor.WithGlobals(globalModules()),
		risor.WithGlobal("_mainGo", _mainGo),
		risor.WithGlobal("_generatorGo", _generatorGo),
		risor.WithGlobal("_goMod", _goMod),
		risor.WithGlobal("_goSum", _goSum),
		risor.WithGlobal("_importerGo", _importerGo),
		risor.WithGlobal("_generatorGo", _generatorGo),
		risor.WithGlobal("rsx", _rsxLib),
		risor.WithGlobal("_rsxLib", _rsxLib),
		risor.WithGlobal("pool", _rsxPool),
		risor.WithGlobal("_rsxPool", _rsxPool),
		risor.WithGlobal("_rsModules", _rsModules)}

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
	} else {
		ros.SetScriptArgs(os.Args)
	}

	_, err := risor.Eval(
		ctx,
		_mainRsr,
		opts...,
	)
	if err != nil {
		logger.Fatal(err)
	}
}
