package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/16/logic"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

/**
 * Rule shit
 */

/**
 *         +----+
 *	      +----+ |
 *		  |    | +
 *		  +----+
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
 * Here we go! Let's deduce this bad boy!
 * Also, everything here is probably inefficient as hell
 */
func deduce(t logic.Table) (logic.Table, bool) {

	numRows := len(t.Rows)
	numCols := len(t.Columns)

	// isSolved := false
	// for !isSolved {
	// first, let's eliminate any entries we can
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if t.IsValid(i, j) && !t.IsSolved(i, j) {
				// fields, rule := t.Get(i, j)
				cell := t.Get(i, j)
				fields := cell.RowValue.([]int)
				rule := cell.ColumnValue.(Rule)
				for _, f := range fields {
					if !(rule.acceptedValues[f]) {
						t.MarkInvalid(i, j)
					}
				}
			}
		}
	}
	// now, let's see if we can solve any rows or columns
	for i := 0; i < numRows; i++ {
		unsolved := t.GetUnsolvedRow(i)
		if len(unsolved) == 1 {
			t.MarkSolved(unsolved[0].Row, unsolved[0].Column)
		}
	}

	_, hasSolution := t.GetSolution()
	if hasSolution {
		return t, true
	} else {
		return t, false
	}
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
	myTicketBlock := blocks[1]
	nearTicketsBlock := blocks[2]

	rules, err := parseRules(ruleBlock)
	if err != nil {
		fmt.Println(err)
		return
	}

	tstrs := strings.Split(myTicketBlock, "\n")
	myFields, err := parseFields(tstrs[1])
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

	// fuck it. Convert it all to interfaces
	potentialFieldsI := []interface{}{}
	for _, f := range potentialFields {
		potentialFieldsI = append(potentialFieldsI, f)
	}
	rulesI := []interface{}{}
	for _, r := range rules {
		rulesI = append(rulesI, r)
	}

	t := logic.NewTable(potentialFieldsI, rulesI)
	hasSolution := false
	tm.MoveCursor(1, 1)
	tm.Clear()
	tm.Println(t)
	tm.Flush()
	for !hasSolution {
		t, hasSolution = deduce(t)
		tm.MoveCursor(1, 1)
		tm.Clear()
		tm.Println(t)
		tm.Flush()

		time.Sleep(time.Second / 2)
	}

	// calculate the solution
	solutionCells, _ := t.GetSolution()
	departureIndices := []int{}
	for _, c := range solutionCells {
		if strings.Contains(c.ColumnValue.(Rule).name, "departure") {
			departureIndices = append(departureIndices, c.Row)
		}
	}
	// fmt.Println(departureIndices)

	product := 1
	for _, i := range departureIndices {
		product *= myFields[i]
	}

	fmt.Println(product)

}
