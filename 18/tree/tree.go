package tree

import (
	"fmt"
)

// Literal type

type Literal struct {
	value int
}

func (l Literal) String() string {
	return fmt.Sprintf("%d", l.value)
}

func MakeLiteral(value int) Literal {
	return Literal{value}
}

// Operator Type

type Operator struct {
	operation string
	// higher priority gets precedence
	priority int
}

func (o Operator) String() string {
	return fmt.Sprintf("%s", o.operation)
}

func MakeOperator(operation string, priority int) Operator {
	return Operator{
		operation,
		priority,
	}
}

func MakeTimes() Operator {
	return Operator{
		operation: "*",
		priority:  1,
	}
}

func MakePlus() Operator {
	return Operator{
		operation: "+",
		priority:  2,
	}
}

func MakeGroup() Operator {
	return Operator{
		operation: "()",
		priority:  3,
	}
}

// Tree types

type EmptyNode struct{}

type Node struct {
	value interface{}
	next  []Node
}

func (n Node) String() string {
	switch n.value.(type) {
	case Literal:
		return fmt.Sprintf("%v ", n.value)
	default:
		next := ""
		for i, n := range n.next {
			next += fmt.Sprintf("%v", n)
			if i < len(n.next)-1 {
				next += ","
			}
		}
		return fmt.Sprintf("(%v %s)", n.value, next)
	}
}

func MakeEmpty() Node {
	return Node{
		value: EmptyNode{},
		next:  []Node{},
	}
}

func Make(value interface{}) Node {
	return Node{
		value: value,
		next:  []Node{},
	}
}

// Group type

// func MakeGroupNode(n Node) Node {
// 	next := []Node{}
// 	next = append(next, n)
// 	return Node{
// 		value: MakeGroup(),
// 		next:  next,
// 	}
// }

func (n Node) Add(value interface{}) (Node, bool) {
	// we'll maintain references to the "parent" and "next"
	// nodes. If we're able to add our node somewhere in the
	// tree, add it. Otherwise, see if we can make it the new
	// root node
	new, ok := n.attemptAdd(value)
	if ok {
		return new, true
	}
	return n.newRoot(value)
}

func (n Node) attemptAdd(value interface{}) (Node, bool) {
	switch n.value.(type) {
	case EmptyNode:
		// if our current node is an empty node, just replace it
		n.value = value
		return n, true
	case Operator:
		// if our node is holding an operator value, and
		// there's an empty slot, add our node to that. If
		// there isn't an empty slot, but one of the next nodes
		// is another operator, attempt to add it to that
		if len(n.next) < 2 {
			newNode := Make(value)
			n.next = append(n.next, newNode)
			return n, true
		} else {
			switch value.(type) {
			// if the node is an operator whose priority is higher than
			// the current node's priority, swap it out
			case Operator:
				curPriority := n.value.(Operator).priority
				valPriority := value.(Operator).priority

				if curPriority < valPriority {
					index := len(n.next) - 1
					new, ok := n.next[index].newRoot(value)
					if ok {
						n.next[index] = new
						return n, true
					}
				}
			}

			for i, next := range n.next {
				new, ok := next.attemptAdd(value)
				if ok {
					n.next[i] = new
					return n, true
				}
			}
		}
	}
	return n, false
}

// This will make a value the new root of the tree. We'll only
// do this if we can't insert a node somewhere further down
// the tree
func (n Node) newRoot(value interface{}) (Node, bool) {
	// if the the value is an Operator,
	// we'll make it the new root
	switch value.(type) {
	case Operator:
		newParent := Make(value)
		newParent.next = append(newParent.next, n)
		n = newParent
		return n, true
	}
	return n, false
}

func (n Node) Evaluate() int {
	switch n.value.(type) {
	case Literal:
		// fmt.Println("Literal:", n.value.(Literal).value)
		return n.value.(Literal).value
	case Operator:
		vals := []int{}
		for _, n := range n.next {
			val := n.Evaluate()
			vals = append(vals, val)
		}
		switch n.value.(Operator).operation {
		case "*":
			product := 1
			for _, v := range vals {
				product *= v
			}
			// fmt.Println("Product: Evaluating", vals)
			return product
		case "+":
			sum := 0
			for _, v := range vals {
				sum += v
			}
			return sum
		}
	case Node:
		return n.value.(Node).Evaluate()
	}
	return 0
}
