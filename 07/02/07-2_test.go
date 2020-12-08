package main

import (
	"testing"
)

func TestGetSeatLocation(t *testing.T) {
	loc := getSeatLocation("FBFBBFFRLR")
	if loc.row != 44 {
		t.Errorf("row calculated incorrectly")
	}
	if loc.column != 5 {
		t.Errorf("column calculated incorrectly")
	}
}

func TestGetSeatID(t *testing.T) {
	loc := getSeatLocation("FBFBBFFRLR")
	id := getSeatID(loc)
	if id != 357 {
		t.Errorf("id calculated incorrectly")
	}
}
