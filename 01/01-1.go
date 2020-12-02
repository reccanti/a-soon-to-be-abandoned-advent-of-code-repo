package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

// A utility function for parsing an input file into an
// array of ints.
//
// @TODO If this is a common problem later, it might make
// sense to break this down into some utility functions
// that can be reused in future problems
//
// ~reccanti 12/2/2020
func parseFile(filename string) ([]int, error) {
	// 1. get the file name
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filepath := path.Join(wd, filename)

	// 2. get the string contents of the file
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// 3. Parse that string into an array of numbers
	strs := strings.Split(string(dat), "\n")
	ints := []int{}
	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func main() {

	filename := os.Args[1]

	inputs, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	// This is probably inefficient because it
	// runs in O(n^2) time. We might be able to
	// improve this by sorting the list first
	// and iterating values on both ends
	//
	// ~reccanti 12/2/2020
	// for i, val1 := range inputs {
	// 	remaining := inputs[i+1:]
	// 	for _, val2 := range remaining {
	// 		if val1+val2 == 2020 {
	// 			fmt.Println(val1 * val2)
	// 		}
	// 	}
	// }

	/**
	 * @NOTE Here we sort our inputs and iterate from both the
	 * beginning and end of the array, i.e. a "low-value" and a
	 * "high-value" index.
	 *
	 * If the sum of those values is less than our "target" value,
	 * we need to increment our "low-value" index to the next higher
	 * value.
	 *
	 * If the sum of those values is greater than our "target" value,
	 * we need to decrement our "high-value" index to the next lower
	 * value.
	 *
	 * If our indexes are ever the same, there are no possible
	 * combination of values that will work.
	 *
	 * ~bwilcox 12/2/2020
	 */
	sort.Ints(inputs)
	i := 0
	j := len(inputs) - 1
	for {
		// if our indexes converge, we'll never produce a value of
		// 2020, so we should exit out
		if i == j {
			fmt.Println("No combination of values adds up to 2020")
			break
		}

		// get the sum of our current "low" and "high" values
		val1 := inputs[i]
		val2 := inputs[j]
		sum := val1 + val2

		// if the sum is less than 2020, increment our "low" index
		if sum < 2020 {
			i += 1
		}
		// if the sum is greater than 2020, decrement our "high" index
		if sum > 2020 {
			j -= 1
		}
		// if our sum equals 2020, print the product and exit
		if sum == 2020 {
			fmt.Println(val1 * val2)
			break
		}
	}

}
