// package main

// import (
// 	"fmt"
// 	"testing"
// )

// // func TestGetNeighborInfo(t *testing.T) {
// // 	gridStr := "" +
// // 		"LLL\n" +
// // 		"#..\n" +
// // 		"..."
// // 	g := parseGrid(gridStr)
// // 	info := getNeighborInfo(g, 1, 1)
// // 	if info.numEmpty != 3 {
// // 		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'empty' neighbors. Should be '3' but instead is '%d'", info.numEmpty))
// // 	}
// // 	if info.numOccupied != 1 {
// // 		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'occupied' neighbors. Should be '1' but instead is '%d'", info.numOccupied))
// // 	}
// // 	if info.numFloor != 4 {
// // 		t.Errorf(fmt.Sprintf("improperly calculated the nuber of 'floor' neighbors. Should be '4' but instead is '%d'", info.numFloor))
// // 	}
// // }

// func TestGetOccupiedVisible(t *testing.T) {
// 	gridStr := "" +
// 		"LLL#L\n" +
// 		"#..L#\n" +
// 		".L...\n" +
// 		".....\n" +
// 		"L#LLL"
// 	g := parseGrid(gridStr)
// 	for r := 0; r < g.Rows; r++ {
// 		for c := 0; c < g.Columns; c++ {
// 			num, err := getOccupiedVisible(g, r, c)
// 			fmt.Println(r, c, "-- START --")
// 			fmt.Println("Occupied Visible:", *num)
// 			fmt.Println(g)
// 			// fmt.Println(err)
// 			if err == nil {
// 				t.Errorf("should be out-of-bounds")
// 				// return
// 			}
// 			if err == nil {
// 				fmt.Println(*num)
// 			}
// 		}
// 	}
// 	t.Errorf("uh")
// }
