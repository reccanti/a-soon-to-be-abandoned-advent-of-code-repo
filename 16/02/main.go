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

func (r Rule) HashKey() string {
	return r.name + "_" + string(len(r.acceptedValues))
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
	// we can't use rules as a key directly, so we construct
	// this intermediary lookup map
	stupidRuleLookup := map[string]Rule{}
	for _, r := range rules {
		stupidRuleLookup[r.HashKey()] = r
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

	// filter out invalid tickets

	validTickets := [][]int{}
	for _, field := range nearFields {
		isValid := true
		for _, f := range field {
			if !allRules[f] {
				isValid = false
			}
		}
		if isValid {
			validTickets = append(validTickets, field)
		}
	}

	// "flip" our array to the side, so that we have
	// potential values for each field
	potentialFields := make([][]int, len(rules))
	for _, ticket := range validTickets {
		for i, v := range ticket {
			potentialFields[i] = append(potentialFields[i], v)
		}
	}
	fmt.Println(potentialFields)

	// let's maintain two maps, a map of "solved" entries
	// (fields assigned to a rule), and a map of "unsolved" entries
	// (fields with several potential rules). We'll iterate over
	// this structure until we can deduce a rule. Once all the
	// rules have been deduced, we'll make note of that, narrowing
	// down the potential rules for each field to see if we can
	// deduce any more

	// Step 1: Construct an initial list
	// in this case, "i" refers to the index of an entry in the potential
	// fields list and will be used to look it up
	unsolvedFields := map[int][]Rule{}
	for i, fieldValues := range potentialFields {

		// create a lookup chart for potential rules. Everything
		// is true by default. We'll remove values that we can
		// rule out
		potentialRulesLookup := map[string]bool{}
		for _, r := range rules {
			potentialRulesLookup[r.HashKey()] = true
		}

		for _, val := range fieldValues {
			for hashkey, _ := range potentialRulesLookup {
				r := stupidRuleLookup[hashkey]
				if !r.acceptedValues[val] {
					potentialRulesLookup[r.HashKey()] = false
					break
				}
			}
		}

		potentialRules := []Rule{}
		for hashkey, isValid := range potentialRulesLookup {
			if isValid {
				r := stupidRuleLookup[hashkey]
				potentialRules = append(potentialRules, r)
			}
		}

		unsolvedFields[i] = potentialRules
	}
	for _, v := range unsolvedFields {
		fmt.Println(v)
		fmt.Println(len(v))
	}

	// sum 'em together

	// sum := 0
	// for _, i := range invalidFields {
	// 	sum += i
	// }
	// fmt.Println(sum)
}
