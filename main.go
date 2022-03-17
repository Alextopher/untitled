package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open file and read it into a string
	if len(os.Args) != 2 {
		fmt.Println("Usage: lexer <file>")
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l := lex(string(bytes))

	// Keep track if we just printed a newline
	var lastNewLine bool

	for item := range l.items {
		if lastNewLine {
			if item.typ == itemNewLine {
				continue
			} else {
				lastNewLine = false
			}
		}

		if item.typ == itemNewLine {
			lastNewLine = true
		}

		fmt.Print(item)
	}
}
