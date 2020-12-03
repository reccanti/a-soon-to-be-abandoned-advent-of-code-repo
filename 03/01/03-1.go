package main

import (
	"os"
)

func main() {
	filename := os.Args[1]
	input, err := parseRelativeFile(filename)
}
