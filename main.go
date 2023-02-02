package main

import (
	"flag"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"log"
)

var cliVersion = flag.String("version", "v1.0.1", "show version of this cli")

func main() {
	flag.Parse()
	log.Printf("=> now version %v", *cliVersion)
	drone_urfave_cli_v2.DroneInfoUrfaveCliFlag()
}
