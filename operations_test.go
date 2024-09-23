//go:build test

package main

import (
	"errors"
	"os"
	"testing"
	"time"
)

var (
	newTime      time.Time = time.Time{}
	noCreate               = new(bool)
	modOnly                = new(bool)
	accessedOnly           = new(bool)
)

func TestCreateFile(t *testing.T) {
	newTime = time.Date(1980, 6, 6, 23, 45, 00, 0, time.Local)
	useCurrentTime = false
	fileName := "a file.dat"

	create(fileName)

	// Check...
	fi, err := os.Stat(fileName)
	if err != nil {
		t.Errorf("Could not create file")
	}
	defer os.Remove(fileName)

	if fi.ModTime() != newTime {
		t.Errorf("Created time does ot match what we want. Have: %v, want: %v", fi.ModTime(), newTime)
	}

	// Don't create....
	*noCreate = true
	fileName = "another file.txt"
	create(fileName)

	// Check...
	if fi, err = os.Stat(fileName); err == nil {
		t.Errorf("A file was created when -noCreate was specified")
		defer os.Remove(fileName)
	} else if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Want: 'file not found error', have: %v. Options are --no-create", err)
	}
}
