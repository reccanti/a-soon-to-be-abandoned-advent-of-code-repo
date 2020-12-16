package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/11/grid"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

// utility function to parse the input into a grid
func parseGrid(str string) grid.Grid {
	strs := strings.Split(str, "\n")
	numRows := len(strs)
	numColumns := len(strs[0])
	cells := []interface{}{}
	for _, rowStr := range strs {
		for _, colStr := range rowStr {
			cells = append(cells, string(colStr))
		}
	}
	return grid.New(numRows, numColumns, cells)
}

type neighborInfo struct {
	numOccupied int
	numEmpty    int
	numFloor    int
}

func getOccupiedVisible(g grid.Grid, row int, column int) (*int, error) {
	occupied := 0
	if !g.IsInBounds(row, column) {
		return nil, errors.New("row and column is out-of-bounds")
	}
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
			} else {
				// iterate until we reach an occupied seat
				// or until we reach the end of the grid
				mod := 1
				for {
					curRow := r*mod + row
					curCol := c*mod + column
					if !g.IsInBounds(curRow, curCol) {
						break
					}
					// fmt.Println(curRow, curCol)
					cell, _ := g.Get(curRow, curCol)
					if *cell == "L" {
						break
					} else if *cell == "#" {
						// fmt.Println("got one")
						occupied += 1
						break
					}
					mod++
				}
			}
		}
	}
	// fmt.Println(occupied)
	return &occupied, nil
}

// determine the next tile based on neighbor iteration
func getNewTile(g grid.Grid, row int, column int) (interface{}, error) {
	cell, err := g.Get(row, column)
	if err != nil {
		return nil, err
	}
	occupied, err := getOccupiedVisible(g, row, column)
	if err != nil {
		return nil, err
	}
	// see if we should flip an empty cell
	if *cell == "L" {
		if *occupied == 0 {
			newStr := "#"
			return newStr, nil
		} else {
			return *cell, nil
		}
	} else if *cell == "#" {
		if *occupied >= 5 {
			newStr := "L"
			return newStr, nil
		} else {
			return *cell, nil
		}
	}
	return *cell, nil
}

// get the next iteration of the grid
func iterate(g grid.Grid) (grid.Grid, int) {
	newCells := []interface{}{}
	changes := 0
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Columns; col++ {
			oldCell, getErr := g.Get(row, col)
			newCell, newErr := getNewTile(g, row, col)
			if getErr != nil || newErr != nil {
				fmt.Println(fmt.Errorf("Something went terribly wrong"))
				break
			}
			if *oldCell != newCell {
				changes += 1
			}
			newCells = append(newCells, newCell)
		}
	}
	return grid.New(g.Rows, g.Columns, newCells), changes
}

// count the number of
func countOccupiedSeats(g grid.Grid) int {
	occupiedSeats := 0
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Columns; col++ {
			cell, err := g.Get(row, col)
			if err != nil {
				fmt.Println(fmt.Errorf("Something went terribly wrong"))
				break
			}
			if *cell == "#" {
				occupiedSeats += 1
			}
		}
	}
	return occupiedSeats
}

func main() {
	// get the input
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	gridBlock := *input
	g := parseGrid(gridBlock)

	// print initial state
	tm.MoveCursor(1, 1)
	tm.Clear()
	tm.Println(g)
	tm.Println("Occupied Seats:")
	tm.Flush()

	// loop until the number of changes is zero
	numChanges := 0
	g, numChanges = iterate(g)
	for numChanges > 0 {
		tm.MoveCursor(1, 1)
		tm.Clear()
		tm.Println(g)
		tm.Println("Occupied Seats:")
		tm.Flush()

		g, numChanges = iterate(g)
		time.Sleep(time.Second / 20)
	}

	// print output
	occupiedSeats := countOccupiedSeats(g)
	tm.MoveCursor(1, 1)
	tm.Clear()
	tm.Println(g)
	tm.Println("Occupied Seats:", occupiedSeats)
	tm.Flush()
}
