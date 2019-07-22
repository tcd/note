// note provides a command line interface for managing note files.
package main

import (
	"flag"
	"fmt"
	"os"
)

var newFlag = flag.String("new", "", "Pass a string for the name of a new note")
var listFlag = flag.Bool("list", false, "List all notes")

func init() {
	flag.StringVar(newFlag, "n", "", "")
	flag.BoolVar(listFlag, "l", false, "")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if !fileExists(getNoteDir()) {
		makeDir(getNoteDir())
	}

	if *listFlag {
		for _, note := range ls(getNoteDir()) {
			fmt.Println(note.Name())
		}
	} else if len(*newFlag) > 0 {
		openFile(newNamedNote(*newFlag))
	} else {
		openFile(newDefaultNote())
	}
}

func usage() {
	title := `                __
   ____  ____  / /____
  / __ \/ __ \/ __/ _ \
 / / / / /_/ / /_/  __/
/_/ /_/\____/\__/\___/
`
	fmt.Println(title)
	fmt.Printf("Usage: %s [OPTIONS] argument\n", os.Args[0])
	flag.PrintDefaults()
}
