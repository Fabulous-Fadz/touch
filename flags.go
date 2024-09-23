package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	newTime      time.Time = time.Now().UTC()
	noCreate               = flag.Bool("c", false, "do not create any files")
	accessedOnly           = flag.Bool("a", false, "only changes the accessed time")
	help                   = flag.Bool("help", false, "displays this help text and exits")
	modOnly                = flag.Bool("m", false, "only changes the modified time")
)

func init() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: touch <file1, file2 ... fileN>\n")
	}
	flag.BoolVar(noCreate, "no-create", false, "do not create any files")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}
}
