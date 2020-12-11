package main

import (
	"fmt"
	"os"

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

func traverseStr(cur int, graph Graph, str string) {
	// paths := [][]int{}

	pathStr := fmt.Sprintf("%s->%d", str, cur)

	potentials, hasArray := graph[cur]
	if !hasArray {
		fmt.Println(pathStr)
	}
	for _, potential := range potentials {
		traverseStr(potential, graph, pathStr)
	}
}

func getSuccessfulPaths(graph Graph) [][]int {

	paths := [][]int{}
	var internalGetSuccessfulPaths func(cur int, path []int, graph Graph)
	internalGetSuccessfulPaths = func(cur int, path []int, graph Graph) {
		fmt.Println(cur)
		path = append(path, cur)
		potentials, hasArray := graph[cur]
		if !hasArray {
			fmt.Println(path)
			paths = append(paths, path)
			return
		}
		for _, potential := range potentials {
			internalGetSuccessfulPaths(potential, path, graph)
		}
	}
	internalGetSuccessfulPaths(0, []int{}, graph)
	return paths
}

func countSuccessfulPaths(graph Graph) int {

	paths := 0
	var internalGetSuccessfulPaths func(cur int, graph Graph)
	internalGetSuccessfulPaths = func(cur int, graph Graph) {
		// fmt.Println(cur)
		potentials, hasArray := graph[cur]
		if !hasArray {
			paths += 1
			fmt.Println(paths)
			return
		}
		for _, potential := range potentials {
			internalGetSuccessfulPaths(potential, graph)
		}
	}
	internalGetSuccessfulPaths(0, graph)
	return paths
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
	// fmt.Println(graph)

	// begin exhaustively searching our graph
	count := countSuccessfulPaths(graph)
	fmt.Println(count)
	// fmt.Println(len(paths))
}
