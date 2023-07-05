//go:build !test

package main

import (
	_ "embed"
	"flag"
	resource "github.com/sinlov/drone-info-tools"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"log"
)

const (
	Name = "drone-info-tools"
)

var cliVersion *string

func main() {
	pkgJson.InitPkgJsonContent(resource.PackageJson)
	cliVersion = flag.String("version", pkgJson.GetPackageJsonVersionGoStyle(), "show version of this cli")

	flag.Parse()
	log.Printf("=> now version %v", *cliVersion)
	drone_urfave_cli_v2.DroneInfoUrfaveCliFlag()
}
