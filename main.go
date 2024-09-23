package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

func main() {
	for _, file := range flag.Args() {
		if file == "-" {
			continue
		}

		switch fi, err := os.Stat(file); {
		default:
			touch(fi)
		case errors.Is(err, os.ErrNotExist): // we don't have the file, create it.
			create(file)
		case err != nil:
			log.Printf("Could not touch the file %q: %v\n", file, err)
			os.Exit(readFileAttributeExitCode)
		}
	}
}
