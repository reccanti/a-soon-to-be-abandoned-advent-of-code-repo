package main

import (
	"fmt"
	// "math"
	"os"
	// "sort"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type tuple = []int

func findMatchingRemainder(t1 tuple, t2 tuple) tuple {
	// keep incrementing according to the formula until we find
	// a common value that matches both remainders
	fmt.Println(t1)
	fmt.Println(t2)
	i := 0
	tNew := tuple{}
	for {
		val := t1[0]*i + t1[1]
		fmt.Println(val % t2[0])
		if val%t2[0] == t2[1] {
			tNew = append(tNew, t1[0]*t2[0], val)
			return tNew
		}
		i++
	}
	// return tuple{0, 0}
}

/**
 * Naive implementation to find the modular inverse. Given a
 * value A and a modulo M, increment from 0 to M until we get
 * a value where A * B % M = 1
 */
func findModularInverse(a int, m int) int {
	for b := 0; b < m; b++ {
		if (a*b)%m == 1 {
			return b
		}
	}
	return -1
}

func main() {
	filename := os.Args[1]
	lines, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		return
	}
	buses := map[int]int{}
	for i, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			num, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			buses[num] = i
		}
	}

	remainders := []tuple{}
	for key, val := range buses {
		t := tuple{key, (key - val) % key}
		remainders = append(remainders, t)
	}

	// I'm still not entirely sure how this works, but I followed
	// these instructions:
	//
	// http://homepages.math.uic.edu/~leon/mcs425-s08/handouts/chinese_remainder.pdf

	// first, we're going to get the combined product
	// of all of our stuff!
	total := 1
	for _, r := range remainders {
		total *= r[0]
	}

	// now, for each remainder, we're going to divide
	// the total by our remainder
	zs := map[int]int{}
	for _, r := range remainders {
		zs[r[0]] = total / r[0]
	}

	// next, we're going to find the inverse of each of
	// these values for our desired remainders
	ys := map[int]int{}
	for _, r := range remainders {
		inv := findModularInverse(zs[r[0]], r[0])
		if inv == -1 {
			fmt.Println("something has gone horribly wrong")
			return
		}
		ys[r[0]] = inv
	}

	// now we'll multiply our respective zs and ys together
	// this will create a value where zi * yi % mi = 1,
	// but zi * yi % mj = 0 for all other values j
	ws := map[int]int{}
	for _, r := range remainders {
		w := ys[r[0]] * zs[r[0]] % total
		ws[r[0]] = w
	}

	// now we'll calculate the final value
	value := 0
	for _, r := range remainders {
		value += ws[r[0]] * r[1]
	}

	fmt.Println(value % total)
}
