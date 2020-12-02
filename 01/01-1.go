package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	for i, val1 := range inputs {
		remaining := inputs[i+1:]
		for _, val2 := range remaining {
			if val1+val2 == 2020 {
				fmt.Println(val1 * val2)
			}
		}
	}
}
