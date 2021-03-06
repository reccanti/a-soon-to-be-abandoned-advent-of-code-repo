package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func checkSum(val int, preamble []int) bool {
	// create a sorted copy of the array
	sortedPreamble := make([]int, len(preamble))
	copy(sortedPreamble, preamble)
	sort.Ints(sortedPreamble)

	// create two indexes, one at the beginning, and one at the end
	// if the sum is greater than the given value, decrement the "end" index
	// if hte sum is less than the given value, increment the "start" index
	// loop until a solution is found or the indexes are equal
	start := 0
	end := len(sortedPreamble) - 1
	for start != end {
		sum := sortedPreamble[start] + sortedPreamble[end]
		if sum > val {
			end -= 1
		} else if sum < val {
			start += 1
		} else {
			return true
		}
	}
	return false
}

// traverse the numbers list and check that each element
// is the sum of two of the previous values
func findInvalidEntry(preambleLength int, numbers []int) (*int, error) {
	window := numbers[0:preambleLength]
	for i, val := range numbers[preambleLength:] {
		if !checkSum(val, window) {
			return &val, nil
		}
		window = numbers[i+1 : i+preambleLength+1]
	}
	return nil, errors.New("all entries in the list are valid")
}

// scan the list to find contiguous values that add to a sum
func sum(ints ...int) int {
	acc := 0
	for _, i := range ints {
		acc += i
	}
	return acc
}

func findContiguousValues(target int, numbers []int) ([]int, error) {
	window := []int{}
	for _, val := range numbers {
		window = append(window, val)
		s := sum(window...)
		// if we're larger than the target, any other value will just
		// make it larger, so we'll remove entries until we're less
		// than or equal to the target value
		for s > target {
			window = window[1:]
			s = sum(window...)
		}
		if s == target {
			return window, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no sum of entries adds to %d", target))
}

func main() {
	// get the input
	filename := os.Args[2]
	preambleLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	numbersBlock := *input
	numbersStr := strings.Split(numbersBlock, "\n")
	numbers := []int{}
	for _, str := range numbersStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, num)
	}

	// find the invalid number
	invalidNumber, err := findInvalidEntry(preambleLength, numbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// find the contiguous values that sum to the given number
	values, err := findContiguousValues(*invalidNumber, numbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get the sum of the smallest and largest number in the list
	sort.Ints(values)
	smallest := values[0]
	largest := values[len(values)-1]

	fmt.Println(smallest + largest)
}
