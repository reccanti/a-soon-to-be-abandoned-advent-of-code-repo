package tree

import (
	"fmt"
)

// Literal - Holds "int" values

// type Literal struct {
// 	value int
// }

// func (l Literal) String() string {
// 	return fmt.Sprintf("%d", l.value)
// }

// func MakeLiteral(value int) Literal {
// 	return Literal{
// 		value,
// 	}
// }

// // Operator - combines "Literals" and other Operations

// type Operator struct {
// 	value string
// 	left  *interface{}
// 	right *interface{}
// }

// func (o Operator) String() string {
// 	left := ""
// 	if o.left != nil {
// 		left = fmt.Sprintf("%v", *o.left)
// 	}
// 	right := ""
// 	if o.right != nil {
// 		right = fmt.Sprintf("%v", *o.right)
// 	}
// 	return fmt.Sprintf("(%s %s %v)", left, o.value, right)
// }

// func MakeOperator(value string) Operator {
// 	return Operator{
// 		value: value,
// 		left:  nil,
// 		right: nil,
// 	}
// }

// func (o *Operator) Add(node interface{}) bool {
// 	if o.left == nil {
// 		o.left = &node
// 		return true
// 	} else if o.right == nil {
// 		o.right = &node
// 		return true
// 	} else {
// 		return false
// 	}
// }

// // Tree - The structure which holds everything together!

// type Tree struct {
// 	root *interface{}
// }

// func (t Tree) String() string {
// 	return fmt.Sprintf("%v", *t.root)
// }

// func MakeTree(root interface{}) Tree {
// 	return Tree{
// 		root: &root,
// 	}
// }

// func (t Tree) HasNext() bool {
// 	next := []*interface{}{t.root}
// 	for len(next) > 0 {
// 		curptr := next[0]
// 		cur := *curptr
// 		next = next[1:]

// 		switch cur.(type) {
// 		case Literal:
// 			// do nothing
// 		case Operator:
// 			// if the left or right operator exists, add it to our "next" array
// 			if cur.(Operator).left != nil {
// 				next = append(next, cur.(Operator).left)
// 			}
// 			if cur.(Operator).right != nil {
// 				next = append(next, cur.(Operator).right)
// 			}
// 			// if left or right is nil, we've found an empty slot!
// 			if cur.(Operator).left == nil {
// 				return true
// 			} else if cur.(Operator).right == nil {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func (t *Tree) Add(node interface{}) bool {
// 	// attempt to get the next available node. If we
// 	// cannot do this, make the current node the new root
// 	next := []*interface{}{t.root}
// 	for len(next) > 0 {
// 		curptr := next[0]
// 		cur := *curptr
// 		next = next[1:]

// 		switch cur.(type) {
// 		case Literal:
// 			// do nothing
// 		case Operator:
// 			op := cur.(Operator)
// 			// if the left or right operator exists, add it to our "next" array
// 			if cur.(Operator).left != nil {
// 				next = append(next, curptr)
// 			}
// 			if cur.(Operator).right != nil {
// 				next = append(next, curptr)
// 			}
// 			// if left or right is nil, we've found an empty slot!
// 			op.Add(node)
// 			var newNode interface{}
// 			newNode = op
// 			fmt.Println(newNode)
// 			curptr = &newNode
// 			return true
// 		}
// 	}
// 	switch node.(type) {
// 	case Literal:
// 		// do nothing because there are no open slots
// 	case Operator:
// 		// if the next node is an operator, make it the new root
// 		// of our tree
// 		op := node.(Operator)
// 		curRoot := *t.root
// 		op.Add(curRoot)

// 		var newRoot interface{}
// 		newRoot = op
// 		t.root = &newRoot

// 		return true
// 	}
// 	return false
// }

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
}

func (o Operator) String() string {
	return fmt.Sprintf("%s", o.operation)
}

func MakeOperator(operation string) Operator {
	return Operator{operation}
}

// Tree types

type Node struct {
	value interface{}
	next  []Node
}

func (n Node) String() string {
	// str := fmt.Sprintf("%v", n.value)
	// for _, val := range n.next {
	// 	str += fmt.Sprintf("\n-> %v", val)
	// }
	// return str
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

func Make(value interface{}) Node {
	return Node{
		value: value,
		next:  []Node{},
	}
}

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
	// switch parent.value.(type) {
	// // if the parent is a Literal and the value is an Operator,
	// // we'll make it the new root
	// case Literal:
	switch value.(type) {
	case Operator:
		newParent := Make(value)
		newParent.next = append(newParent.next, n)
		n = newParent
		return n, true
	}
	// }
	return n, false
}
