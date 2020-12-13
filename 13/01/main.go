package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	filename := os.Args[1]
	lines, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		return
	}
	timestamp, err := strconv.ParseFloat(lines[0], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	buses := []int{}
	for _, val := range strings.Split(lines[1], ",") {
		if val != "x" {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return
			}
			buses = append(buses, num)
		}
	}
	// fmt.Println(timestamp)
	// fmt.Println(buses)
	times := map[float64]int{}
	keys := []float64{}
	for _, id := range buses {
		iterations := math.Floor(timestamp/float64(id)) + 1
		waitDuration := float64(id)*iterations - timestamp
		times[waitDuration] = id
		keys = append(keys, waitDuration)
	}
	sort.Float64s(keys)
	wait := times[keys[0]]
	fmt.Println(keys[0] * float64(wait))
}
