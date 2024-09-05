package main

import (
	"context"
	"embed"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/risor-io/risor"
	"github.com/risor-io/risor/modules/cli"
	"github.com/risor-io/risor/modules/sql"
	"github.com/risor-io/risor/modules/uuid"
	ros "github.com/risor-io/risor/os"
)

//go:embed lib/*.risor
var _rsrLib embed.FS

//go:embed lib/rsx.risor
var _rsxLib string

//go:embed main.risor
var _mainRsr string

//go:embed main.go
var _mainGo string

//go:embed importer.go
var _importerGo string

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
	importerOpt := risor.WithImporter(newEmbedImporter())
	//cfg := risor.NewConfig()

	ros.SetScriptArgs(os.Args)
	_, err := risor.Eval(
		ctx,
		_mainRsr,
		importerOpt,
		risor.WithConcurrency(),
		//risor.WithGlobals(cfg.Globals()),
		risor.WithGlobal("cli", cli.Module()),
		risor.WithGlobal("sql", sql.Module()),
		risor.WithGlobal("uuid", uuid.Module()),
		risor.WithGlobal("_mainGo", _mainGo),
		risor.WithGlobal("_importerGo", _importerGo),
		risor.WithGlobal("rsxLib", _rsxLib),
	)
	if err != nil {
		logger.Fatal(err)
	}
}
