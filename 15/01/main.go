package main

import (
	"fmt"
	"os"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	filename := os.Args[1]
	nums, err := util.ParseRelativeFileInts(filename, ",")
	if err != nil {
		fmt.Println(err)
		return
	}

	entries := map[int]int{}
	speak := 0
	for turn, n := range nums {
		lastSaid, wasSaid := entries[n]
		fmt.Println(fmt.Sprintf("Turn %d: %d", turn+1, n))
		if !wasSaid {
			speak = 0
		} else {
			speak = turn - lastSaid
		}
		entries[n] = turn
	}

	for turn := len(nums); turn < 2020; turn++ {
		fmt.Println(fmt.Sprintf("Turn %d: %d", turn+1, speak))
		lastSaid, wasSaid := entries[speak]
		entries[speak] = turn
		if !wasSaid {
			speak = 0
		} else {
			speak = turn - lastSaid
		}
	}
}
