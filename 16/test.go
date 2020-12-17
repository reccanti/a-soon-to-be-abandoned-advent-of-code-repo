package main

import (
	"fmt"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/16/logic"
)

func main() {
	r := []interface{}{}
	r = append(r, 1, 2, 3, 4, 5)
	c := []interface{}{}
	c = append(c, "one", "two", "three", "four", "five")

	t := logic.NewTable(r, c)

	// Getters
	fmt.Println(t.Get(2, 2))
	fmt.Println(t.GetRow(2))
	fmt.Println(t.GetColumn(2))

	// Invalid funcs
	t.MarkInvalid(2, 2)
	fmt.Println(t.IsValid(2, 2))
	fmt.Println(t.IsValid(2, 3))

	// Solved funcs
	t.MarkSolved(3, 3)

	// "Unsolved" funcs
	fmt.Println(t.GetUnsolvedRow(2))
	fmt.Println(t.GetUnsolvedColumn(4))

	// test the damn table
	fmt.Println(t)
}
