package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		println("Please provide a file name")
		os.Exit(1)
	}

	file, err := os.Open(args[0])
	if err != nil {
		return
	}

	_, err1 := io.Copy(os.Stdout, file)
	if err1 != nil {
		return
	}
}
