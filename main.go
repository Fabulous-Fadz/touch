package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

func main() {
	for _, file := range flag.Args() {
		switch fi, err := os.Stat(file); {
		case errors.Is(err, os.ErrNotExist): // we don't have the file, create it.
			create(file)
		case err != nil:
			log.Fatalf("Could not touch the file %q: %v\n", file, err)
		default:
			touch(fi)
		}
	}
}
