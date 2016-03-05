package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func isExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "invalid args")
		os.Exit(1)
	}

	var wantPath = os.Args[1]
	var startDir string
	if len(os.Args) == 3 {
		if path.IsAbs(startDir) {
			startDir = os.Args[2]
		} else {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			startDir = filepath.Join(wd, startDir)
		}
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		startDir = wd
	}

	var prevDir string
	for dir := startDir; ; {
		if isExist(filepath.Join(dir, wantPath)) {
			fmt.Printf("%s\n", dir)
			return
		}

		dir = path.Dir(dir)
		if prevDir == dir {
			fmt.Fprintln(os.Stderr, "file does not exist")
			os.Exit(1)
		}
		prevDir = dir
	}
}
