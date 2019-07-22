package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func newNamedNote(name string) string {
	note := filepath.Join(getNoteDir(), name)
	fileName := addMdExtension(note)
	createFile(fileName, ("# " + name + "\n"))
	return fileName
}

func newDefaultNote() string {
	note := defaultNoteName()
	createFile(note, defaultNoteTitle())
	return note
}

func openFile(filePath string) {
	cmd := exec.Command(getEditor(), filePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// createFile makes a file at name only if the file doesn't already exist.
func createFile(name string, contents string) {
	fileContent := []byte(contents)
	permissions := os.FileMode(0666)

	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := ioutil.WriteFile(name, fileContent, permissions)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// ls returns an array of os.Fileinfo objects contained in a given directory.
func ls(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func defaultNoteTitle() string {
	dateFmt := "January 2, 2006"
	currentTime := time.Now()
	return ("# " + currentTime.Format(dateFmt) + "\n\n")
}

func defaultNoteName() string {
	currentTime := time.Now()
	date := currentTime.Format("01-02-2006")
	noteName := filepath.Join(getNoteDir(), date) + ".md"
	return noteName
}

// Exists returns a true if a file exists and false if it doesn't.
func fileExists(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

// makeDir creates the directory at `name` if it doesn't already exist.
func makeDir(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.MkdirAll(name, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addMdExtension(file string) string {
	ext := filepath.Ext(file)
	if ext != ".md" {
		newName := file[0:len(file)-len(ext)] + ".md"
		return newName
	}
	return file
}
