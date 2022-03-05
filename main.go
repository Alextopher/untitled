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

	for item := range l.items {
		fmt.Print(item)
	}
}
