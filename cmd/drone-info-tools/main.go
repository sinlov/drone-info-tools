//go:build !test

package main

import (
	_ "embed"
	resource "github.com/sinlov/drone-info-tools"
	"github.com/sinlov/drone-info-tools/cmd/cli"
	"github.com/sinlov/drone-info-tools/drone_log"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"os"
)

func main() {
	pkgJson.InitPkgJsonContent(resource.PackageJson)
	template.RegisterSettings(template.DefaultFunctions)

	app := cli.NewCliApp()

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		drone_log.Warnf("run err: %v", err)
	}
}
