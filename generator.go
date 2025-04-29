//go:generate go run generator.go
//go:build generator

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

const modsFile = ".modules"
const baseModulePath = "github.com/risor-io/risor/modules/"
const modPrefix = "rsm"

func main() {
	source := jen.NewFile("main")
	source.NoFormat = false
	source.HeaderComment("//go:generate go run generator.go")

	list := datasources()
	if len(list) == 0 {
		return
	}

	globals := map[string]string{}
	for _, ds := range list {
		mAlias := modPrefix + filepath.Base(ds)
		globals[ds] = mAlias
	}

	var stmts []jen.Code
	for k, _ := range globals {
		key := filepath.Base(k)
		mPath := k
		if strings.Contains(k, "@") {
			tokens := strings.Split(k, "@")
			key = tokens[0]
			mPath = tokens[1]
		}
		alias := modPrefix + strings.Replace(key, "-", "", -1)
		source.ImportAlias(mPath, alias)
		stmts = append(stmts, jen.Lit(key).Op(":").Qual(mPath, "Module").Call())
	}

	source.Func().Id("globalModules").Params().Map(jen.String()).Any().Block(
		jen.Id("a").Op(":=").Map(jen.String()).Any().Values(
			stmts...,
		),
		jen.Return(jen.Id("a")),
	)

	f, err := os.Create("modules.go")
	if err != nil {
		genFailed(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%#v", source)
}

func datasources() []string {
	var datasources []string

	f, err := os.Open(modsFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "** no %s file found\n", modsFile)
			return datasources
		}
		genFailed(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ds := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(ds, "//") || ds == "" {
			continue
		}

		// local data sources
		if !strings.Contains(ds, "/") {
			ds = baseModulePath + ds
		}

		datasources = append(datasources, ds)
	}

	return datasources
}

func genFailed(err error) {
	fmt.Fprintf(os.Stderr, "generating modules.go failed: %s", err)
	os.Exit(1)
}
