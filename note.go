package main

import (
	"os"
	"path"
)

// getNoteDir returns the full path of $NOTE_DIR, or ~/notes if no value is set.
func getNoteDir() string {
	noteDir := ""
	envSetting := os.Getenv("NOTE_DIR")
	if len(envSetting) > 0 {
		noteDir += envSetting
	} else {
		home := os.Getenv("HOME")
		noteDir = path.Join(home, "notes")
	}
	return noteDir
}
