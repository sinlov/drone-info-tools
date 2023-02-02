package main

import (
	"flag"
	"log"
)

var cliVersion = flag.String("version", "v1.0.1", "show version of this cli")

func main() {
	flag.Parse()
	log.Printf("=> now version %v", *cliVersion)
}
