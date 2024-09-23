package main

import (
	"log"
	"os"
	"time"
)

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
	var accessed, modded = newTime, newTime
	if *accessedOnly {
		modded = fi.ModTime().UTC()
	} else if *modOnly {
		accessed = time.Time{}
	}

	if err := os.Chtimes(fi.Name(), accessed, modded); err != nil {
		log.Printf("touch: could not change times for file %q: %v", fi.Name(), err)
	}
}
