package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/18/tree"
	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
	"github.com/rivo/tview"
)

func makeNumber(numBuffer []rune) int {
	num := 0
	for _, n := range numBuffer {
		num = num*10 + int(n-'0')
	}
	return num
}

func parseString(input string) []interface{} {
	removeWhitespace := strings.ReplaceAll(input, " ", "")

	numBuffer := []rune{}
	tokens := []interface{}{}
	for _, char := range removeWhitespace {
		if unicode.IsDigit(char) {
			numBuffer = append(numBuffer, char)
		} else {
			if len(numBuffer) > 0 {
				num := makeNumber(numBuffer)
				tokens = append(tokens, num)
				numBuffer = []rune{}
			}
			tokens = append(tokens, string(char))
		}
	}
	if len(numBuffer) > 0 {
		num := makeNumber(numBuffer)
		tokens = append(tokens, num)
	}
	return tokens
}

func parseTokens(inputs []interface{}) (tree.Node, []interface{}) {
	t := tree.MakeEmpty()
	for len(inputs) > 0 {
		in := inputs[0]
		inputs = inputs[1:]

		switch in.(type) {
		case int:
			l := tree.MakeLiteral(in.(int))
			newT, ok := t.Add(l)
			if !ok {
				fmt.Sprintln("error adding literal %v", l)
			}
			t = newT
		case string:
			switch in.(string) {
			case "*":
				o := tree.MakeTimes()
				newT, ok := t.Add(o)
				if !ok {
					fmt.Sprintln("error adding operator %v", o)
				}
				t = newT
			case "+":
				o := tree.MakePlus()
				newT, ok := t.Add(o)
				if !ok {
					fmt.Sprintln("error adding operator %v", o)
				}
				t = newT
			case ")":
				return t, inputs
			case "(":
				group, remaining := parseTokens(inputs)
				// fmt.Println(n)
				newT, ok := t.Add(group)
				if !ok {
					fmt.Sprintln("error adding group %v", group)
				}
				// fmt.Println(newT)
				t = newT
				inputs = remaining
			}
		}
	}
	return t, inputs
}

func main() {
	// initialize the application
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	// get the input
	filename := os.Args[1]
	inputs, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	i := 0
	for _, expression := range inputs {
		tokens := parseString(expression)
		tree, _ := parseTokens(tokens)
		res := tree.Evaluate()
		sum += res

		numstr := strconv.Itoa(res)

		expCell := tview.NewTableCell(expression)
		numCell := tview.NewTableCell(numstr)
		table.SetCell(i, 0, expCell)
		table.SetCell(i, 1, numCell)

		i++
	}

	sumstr := strconv.Itoa(sum)
	sumLabelCell := tview.NewTableCell("Sum")
	sumCell := tview.NewTableCell(sumstr)
	table.SetCell(i, 0, sumLabelCell)
	table.SetCell(i, 1, sumCell)
	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
