package main

import (
	"fmt"
	"os"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	fmt.Println(*input)
}
