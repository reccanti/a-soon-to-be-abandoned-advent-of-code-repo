package main

import (
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

	// traverse the numbers list and check that each element
	// is the sum of two of the previous values
	window := numbers[0:preambleLength]
	for i, val := range numbers[preambleLength:] {
		if !checkSum(val, window) {
			fmt.Println(val)
			return
		}
		window = numbers[i+1 : i+preambleLength+1]
	}
	fmt.Println(fmt.Errorf("All items are valid sums"))
}
