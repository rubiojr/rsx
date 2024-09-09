package main

import (
	"context"
	"embed"
	"os"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/risor-io/risor"
	mbase64 "github.com/risor-io/risor/modules/base64"
	mbcrypt "github.com/risor-io/risor/modules/bcrypt"
	mbytes "github.com/risor-io/risor/modules/bytes"
	mcarbon "github.com/risor-io/risor/modules/carbon"
	mcli "github.com/risor-io/risor/modules/cli"
	mcolor "github.com/risor-io/risor/modules/color"
	mexec "github.com/risor-io/risor/modules/exec"
	mfmt "github.com/risor-io/risor/modules/fmt"
	mhttp "github.com/risor-io/risor/modules/http"
	mjson "github.com/risor-io/risor/modules/json"
	mnet "github.com/risor-io/risor/modules/net"
	mos "github.com/risor-io/risor/modules/os"
	mrand "github.com/risor-io/risor/modules/rand"
	mregexp "github.com/risor-io/risor/modules/regexp"
	msql "github.com/risor-io/risor/modules/sql"
	mstrconv "github.com/risor-io/risor/modules/strconv"
	mstrings "github.com/risor-io/risor/modules/strings"
	mtime "github.com/risor-io/risor/modules/time"
	muuid "github.com/risor-io/risor/modules/uuid"
	myaml "github.com/risor-io/risor/modules/yaml"
	ros "github.com/risor-io/risor/os"
	"github.com/rubiojr/rsx/modules/rsxmod"
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

//go:embed go.mod
var _goMod string

//go:embed go.sum
var _goSum string

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
		risor.WithLocalImporter("lib"),
		//risor.WithGlobals(cfg.Globals()),
		risor.WithGlobal("cli", mcli.Module()),
		risor.WithGlobal("sql", msql.Module()),
		risor.WithGlobal("uuid", muuid.Module()),
		risor.WithGlobal("base64", mbase64.Module()),
		risor.WithGlobal("http", mhttp.Module()),
		risor.WithGlobal("exec", mexec.Module()),
		risor.WithGlobal("json", mjson.Module()),
		risor.WithGlobal("bytes", mbytes.Module()),
		risor.WithGlobal("strings", mstrings.Module()),
		risor.WithGlobal("yaml", myaml.Module()),
		risor.WithGlobal("time", mtime.Module()),
		risor.WithGlobal("os", mos.Module()),
		risor.WithGlobal("regexp", mregexp.Module()),
		risor.WithGlobal("net", mnet.Module()),
		risor.WithGlobal("fmt", mfmt.Module()),
		risor.WithGlobal("color", mcolor.Module()),
		risor.WithGlobal("bcrypt", mbcrypt.Module()),
		risor.WithGlobal("carbon", mcarbon.Module()),
		risor.WithGlobal("rand", mrand.Module()),
		risor.WithGlobal("atoi", mstrconv.Module()),
		risor.WithGlobal("rsxmod", rsxmod.Module()),
		risor.WithGlobal("_mainGo", _mainGo),
		risor.WithGlobal("_goMod", _goMod),
		risor.WithGlobal("_goSum", _goSum),
		risor.WithGlobal("_importerGo", _importerGo),
		risor.WithGlobal("_rsxLib", _rsxLib)}

	if len(os.Args) > 1 && os.Args[1] == "run" {
		m, err := os.ReadFile("main.risor")
		if err != nil {
			logger.Fatal("error reading main.risor")
		}
		_mainRsr = string(m)
		opts = append(opts, risor.WithLocalImporter("lib"))
		if len(os.Args) > 2 {
			ros.SetScriptArgs(append([]string{"main.risor"}, os.Args[2:]...))
		}
	} else {
		opts = append(opts, risor.WithImporter(newEmbedImporter()))
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
