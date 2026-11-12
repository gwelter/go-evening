package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	showHidden := flag.Bool("a", true, "Show hidden files")
	flag.Parse()

	files := listFiles("testdata", *showHidden)
	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(dirname string, showHidden bool) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !showHidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		dirs = append(dirs, f.Name())
	}

	return dirs
}
