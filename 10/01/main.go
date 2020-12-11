package main

import (
	"fmt"
	"os"
	"sort"
	// "strconv"
	// "strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	filename := os.Args[1]
	jolts, err := util.ParseRelativeFileInts(filename, "\n")
	if err != nil {
		fmt.Println(fmt.Errorf("unable to parse file"))
		return
	}

	// sort all of our joltage devices
	sort.Ints(jolts)

	// iterate through the list and add the first element
	// that is within 3 jolts of the current joltage
	curJoltage := 0
	visited := []int{}
	for _, jolt := range jolts {
		if jolt <= curJoltage+3 {
			visited = append(visited, jolt)
			curJoltage = jolt
			jolts = jolts[1:]
		} else {
			fmt.Println(fmt.Errorf("Cannot make a chain of all input devices"))
			return
		}
	}
	final := visited[len(visited)-1] + 3
	visited = append(visited, final)
	// fmt.Println(visited)

	// get the number of 1-jolt differences
	oneJolt := 0
	prev := 0
	for _, cur := range visited {
		if cur-prev == 1 {
			oneJolt += 1
		}
		prev = cur
	}
	// fmt.Println(oneJolt)

	// get the number of 3-jolt differences
	threeJolts := 0
	prev = 0
	for _, cur := range visited {
		if cur-prev == 3 {
			threeJolts += 1
		}
		prev = cur
	}
	// fmt.Println(threeJolts)

	fmt.Println(oneJolt * threeJolts)

	// curJoltage := 0
	// visited := []int{}
	// for len(jolts) > 0 {
	//
	// }

	// sort all inputs into a "toVisit" list
	// create an empty "visited" list
	// iterate through the lists and attempt to
}
