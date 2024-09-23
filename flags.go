package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"time"
)

var (
	newTime       time.Time = time.Now().UTC()
	noCreate                = flag.Bool("c", false, "do not create any files")
	accessedOnly            = flag.Bool("a", false, "only changes the accessed time")
	help                    = flag.Bool("help", false, "displays this help text and exits")
	modOnly                 = flag.Bool("m", false, "only changes the modified time")
	referenceFile           = flag.String("r", "", "use the specified file's times instead of the current system time")
)

func init() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: touch <file1, file2 ... fileN>\n")
	}

	flag.BoolVar(noCreate, "no-create", false, "do not create any files")
	flag.StringVar(referenceFile, "reference", "", "use this file's times insead of current time")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if len(*referenceFile) > 0 {
		switch fi, err := os.Stat(*referenceFile); {
		default:
			newTime = fi.ModTime()
			useCurrentTime = false
		case errors.Is(err, os.ErrNotExist):
			log.Printf("touch: failed to get attributes of %q: No such file or directory\n", *referenceFile)
			fallthrough
		case err != nil:
			os.Exit(2)
		}
	}
}
