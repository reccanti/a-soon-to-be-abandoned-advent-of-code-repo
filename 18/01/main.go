package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
	"github.com/rivo/tview"
)

// we need to define the types of things that can exist in our
// operation string. Right now, I'm thinking of the following:
//
// 1. Literals - only one value
// 2. Expressions - combination of 2 values and an operation
//
// All expressions must be solved and reduced to literals before
// we attempt to solve them!

// tokens

type NumberToken struct {
	value int
}

type OperatorToken struct {
	value string
}

type GroupToken struct {
	value string
}

// AST nodes

type Literal struct {
	value int
}

type Expression struct {
	operation *string
	left      interface{}
	right     interface{}
}

func MakeExp() Expression {
	return Expression{
		operation: nil,
		left:      nil,
		right:     nil,
	}
}

func (e Expression) String() string {
	return fmt.Sprintf("(%v %s %v)", e.left, *e.operation, e.right)
}

func (e Expression) IsComplete() bool {
	if e.left != nil && e.right != nil && e.operation != nil {
		return true
	}
	return false
}

func (e Expression) Add(t interface{}) (*Expression, error) {
	switch t.(type) {
	case OperatorToken:
		if e.operation != nil {
			return nil, errors.New("Unable to parse expression")
		} else {
			operator := t.(OperatorToken).value
			e.operation = &operator
			return &e, nil
		}
	case NumberToken:
		l := Literal{
			value: t.(NumberToken).value,
		}
		if e.left == nil {
			e.left = l
		} else if e.right == nil {
			e.right = l
		} else {
			return nil, errors.New("Unable to add literal to expression")
		}
		return &e, nil
	}
	return nil, errors.New(fmt.Sprintf("Did not recognize token %v", t))
}

// parsing operations

func tokenize(input string) []interface{} {
	numBuffer := []rune{}
	tokens := []interface{}{}
	for _, cur := range input {
		if unicode.IsDigit(cur) {
			numBuffer = append(numBuffer, cur)
		} else if cur == '*' || cur == '+' || cur == '(' || cur == ')' {
			if len(numBuffer) > 0 {
				num := 0
				for _, n := range numBuffer {
					num = num*10 + int(n-'0')
				}
				t := NumberToken{
					value: num,
				}
				tokens = append(tokens, t)

			}
			numBuffer = []rune{}
			if cur == '*' || cur == '+' {
				t := OperatorToken{
					value: string(cur),
				}
				tokens = append(tokens, t)
			} else {
				t := GroupToken{
					value: string(cur),
				}
				tokens = append(tokens, t)
			}
		}
	}
	if len(numBuffer) > 0 {
		num := 0
		for _, n := range numBuffer {
			num = num*10 + int(n-'0')
		}
		t := NumberToken{
			value: num,
		}
		tokens = append(tokens, t)
	}
	return tokens
}

func buildAST(tokens []interface{}) (*Expression, []interface{}, error) {
	e := MakeExp()
	for len(tokens) > 0 {
		t := tokens[0]
		tokens = tokens[1:]

		switch t.(type) {
		case GroupToken:
			if t.(GroupToken).value == "(" {
				g, remainingTokens, err := buildAST(tokens)
				if err != nil {
					return nil, nil, err
				}
				if e.left == nil {
					e.left = *g
				} else if e.right == nil {
					e.right = *g
				} else {
					return nil, nil, errors.New("Can't add group to expression")
				}
				tokens = remainingTokens
			} else {
				return &e, tokens, nil
			}
		default:
			if e.IsComplete() {
				newExp := MakeExp()
				newExp.left = e
				e = newExp
			}
			newExp, err := e.Add(t)
			if err != nil {
				// fmt.Println(err)
				return nil, nil, err
			}
			e = *newExp
		}
	}
	return &e, tokens, nil
}

func parseExpression(opstr string) (*Expression, error) {
	opstr = strings.ReplaceAll(opstr, " ", "")

	// first, we'll "tokenize" our input. Our tokens will be one of these:
	// 1. Number
	// 2. Operator
	// 3. Grouping
	tokens := tokenize(opstr)

	// now we'll parse our tokens into an AST.
	e, _, err := buildAST(tokens)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Evaluation functions

func evaluate(e Expression) int {

	leftVal := 0
	rightVal := 0

	// get the left value
	switch e.left.(type) {
	case Expression:
		leftVal = evaluate(e.left.(Expression))
	case Literal:
		leftVal = e.left.(Literal).value
	}

	// get the right value
	switch e.right.(type) {
	case Expression:
		rightVal = evaluate(e.right.(Expression))
	case Literal:
		rightVal = e.right.(Literal).value
	}

	// combine these fuckers!
	opptr := e.operation
	op := *opptr
	if op == "*" {
		return leftVal * rightVal
	} else {
		return leftVal + rightVal
	}
}

func main() {
	// initialize the application
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	// get the input
	filename := os.Args[1]
	inputs, err := util.ParseRelativeFileSplit(filename, "\n")
	if err != nil {
		return
	}

	// parse all of our inputs and add them to the list
	// list.AddItem("Quit", "Press to exit", 'q', func() {
	// 	app.Stop()
	// })
	sum := 0
	i := 0
	for _, expstr := range inputs {
		expptr, err := parseExpression(expstr)
		if err != nil {
			fmt.Println(err)
			return
		}
		exp := *expptr
		num := evaluate(exp)
		numstr := strconv.Itoa(num)

		expCell := tview.NewTableCell(expstr)
		numCell := tview.NewTableCell(numstr)
		table.SetCell(i, 0, expCell)
		table.SetCell(i, 1, numCell)

		sum += num
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
