package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestMaskFunctions(t *testing.T) {
	maskStr := "000000000000000000000000000000X1001X"
	ignoreMask := constructIgnoreMask(maskStr)
	branches := constructBranches(maskStr)

	expectedBranches := []int{
		0b000000000000000000000000000000011010,
		0b000000000000000000000000000000011011,
		0b000000000000000000000000000000111010,
		0b000000000000000000000000000000111011,
	}

	fmt.Println("Masked")
	fmt.Println(fmt.Sprintf("%036b", ignoreMask))
	fmt.Println(fmt.Sprintf("%036b", 42))

	paths := []int{}
	for _, b := range branches {
		paths = append(paths, 42&ignoreMask|b)
	}

	sort.Ints(paths)

	fmt.Println("")
	fmt.Println("Expected Branches")
	for i := 0; i < len(paths)-1; i++ {
		if paths[i] != expectedBranches[i] {
			t.Errorf("given paths don't match\n%v\n%v", paths[i], expectedBranches[i])
		}
	}
}
