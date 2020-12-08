// package main

// import (
// 	"testing"
// )

// func TestGetSeatLocation(t *testing.T) {
// 	loc := getSeatLocation("FBFBBFFRLR")
// 	if loc.row != 44 {
// 		t.Errorf("row calculated incorrectly")
// 	}
// 	if loc.column != 5 {
// 		t.Errorf("column calculated incorrectly")
// 	}
// }

// func TestGetSeatID(t *testing.T) {
// 	loc := getSeatLocation("FBFBBFFRLR")
// 	id := getSeatID(loc)
// 	if id != 357 {
// 		t.Errorf("id calculated incorrectly")
// 	}
// }

package main

import (
	"fmt"
	"testing"
)

func TestParseBagString(t *testing.T) {
	bagString := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	container, containing, err := parseBagString(bagString)
	if err != nil {
		t.Errorf("should properly parse bag string")
	}
	if container != "light red" {
		t.Errorf("'container' improperly parsed")
	}
	fmt.Println(containing)
	if len(containing) != 2 || containing["bright white"] != 1 || containing["muted yellow"] != 2 {
		t.Errorf("'containing' map improperly parsed")
	}
}
