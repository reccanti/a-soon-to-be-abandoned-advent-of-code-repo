package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type HillInfo struct {
	trees int
	empty int
}

func traverseHill(hill string, incRight int, incDown int) HillInfo {
	rows := strings.Split(hill, "\n")
	xPos := 0
	yPos := 0
	treeCount := 0
	emptyCount := 0

	for yPos < len(rows)-1 {
		// move to the next tile on the hill
		yPos += incDown
		xPos = (xPos + incRight) % len(rows[yPos])

		// retrieve the character at that tile
		char := string(rows[yPos][xPos])

		// update the count appropriately
		if char == "#" {
			treeCount += 1
		} else {
			emptyCount += 1
		}
	}

	// return the hill info
	return HillInfo{
		trees: treeCount,
		empty: emptyCount,
	}
}

func main() {
	// get the "hill" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	hill := *input

	// traverse the hill and gather data
	hillInfo1 := traverseHill(hill, 1, 1)
	hillInfo2 := traverseHill(hill, 3, 1)
	hillInfo3 := traverseHill(hill, 5, 1)
	hillInfo4 := traverseHill(hill, 7, 1)
	hillInfo5 := traverseHill(hill, 1, 2)

	fmt.Println(hillInfo1.trees * hillInfo2.trees * hillInfo3.trees * hillInfo4.trees * hillInfo5.trees)
}
