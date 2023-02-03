package main

import (
	"flag"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"log"
)

const (
	Version = "v1.3.0"
)

var cliVersion = flag.String("version", Version, "show version of this cli")

func main() {
	flag.Parse()
	log.Printf("=> now version %v", *cliVersion)
	drone_urfave_cli_v2.DroneInfoUrfaveCliFlag()
}
