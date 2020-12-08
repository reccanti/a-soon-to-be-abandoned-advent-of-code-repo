package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

type BagContainsMap = map[string]map[string]int
type BagContainedByMap = map[string][]string

// ^(\w+ \w+) bags contain (.*)$
// ^(\d+) (\w+ \w+) bags\W?$

func parseBagString(bag string) (string, map[string]int, error) {
	// initial parse. Get the type of the "containing" bag,
	containerExp := regexp.MustCompile(`^(\w+ \w+) bags contain (.*)$`)
	res := containerExp.FindStringSubmatch(bag)
	if len(res) != 3 {
		return "", nil, errors.New("unable to find 'containing' bag in bag string")
	}
	container := res[1]
	containedStringBlock := res[2]

	// second parse, for each comma-separated "contained" string,
	// extract the quantity and type of each bag
	containedStrings := strings.Split(containedStringBlock, ",")
	containedExp := regexp.MustCompile(`^\W*(\d+) (\w+ \w+) bags?\W?$`)
	contained := map[string]int{}
	for _, str := range containedStrings {
		res := containedExp.FindStringSubmatch(str)
		if len(res) != 3 {
			break
		}
		bagType := res[2]
		quantity, err := strconv.Atoi(res[1])
		if err != nil {
			break
		}
		contained[bagType] = quantity
	}

	// return the parsed values
	return container, contained, nil
}

func main() {
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	bagText := *input
	bagStrings := strings.Split(bagText, "\n")

	/**
	 * STEP 1: Create a map of the types and quantity of bags that
	 * each bag contains
	 */
	containerBags := make(BagContainsMap)
	for _, str := range bagStrings {
		container, containing, err := parseBagString(str)
		if err != nil {
			fmt.Println(fmt.Errorf("Error parsing the string %s", str))
			return
		}
		containerBags[container] = containing
	}
	// fmt.Println(containerBags)

	/**
	 * STEP 2: Create a second map that inverts this, showing which
	 * types of bags this current bag is contained by
	 */
	containingBags := make(BagContainedByMap)
	for bagType, _ := range containerBags {
		containingBags[bagType] = []string{}
	}
	for bagType, containing := range containerBags {
		for bag, _ := range containing {
			// fmt.Println(containing)
			containedBy := containingBags[bag]
			containingBags[bag] = append(containedBy, bagType)
		}
	}
	// fmt.Println(containingBags)

	/**
	 * STEP 3: Using these 2 maps, determine which bags can contain
	 * a "shiny golden" bag
	 */
	toVisit := containingBags["shiny gold"]
	visited := map[string]bool{}
	contains := 0
	for len(toVisit) > 0 {
		currentBag := toVisit[0]
		hasVisited := visited[currentBag]
		if !hasVisited {
			toVisit = append(toVisit[1:], containingBags[currentBag]...)
			visited[currentBag] = true
			contains += 1
		} else {
			toVisit = toVisit[1:]
		}
	}
	fmt.Println(contains)
}
