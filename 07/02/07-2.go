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

func countBags(currentBag string, bagMap BagContainsMap) int {
	bags, containsBags := bagMap[currentBag]
	if !containsBags {
		fmt.Println(fmt.Errorf("'%s' bag not found", currentBag))
		return 1
	}
	totalBags := 0
	for bag, quantity := range bags {
		totalBags += quantity * (countBags(bag, bagMap) + 1)
	}
	return totalBags
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

	/**
	 * STEP 2: Recursively determine the total number of bags that
	 * "shiny gold" bag can contain
	 */
	count := countBags("shiny gold", containerBags)
	fmt.Println(count)
}
