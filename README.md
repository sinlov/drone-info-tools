[![go-ubuntu](https://github.com/sinlov/drone-info-tools/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-info-tools/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-info-tools?status.png)](https://godoc.org/github.com/sinlov/drone-info-tools/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-info-tools)](https://goreportcard.com/report/github.com/sinlov/drone-info-tools)
[![codecov](https://codecov.io/gh/sinlov/drone-info-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-info-tools)

## for what

- this tools use for bind drone env for [https://plugins.drone.io/](https://plugins.drone.io/)

## usage

```go
package main

import (
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	droneFlag := drone_urfave_cli_v2.DroneInfoUrfaveCliFlag()
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(droneFlag, pluginFlag())

	app.Flags = flags
}

// action
// do cli Action before flag.
func action(c *cli.Context) error {
	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)
}
```

### tools

```go
package main

import (
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
)

func main()  {
	tools.StrInArr(foo, bar)
}
```

### template

```
	github.com/Masterminds/sprig/v3 v3.1.0
	github.com/aymerick/raymond v2.0.2+incompatible
```

## usage

- `only once template.RegisterSettings` like file `init.go`

```go
package main

import (
	"github.com/sinlov/drone-info-tools/template"
)

func init() {
  template.RegisterSettings(template.DefaultFunctions)
}
```

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlov/drone-info-tools.git

# test depends see full version
$ go list -mod=readonly -v -m -versions github.com/sinlov/drone-info-tools
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov/drone-info-tools | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.17+

# dev

```bash
make init dep
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# see how to use
$ drone-info-tools -h
```
