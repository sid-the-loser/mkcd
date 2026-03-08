package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No arguments provided!")
	} else if len(args) > 2 {
		log.Fatal("Only two arguments are expected!")
	}

	dir := ""
	pFlag := false

	for i := range args {
		if args[i] == "-p" {
			pFlag = true
		} else {
			dir = args[i]
		}
	}

	if !pFlag {
		Check(os.Mkdir(dir, os.ModeDir))
	} else {
		Check(os.MkdirAll(dir, os.ModeDir))
	}

	Check(os.Chdir(dir))
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
