package main

import (
	"fmt"
	"testing"
	// "github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/11/grid"
)

func TestGetNeighborInfo(t *testing.T) {
	gridStr := "" +
		"LLL\n" +
		"#..\n" +
		"..."
	g := parseGrid(gridStr)
	info := getNeighborInfo(g, 1, 1)
	if info.numEmpty != 3 {
		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'empty' neighbors. Should be '3' but instead is '%d'", info.numEmpty))
	}
	if info.numOccupied != 1 {
		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'occupied' neighbors. Should be '1' but instead is '%d'", info.numOccupied))
	}
	if info.numFloor != 4 {
		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'floor' neighbors. Should be '4' but instead is '%d'", info.numFloor))
	}
}
