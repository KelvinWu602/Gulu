package main

import (
	"fmt"
	"io"
	"os"

	"github.com/KelvinWu602/Gulu/core/scanner"
)

func main() {
	if len(os.Args) > 2 {
		os.Exit(1)
	} else if len(os.Args) == 2 {
		readFile(os.Args[1])
	} else {
		REPL()
	}
}

func check(e error) {
	if e != nil {
		fmt.Println("Unable to read the file.")
		os.Exit(1)
	}
}

func readFile(arg string) {
	src, err := os.ReadFile(arg)
	check(err)
	run(string(src))
}

func REPL() {
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		run(line)
	}
}

func run(code string) {
	scanner.Scan("wkefu")
}
