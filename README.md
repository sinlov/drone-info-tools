[![golang-ci](https://github.com/sinlov/drone-info-tools/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/sinlov/drone-info-tools/actions/workflows/golang-ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov/drone-info-tools?label=go.mod)](https://github.com/sinlov/drone-info-tools)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-info-tools?status.png)](https://godoc.org/github.com/sinlov/drone-info-tools)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov/drone-info-tools)](https://goreportcard.com/report/github.com/sinlov/drone-info-tools)
[![GitHub license](https://img.shields.io/github/license/sinlov/drone-info-tools)](https://github.com/sinlov/drone-info-tools)
[![codecov](https://codecov.io/gh/sinlov/drone-info-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-info-tools)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov/drone-info-tools)](https://github.com/sinlov/drone-info-tools/tags)

## for what

- this tools use for bind drone env for [https://plugins.drone.io/](https://plugins.drone.io/)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/drone-info-tools)](https://github.com/sinlov/drone-info-tools/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## usage

```go
package main

import (
	_ "embed"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	resource "github.com/sinlov/drone-info-tools"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/urfave/cli/v2"
)

//go:embed package.json
var packageJson string

func main() {
	pkgJson.InitPkgJsonContent(resource.PackageJson)

	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
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

func main() {
	tools.StrInArr(foo, bar)
}
```

### template

```
	github.com/Masterminds/sprig/v3 v3.1.0
	github.com/aymerick/raymond v2.0.2+incompatible
```

## usage

- `only once template.RegisterSettings` like file `main.go`

```go
package main

import (
	"github.com/sinlov/drone-info-tools/template"
)

func main() {
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

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                  | version             |
|:-------------------------------------|:--------------------|
| https://github.com/stretchr/testify  | v1.8.4              |
| https://github.com/urfave/cli        | v2.25.7             |
| https://github.com/aymerick/raymond  | v2.0.2+incompatible |
| https://github.com/Masterminds/sprig | v3.2.3              |

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
