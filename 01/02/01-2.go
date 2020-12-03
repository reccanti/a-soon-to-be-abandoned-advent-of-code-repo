package main

import (
	"errors"
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

/**
 * For this, we'd need to do the following:
 *
 * 1. Create N iterators, where "N" is the amount of values we want
 *    to sum together to get the "target" value (i.e. if we want to
 * 	  find 3 values that sum to 100, we'd create "3" iterators)
 * 2. Place these iterators at separate values in the beginning of
 * 	  the list (so if we have 3 iterators, they would index "0", "1", and "2")
 * 3. Start incrementing these values. We'll need to follow these rules:
 * 	  a. If the current sum is less than the target, increment the current
 *       iterator.
 *    b. If the current sum is greater than the target, decrement the
 *		 next iterator ahead
 *	  c. An iterator must be a value between the next and previous iterators.
 *        If we can't do a or b, move down to the previous iterator
 */
func perfectSums(inputs []int, target int, numValues int) ([]int, error) {

	// basic error checking. Make sure the number of values is less than
	// the total number of potential inputs in the array
	if len(inputs) < numValues {
		fmt.Println("This won't work!")
		return nil, errors.New("Not enough values in this array for the desired number of values")
	}

	// create our initial list of iterators:
	indices := []int{}
	for i := 0; i < numValues; i += 1 {
		indices = append(indices, i)
	}

	// this loop is weird and probably bad, but I think it works???
	// It's not very elegant, but I'll do my best to annotate
	// ~reccanti 12/2/2020
	cur := 0
	for {
		// get the sum of the current index values we're recording
		sum := 0
		for _, index := range indices {
			sum += inputs[index]
		}

		// if the sum of all of our values is less than the target,
		// increment EVERY index by one. Don't worry if this puts us
		// over the target value. We'll handle that in the "greater
		// than" case.
		if sum < target {
			for i := range indices {
				indices[i] += 1
			}
		}
		// if the sum of all of our values is greater than the target,
		// we'll decrement our indices one-by-one until we reach the
		// "less than" case again.
		if sum > target {
			// EDGE CASE: if we can't increment the current value any lower,
			// decrement the last value (i.e. the greatest value)
			// by one. This will let us restart our scanning. Otherwise
			if indices[cur] == 0 {
				indices[len(indices)-1] -= 1
			} else {
				// EDGE CASE: We can't have multiple copies of the same index
				// in our returned values, if we've decremented in such a way
				// where the only possible option is to use multiple copies of
				// the same index, that means we've exhausted all of our options
				if cur > 0 && indices[cur-1] == indices[cur] {
					return nil, errors.New("No combination of values works")
				}
				indices[cur] -= 1
				cur = (cur + 1) % (len(indices) - 1)
			}
		}
		// if the sum of all our values is equal to the target, we've
		// found the values t so we return them!
		if sum == target {
			values := []int{}
			for _, i := range indices {
				values = append(values, inputs[i])
			}
			return values, nil
		}

	}
}

func main() {

	filename := os.Args[1]

	inputs, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	sort.Ints(inputs)

	values, err := perfectSums(inputs, 2020, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(values)

	product := 1
	for _, val := range values {
		product *= val
	}
	fmt.Println(product)
}
