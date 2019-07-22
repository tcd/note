package main

import (
	"os"
	"path/filepath"
)

// getNoteDir returns the full path of $NOTE_DIR, or ~/notes if no value is set.
func getNoteDir() string {
	noteDir := ""
	envSetting := os.Getenv("NOTE_DIR")
	if len(envSetting) > 0 {
		noteDir += envSetting
	} else {
		home := os.Getenv("HOME")
		noteDir = filepath.Join(home, "notes")
	}
	return noteDir
}

// getEditor returns the value of $EDITOR, or "vim" if no value is set
func getEditor() string {
	var editor string
	envSetting := os.Getenv("EDITOR")
	if len(envSetting) > 0 {
		editor += envSetting
	} else {
		editor = "vim"
	}
	return editor
}
