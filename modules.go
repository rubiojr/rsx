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
	rsmsql "github.com/risor-io/risor/modules/sql"
	rsmstrconv "github.com/risor-io/risor/modules/strconv"
	rsmstrings "github.com/risor-io/risor/modules/strings"
	rsmtablewriter "github.com/risor-io/risor/modules/tablewriter"
	rsmtime "github.com/risor-io/risor/modules/time"
	rsmuuid "github.com/risor-io/risor/modules/uuid"
	rsmyaml "github.com/risor-io/risor/modules/yaml"
	rsmsched "github.com/rubiojr/risor-modules/sched"
	rsmsemver "github.com/rubiojr/risor-modules/semver"
)

func globalModules() map[string]any {
	a := map[string]any{"time": rsmtime.Module(), "base64": rsmbase64.Module(), "bytes": rsmbytes.Module(), "fmt": rsmfmt.Module(), "net": rsmnet.Module(), "rand": rsmrand.Module(), "sql": rsmsql.Module(), "cli": rsmcli.Module(), "strconv": rsmstrconv.Module(), "strings": rsmstrings.Module(), "sched": rsmsched.Module(), "semver": rsmsemver.Module(), "json": rsmjson.Module(), "yaml": rsmyaml.Module(), "regexp": rsmregexp.Module(), "tablewriter": rsmtablewriter.Module(), "bcrypt": rsmbcrypt.Module(), "carbon": rsmcarbon.Module(), "color": rsmcolor.Module(), "exec": rsmexec.Module(), "http": rsmhttp.Module(), "os": rsmos.Module(), "uuid": rsmuuid.Module()}
	return a
}
