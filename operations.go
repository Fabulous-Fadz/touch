package main

import (
	"log"
	"os"
	"time"
)

// The various exit codes used by touch when something goes wrong. Add new ones to the end to avoid changing these
const (
	normalExitCode = iota
	_              // Skipping code 1 to leave it as unclassified in cases we use log.Fatal or log.Fatalf since that uses exit code 1.
	unrecognizedFlagExitCode
	noFilesExitCode
	createFileExitCode
	parseTimeExitCode
	readFileAttributeExitCode
)

var useCurrentTime = true

const (
	version = "(touchwin) touch for Windows version: 0.1.0"
)

func create(file string) {
	if *noCreate {
		return
	}

	f, err := os.Create(file)
	if err != nil {
		log.Printf("create: cannot create the file %q: %v\n", file, err)
		os.Exit(createFileExitCode)
	}
	f.Close()

	// We can set the date to the one picked by the user...
	if !useCurrentTime {
		fi, _ := os.Stat(file) // It shouldn't fail seeing as we just created the file a nanosecond ago (or so)...
		touch(fi)
	}
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
