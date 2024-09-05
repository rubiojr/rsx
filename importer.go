package main

import (
	"context"
	"fmt"

	"github.com/risor-io/risor"
	"github.com/risor-io/risor/compiler"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/parser"
)

type embedImporter struct {
	globals []string
}

func newEmbedImporter() *embedImporter {
	cfg := risor.NewConfig()
	return &embedImporter{
		globals: cfg.GlobalNames(),
	}
}

func (i *embedImporter) Import(ctx context.Context, name string) (*object.Module, error) {
	source, err := _rsrLib.ReadFile("lib/" + name + ".risor")
	if err != nil {
		return nil, fmt.Errorf("import error: module %q not found", name)
	}

	ast, err := parser.Parse(ctx, string(source))
	if err != nil {
		return nil, err
	}

	var opts []compiler.Option
	opts = append(opts, compiler.WithGlobalNames(i.globals))
	code, err := compiler.Compile(ast, opts...)
	if err != nil {
		return nil, err
	}

	return object.NewModule(name, code), nil
}
