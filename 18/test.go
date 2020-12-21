package main

import (
	"fmt"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/18/tree"
)

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
				n := tree.MakeGroupNode(group)
				newT, ok := t.Add(n)
				if !ok {
					fmt.Sprintln("error adding group %v", group)
				}
				t = newT
				inputs = remaining
			}
		}
	}
	return t, inputs
}

func main() {

	inputs := []interface{}{2, "*", 3, "+", "(", 4, "*", 5, ")"}
	t, _ := parseTokens(inputs)
	fmt.Println(t.Evaluate())
}
