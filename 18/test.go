package main

import (
	"fmt"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/18/tree"
)

func main() {
	// l1 := tree.MakeLiteral(1)
	// l2 := tree.MakeLiteral(2)
	// // l3 := tree.MakeLiteral(3)
	// o1 := tree.MakeOperator("+")
	// o2 := tree.MakeOperator("*")

	// // o1.Add(l1)
	// // o1.Add(l2)
	// // o2.Add(l3)
	// // o2.Add(o1)

	// // nodes := []interface{}{l1, o, l2}

	// tree := tree.MakeTree(l1)
	// success1 := tree.Add(o1)
	// success2 := tree.Add(l2)
	// tree.Add(o2)

	// fmt.Println(success1)
	// fmt.Println(success2)

	// fmt.Println(tree)

	// 1 * 2 + 3
	// 1 * (2 + 3)

	// l1 := tree.MakeLiteral(1)
	// o1 := tree.MakeTimes()
	// l2 := tree.MakeLiteral(2)
	// o2 := tree.MakePlus()
	// l3 := tree.MakeLiteral(3)

	// n := tree.Make(l1)
	// n, _ = n.Add(o1)
	// n, _ = n.Add(l2)
	// n, _ = n.Add(o2)
	// n, _ = n.Add(l3)

	// 2 * 3 + 4 * 5
	l1 := tree.MakeLiteral(2)
	o1 := tree.MakeTimes()
	l2 := tree.MakeLiteral(3)
	o2 := tree.MakePlus()
	l3 := tree.MakeLiteral(4)
	o3 := tree.MakeTimes()
	l4 := tree.MakeLiteral(5)

	n := tree.Make(l1)
	n, _ = n.Add(o1)
	n, _ = n.Add(l2)
	n, _ = n.Add(o2)
	n, _ = n.Add(l3)
	n, _ = n.Add(o3)
	n, _ = n.Add(l4)

	fmt.Println(n.Evaluate())
}
