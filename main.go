package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"time"
)

var (
	noCreate     = flag.Bool("c", false, "do not create any files")
	accessedOnly = flag.Bool("a", false, "only changes the accessed time")
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: touch <file1, file2 ... fileN>\n")
	}
	flag.BoolVar(noCreate, "no-create", false, "do not create any files")
	flag.Parse()

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

func create(file string) {
	if *noCreate {
		return
	}
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("create: cannot create the file %q: %v\n", file, err)
	}
	f.Close()
}

func touch(fi os.FileInfo) {
	var accessed, modded = time.Now().UTC(), time.Now().UTC()
	if *accessedOnly {
		modded = fi.ModTime().UTC()
	}

	if err := os.Chtimes(fi.Name(), accessed, modded); err != nil {
		log.Printf("touch: could not change times for file %q: %v", fi.Name(), err)
	}
}
