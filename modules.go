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
	a := map[string]any{"net": rsmnet.Module(), "strings": rsmstrings.Module(), "time": rsmtime.Module(), "uuid": rsmuuid.Module(), "sched": rsmsched.Module(), "fmt": rsmfmt.Module(), "json": rsmjson.Module(), "sql": rsmsql.Module(), "strconv": rsmstrconv.Module(), "tablewriter": rsmtablewriter.Module(), "rand": rsmrand.Module(), "semver": rsmsemver.Module(), "bytes": rsmbytes.Module(), "carbon": rsmcarbon.Module(), "color": rsmcolor.Module(), "exec": rsmexec.Module(), "os": rsmos.Module(), "yaml": rsmyaml.Module(), "base64": rsmbase64.Module(), "bcrypt": rsmbcrypt.Module(), "cli": rsmcli.Module(), "http": rsmhttp.Module(), "regexp": rsmregexp.Module()}
	return a
}
