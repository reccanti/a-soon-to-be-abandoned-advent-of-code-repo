package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type SeatLocation struct {
	row    int
	column int
}

/**
 * @NOTE this function assumes that str will be 10 characters
 * long, where the first 7 are used to determine one of the 128
 * rows, and the last 3 are used to determine one of the 8 columns.
 * If it turns out that we need to change this to support an arbitrary
 * number of rows and columns, we should change how we calculate the
 * "pivot" point, where we split the row and column data
 *
 * ~reccanti 12/5/2020
 */
func getSeatLocation(seatStr string) SeatLocation {

	// calculate the row
	rowChars := seatStr[:7]
	maxRow := 127
	minRow := 0
	for _, char := range rowChars {
		change := (maxRow - minRow + 1) / 2
		if string(char) == "F" {
			maxRow -= change
		} else if string(char) == "B" {
			minRow += change
		}
	}

	// calculate the column
	columnChars := seatStr[7:]
	maxColumn := 7
	minColumn := 0
	for _, char := range columnChars {
		change := (maxColumn - minColumn + 1) / 2
		if string(char) == "L" {
			maxColumn -= change
		} else if string(char) == "R" {
			minColumn += change
		}
	}

	return SeatLocation{
		row:    maxRow,
		column: maxColumn,
	}
}

/**
 * @NOTE Again, this function assumes that there will
 * be 8 columns, so this might need to be expanded in
 * the future
 *
 * ~reccanti 12/5/2020
 */
func getSeatID(seat SeatLocation) int {
	return seat.row*8 + seat.column
}

func main() {
	// get the "seat" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	seatData := *input
	seats := strings.Split(seatData, "\n")

	// convert our seats list into an array of seat IDs
	seatIDs := []int{}
	for _, seat := range seats {
		loc := getSeatLocation(seat)
		id := getSeatID(loc)
		seatIDs = append(seatIDs, id)
	}
	sort.Ints(seatIDs)

	// search for a gap in our list
	prev := seatIDs[0]
	for _, cur := range seatIDs[1:] {
		if cur-prev > 1 {
			fmt.Println(cur - 1)
			return
		}
		prev = cur
	}

}
