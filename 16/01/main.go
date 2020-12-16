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

/**
 * Rule shit
 */

// (.+): (\d+-\d+) or (\d+-\d+)
type Rule struct {
	name           string
	acceptedValues map[int]bool
}

func parseRules(block string) ([]Rule, error) {
	exp := regexp.MustCompile(`(.+): (\d+-\d+) or (\d+-\d+)`)
	ruleStrs := strings.Split(block, "\n")
	rules := []Rule{}
	for _, r := range ruleStrs {
		res := exp.FindStringSubmatch(r)
		if len(res) != 4 {
			return nil, errors.New("something has gone terribly wrong parsing the rules")
		}
		name := res[1]
		values := res[2:]

		r, err := constructRule(name, values)
		if err != nil {
			return nil, err
		}

		rules = append(rules, *r)
	}

	return rules, nil
}

func constructRule(name string, values []string) (*Rule, error) {
	// convert our number strings to ints
	acceptedValues := map[int]bool{}
	for _, v := range values {
		numStrings := strings.Split(v, "-")
		vals := []int{}
		for _, s := range numStrings {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("soething has gone terribly wrong constructing the rules")
				return nil, err
			}
			vals = append(vals, n)
		}
		// construct our intermediate values
		for i := vals[0]; i <= vals[1]; i++ {
			acceptedValues[i] = true
		}
	}

	r := Rule{
		name:           name,
		acceptedValues: acceptedValues,
	}

	return &r, nil
}

/**
 * Field shit
 */
func parseFields(fieldStr string) ([]int, error) {
	numStrings := strings.Split(fieldStr, ",")
	nums := []int{}
	for _, s := range numStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

/**
 * The thing where the goddamn app runs
 */

func main() {
	// get input
	filename := os.Args[1]
	blocks, err := util.ParseRelativeFileSplit(filename, "\n\n")
	if err != nil {
		return
	}

	// parse all the goddamn fields

	ruleBlock := blocks[0]
	// myTicketBlock := blocks[1]
	nearTicketsBlock := blocks[2]

	rules, err := parseRules(ruleBlock)
	if err != nil {
		fmt.Println(err)
		return
	}

	ntstrs := strings.Split(nearTicketsBlock, "\n")
	nearFields := [][]int{}
	for _, s := range ntstrs[1:] {
		f, err := parseFields(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		nearFields = append(nearFields, f)
	}

	// put all of our rules into a common pool of rule ranges

	allRules := map[int]bool{}
	for _, r := range rules {
		for k, v := range r.acceptedValues {
			allRules[k] = v
		}
	}

	// check to see if any of the values are invalid

	invalidFields := []int{}
	for _, field := range nearFields {
		for _, f := range field {
			if !allRules[f] {
				invalidFields = append(invalidFields, f)
			}
		}
	}

	// sum 'em together

	sum := 0
	for _, i := range invalidFields {
		sum += i
	}
	fmt.Println(sum)
}
