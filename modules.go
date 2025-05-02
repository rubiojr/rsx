//go:generate go run generator.go

package main

import (
	rsmbase64 "github.com/risor-io/risor/modules/base64"
	rsmbcrypt "github.com/risor-io/risor/modules/bcrypt"
	rsmbytes "github.com/risor-io/risor/modules/bytes"
	rsmcarbon "github.com/risor-io/risor/modules/carbon"
	rsmcli "github.com/risor-io/risor/modules/cli"
	rsmcolor "github.com/risor-io/risor/modules/color"
	rsmexec "github.com/risor-io/risor/modules/exec"
	rsmfmt "github.com/risor-io/risor/modules/fmt"
	rsmhttp "github.com/risor-io/risor/modules/http"
	rsmjson "github.com/risor-io/risor/modules/json"
	rsmnet "github.com/risor-io/risor/modules/net"
	rsmos "github.com/risor-io/risor/modules/os"
	rsmrand "github.com/risor-io/risor/modules/rand"
	rsmregexp "github.com/risor-io/risor/modules/regexp"
	rsmsemver "github.com/risor-io/risor/modules/semver"
	rsmsql "github.com/risor-io/risor/modules/sql"
	rsmstrconv "github.com/risor-io/risor/modules/strconv"
	rsmstrings "github.com/risor-io/risor/modules/strings"
	rsmtablewriter "github.com/risor-io/risor/modules/tablewriter"
	rsmtemplate "github.com/risor-io/risor/modules/template"
	rsmtime "github.com/risor-io/risor/modules/time"
	rsmuuid "github.com/risor-io/risor/modules/uuid"
	rsmyaml "github.com/risor-io/risor/modules/yaml"
)

func globalModules() map[string]any {
	a := map[string]any{"exec": rsmexec.Module(), "fmt": rsmfmt.Module(), "json": rsmjson.Module(), "net": rsmnet.Module(), "os": rsmos.Module(), "rand": rsmrand.Module(), "bytes": rsmbytes.Module(), "cli": rsmcli.Module(), "template": rsmtemplate.Module(), "regexp": rsmregexp.Module(), "uuid": rsmuuid.Module(), "yaml": rsmyaml.Module(), "semver": rsmsemver.Module(), "base64": rsmbase64.Module(), "bcrypt": rsmbcrypt.Module(), "http": rsmhttp.Module(), "sql": rsmsql.Module(), "strconv": rsmstrconv.Module(), "strings": rsmstrings.Module(), "tablewriter": rsmtablewriter.Module(), "time": rsmtime.Module(), "carbon": rsmcarbon.Module(), "color": rsmcolor.Module()}
	return a
}
