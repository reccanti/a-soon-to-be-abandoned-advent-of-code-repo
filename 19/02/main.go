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

// fancy go programming. I sure hope this doesn't bite me in the ass!
// lookup := map[int](func() []string){}
// 	lookup[0] = func() []string {
// 		return []string{"a"}
// 	}
// 	fmt.Println(lookup[0]())

// `(\d+): (.+)`
// `"(.*)"`
func parseRules(str string) (int, interface{}, error) {
	// first things first, we want to get the "index"
	// of the rule
	indexExp := regexp.MustCompile(`(\d+): (.+)`)
	res := indexExp.FindStringSubmatch(str)
	if len(res) != 3 {
		return -1, nil, errors.New("did not parse the rule string correctly")
	}
	index, err := strconv.Atoi(res[1])
	if err != nil {
		return -1, nil, err
	}

	// next, we parse the remaining fields. Depending
	// on what we're parsing, it'll be either a string
	// or an array of additional indexes to check
	stringValueExp := regexp.MustCompile(`"(.*)"`)
	stringRes := stringValueExp.FindStringSubmatch(res[2])
	if len(stringRes) == 2 {
		return index, stringRes[1], nil
	}

	groups := strings.Split(res[2], "|")
	inputs := [][]int{}
	for _, g := range groups {
		combos := []int{}
		rules := strings.Split(g, " ")
		for _, r := range rules {
			if r == "" {
				continue
			}
			rule, err := strconv.Atoi(r)
			if err != nil {
				return -1, nil, err
			}
			combos = append(combos, rule)
		}
		inputs = append(inputs, combos)
	}
	return index, inputs, nil
}

/**
 * Utility function to get all the unique permutations of two
 * arrays of strings, so:
 *
 * ["a", "b"], ["a", "b"]
 *
 * would return:
 *
 * ["aa", "ab", "ba", "bb"]
 *
 * and:
 *
 * ["a", "a"], ["b", "b"]
 *
 * would return:
 *
 * ["aa", "bb"]
 */
func zip(allStrs ...[]string) []string {
	// first, create a naive list of all the
	// different combinations of strings
	naiveStrings := []string{""}
	for _, strSet := range allStrs {
		combosInSet := []string{}
		for _, baseStr := range naiveStrings {
			for _, str := range strSet {
				combosInSet = append(combosInSet, baseStr+str)
			}
		}
		naiveStrings = combosInSet
	}

	// now, we'll dedupe any elements in these arrays
	lookup := map[string]bool{}
	res := []string{}
	for _, str := range naiveStrings {
		if !lookup[str] {
			res = append(res, str)
			lookup[str] = true
		}
	}
	return res
}

// just a test main. Delete this later
// func main() {
// 	// should just return ["a"]
// 	strs1 := zip([]string{"a"})
// 	fmt.Println(strs1)

// 	strs2 := zip([]string{"a"}, []string{"a"})
// 	fmt.Println(strs2)

// 	strs3 := zip([]string{"a", "a"}, []string{"b", "b"})
// 	fmt.Println(strs3)

// 	strs4 := zip([]string{"a", "b", "ab"}, []string{"a", "b"})
// 	fmt.Println(strs4)
// }

func main() {
	// get the input
	filename := os.Args[1]
	inputs, err := util.ParseRelativeFileSplit(filename, "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	// get the rules

	ruleBlock := inputs[0]
	ruleLookup := map[int]interface{}{}
	rules := strings.Split(ruleBlock, "\n")
	for _, r := range rules {
		index, value, err := parseRules(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		ruleLookup[index] = value
	}

	// construct a lookup table, where we check against
	// previous keys until we return strings

	lookup := map[int]func() []string{}
	for k, v := range ruleLookup {
		switch v.(type) {
		case [][]int:
			values := v.([][]int)
			lookup[k] = func() []string {
				allStrings := []string{}
				for _, indices := range values {
					res := [][]string{}
					for _, i := range indices {
						res = append(res, lookup[i]())
					}
					allStrings = append(allStrings, zip(res...)...)
				}
				return allStrings
			}
		case string:
			value := v.(string)
			lookup[k] = func() []string {
				return []string{value}
			}
		}
	}

	fmt.Println(lookup[0]())
	// construct a final lookup table of all the values
	// that result from key 0

	zeroVals := map[string]bool{}
	for _, val := range lookup[0]() {
		zeroVals[val] = true
	}

	// parse all of the outputs

	outputBlock := inputs[1]
	outputVals := strings.Split(outputBlock, "\n")
	count := 0
	for _, val := range outputVals {
		if zeroVals[val] {
			count += 1
		}
	}

	// display the count

	fmt.Println(count)
}
