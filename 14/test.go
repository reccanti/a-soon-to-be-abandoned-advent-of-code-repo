package main

import (
	"fmt"
	"sort"
)

func attempt1() {
	// construct a "binary" number from a given value
	testString := "000000000000000000000000000000110011"
	base := 1
	num := 0
	for i := len(testString) - 1; i >= 0; i-- {
		char := string(testString[i])
		if char == "1" {
			num += 1 * base
		}
		base *= 2
	}
	fmt.Println(num)
	fmt.Println(fmt.Sprintf("%036b", num))
}

func attempt2() {
	// construct a "main" binary number from a given value.
	// if we encounter an "X" we log it as a pair of "floaters",
	// which will be a "0" and "1" variant of the given index
	testString := "00000000000000000000000000000000X0XX"
	base := 1
	main := 0
	floaters := [][]int{}
	for i := len(testString) - 1; i >= 0; i-- {
		char := string(testString[i])
		if char == "1" {
			main += 1 * base
		}
		if char == "X" {
			floaters = append(floaters, []int{0 * base, 1 * base})
		}
		base *= 2
	}

	// given our floaters from before, construct some
	// branching paths.
	branches := []int{0}
	for _, f := range floaters {
		newBranches := []int{}
		for _, b := range branches {
			newBranches = append(newBranches, b+f[0])
			newBranches = append(newBranches, b+f[1])
		}
		branches = newBranches
	}

	// construct all of the addresses
	addresses := []int{}
	for _, b := range branches {
		addresses = append(addresses, main+b)
	}

	sort.Ints(addresses)

	for _, a := range addresses {
		fmt.Println(fmt.Sprintf("%036b", a))
	}

}

func main() {
	attempt1()
	fmt.Println("")
	attempt2()
}
