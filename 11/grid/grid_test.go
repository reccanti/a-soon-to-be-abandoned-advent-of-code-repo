package grid

import (
	"fmt"
	"testing"
)

func TestNewAndString(t *testing.T) {
	// should create a grid that matches the given string
	g := New(5, 5, []interface{}{
		1, 0, 0, 0, 0,
		0, 10, 0, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1})
	predictedOutput := fmt.Sprintf("%s%s%s%s%s",
		"1  0  0  0  0 \n",
		"0  10 0  0  0 \n",
		"0  0  1  0  0 \n",
		"0  0  0  1  0 \n",
		"0  0  0  0  1 ",
	)
	if g.String() != predictedOutput {
		fmt.Println(g)
		fmt.Println("")
		fmt.Println(predictedOutput)
		t.Errorf("grid does not match predicted output")
	}
}

func TestGet(t *testing.T) {
	// rows and columns are zero-indexed. Attempt to
	// get the item at row 1 column 1 (10)
	g := New(5, 5, []interface{}{
		1, 0, 0, 0, 0,
		0, 10, 0, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1})
	cell, _ := g.Get(1, 1)
	if *cell != 10 {
		t.Errorf(fmt.Sprintf("Should have retrieved '10' from the given coordinates, but instead retrieved %d", cell))
	}

	// attempt to retrieve a valu that is out-of-bounds
	_, err := g.Get(5, 5)
	if err == nil {
		t.Errorf("should not be able to retrieve a value outside the  bounds of the grid")
	}
}

func TestSet(t *testing.T) {
	// rows and columns are zero-indexed. Attempt to
	// get the item at row 1 column 1 (10)
	g := New(5, 5, []interface{}{
		1, 0, 0, 0, 0,
		0, 10, 0, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1})
	g.Set(1, 1, 9)
	cell, _ := g.Get(1, 1)
	if *cell != 9 {
		t.Errorf(fmt.Sprintf("Should have retrieved '10' from the given coordinates, but instead retrieved %d", cell))
	}

	// attempt to retrieve a valu that is out-of-bounds
	err := g.Set(5, 5, 9)
	if err == nil {
		t.Errorf("should not be able to set a value outside the  bounds of the grid")
	}
}
