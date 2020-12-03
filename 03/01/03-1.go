package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	// get the "hill" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	hill := *input

	// scan the hill
	rows := strings.Split(hill, "\n")
	xPos := 0
	treeCount := 0
	for _, row := range rows {
		// determine if there is a tree at the current position
		char := string(row[xPos])
		if char == "#" {
			treeCount += 1
		}
		// increment the x position
		xPos = (xPos + 3) % len(row)
	}
	fmt.Println(treeCount)
}
