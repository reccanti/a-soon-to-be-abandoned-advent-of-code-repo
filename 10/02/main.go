package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type Graph = map[int][]int

// this will construct a map of all the potential nodes we can traverse
func constructGraph(jolts []int) Graph {
	// add "0" to the beginning of the list
	jolts = append([]int{0}, jolts...)

	// first, construct a lookup table. We'll use this
	// to construct our nodes
	lookup := map[int]bool{}
	lookup[0] = true
	max := 0
	for _, jolt := range jolts {
		if jolt > max {
			max = jolt
		}
		lookup[jolt] = true
	}
	lookup[max+3] = true

	// now, we'll create our graph, which will consist of a "node"
	// and an array of potential nodes it can travel to
	graph := map[int][]int{}
	for _, jolt := range jolts {
		potentials := []int{}
		for i := 1; i <= 3; i++ {
			ok := lookup[jolt+i]
			if ok {
				potentials = append(potentials, jolt+i)
			}
		}
		graph[jolt] = potentials
	}
	return graph
}

func countSuccessfulPaths(graph Graph) int {
	orderedKeys := []int{}
	for key, _ := range graph {
		orderedKeys = append(orderedKeys, key)
	}
	sort.Ints(orderedKeys)

	// create a compounding number paths. Record the current number
	// of paths, then add that to the total number of paths for each
	// potential node. By the time we make it to the end of the list,
	// we'll have counted all the potential ways we could have reached
	// this node.
	pathsPerNode := map[int]int{}
	pathsPerNode[0] = 1
	for _, key := range orderedKeys {
		paths := graph[key]
		curPaths := pathsPerNode[key]
		for _, path := range paths {
			_, hasPath := pathsPerNode[path]
			if !hasPath {
				pathsPerNode[path] = 0
			}
			pathsPerNode[path] += curPaths
		}
	}
	return pathsPerNode[orderedKeys[len(orderedKeys)-1]+3]
}

func main() {
	filename := os.Args[1]
	jolts, err := util.ParseRelativeFileInts(filename, "\n")
	if err != nil {
		fmt.Println(fmt.Errorf("unable to parse file"))
		return
	}
	jolts = append([]int{0}, jolts...)

	// create a graph of all the paths we can take
	graph := constructGraph(jolts)

	// begin exhaustively searching our graph
	count := countSuccessfulPaths(graph)
	fmt.Println(count)
}
