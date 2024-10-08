//go:build !test

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const format = time.RFC3339

var (
	newTime       time.Time = time.Now().UTC()
	noCreate                = flag.Bool("c", false, "do not create any files")
	accessedOnly            = flag.Bool("a", false, "only changes the accessed time")
	_                       = flag.Bool("h", false, "currently ignored in this implementation")
	help                    = flag.Bool("help", false, "displays this help text and exits")
	_                       = flag.Bool("f", false, "(ignored)")
	full                    = flag.Bool("full", false, "displays usage information, including exit codes. Assumes --help is specified")
	modOnly                 = flag.Bool("m", false, "only changes the modified time")
	_                       = flag.Bool("no-dereference", false, "currently ignored in this implementation")
	referenceFile           = flag.String("r", "", "use the specified file's times instead of the current system time")
	userTime                = flag.String("t", "", "-t sets a specified time instead of the default current system time")
	versionOnly             = flag.Bool("version", false, "output version information and exit")
)

func init() {
	if len(os.Args) == 1 {
		println("Usage: touch <file1, file2 ... fileN>\n")
		os.Exit(noFilesExitCode)
	}

	flag.BoolVar(noCreate, "no-create", false, "do not create any files")
	flag.StringVar(referenceFile, "reference", "", "use this file's times insead of current time")

	flag.Parse()

	if *versionOnly {
		println(version)
		os.Exit(normalExitCode)
	}
	if *help || *full {
		flag.Usage()
		if *full {
			fmt.Println(fullHelp)
		}
		os.Exit(normalExitCode)
	}

	// I could just leave them both untouched but the user needs to think about their life choices very seriously.
	// Not to nitpick... but why call a program that touches 2 dates and tell it not to touch either of them? What is that? Come on? Be serious!
	if *modOnly && *accessedOnly {
		log.Fatal("You cannot set both -a and -m. Specify just one or omit both of them to modify both dates.")
	}

	// Date from a reference file takes precedence over any supplied date string in this implementation. Check for either.
	if len(*referenceFile) > 0 {
		switch fi, err := os.Stat(*referenceFile); {
		default:
			newTime = fi.ModTime()
			useCurrentTime = false
		case errors.Is(err, os.ErrNotExist):
			log.Printf("touch: failed to get attributes of %q: No such file or directory\n", *referenceFile)
			fallthrough
		case err != nil:
			os.Exit(readFileAttributeExitCode)
		}
	} else if *userTime != "" {
		t, err := time.Parse(format, *userTime)
		if err != nil {
			log.Printf("%q is invalid as a date of format: %q\n", *userTime, format)
			os.Exit(parseTimeExitCode)
		}
		newTime = t
		useCurrentTime = false
	}
}
