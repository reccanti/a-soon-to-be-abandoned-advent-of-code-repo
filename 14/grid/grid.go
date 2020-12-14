package grid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Grid struct {
	Rows    int
	Columns int
	cells   []interface{}
}

func New(rows int, columns int, cells []interface{}) Grid {
	return Grid{
		Rows:    rows,
		Columns: columns,
		cells:   cells,
	}
}

// get and set methods

func (g Grid) IsInBounds(row int, column int) bool {
	if row > g.Rows-1 || row < 0 {
		return false
	}
	if column > g.Columns-1 || column < 0 {
		return false
	}
	return true
}

func (g Grid) Get(row int, column int) (*interface{}, error) {
	index := row*g.Columns + column
	if !g.IsInBounds(row, column) {
		return nil, errors.New(fmt.Sprintf("The give row %d and column %d are out-of-bounds on the given grid", row, column))
	}
	return &g.cells[index], nil
}

func (g Grid) Set(row int, column int, val interface{}) error {
	index := row*g.Columns + column
	if !g.IsInBounds(row, column) {
		return errors.New(fmt.Sprintf("The give row %d and column %d are out-of-bounds on the given grid", row, column))
	}
	g.cells[index] = val
	return nil
}

/**
 * String methods
 */

func toString(i interface{}) string {
	switch i.(type) {
	case nil:
		return "nil"
	case int:
		return strconv.Itoa(i.(int))
	case string:
		return i.(string)
	case bool:
		if i.(bool) {
			return "T"
		} else {
			return "F"
		}
	default:
		return "?"
	}
}

func padStr(str string, length int) string {
	diff := length - len(str)
	padding := strings.Repeat(" ", diff)
	return fmt.Sprintf("%s%s", str, padding)
}

func (g Grid) String() string {
	str := ""

	// first create "string" representations of everything in the list
	// also determine the longest string in the list
	strCells := []string{}
	longest := 0
	for _, c := range g.cells {
		sc := toString(c)
		if len(sc) > longest {
			longest = len(sc)
		}
		strCells = append(strCells, sc)
	}

	// loop over our stringified cells to
	for i, cell := range strCells {
		str = str + padStr(cell, longest)
		if i == len(strCells)-1 {
		} else if (i+1)%g.Columns == 0 {
			str = str + "\n"
		} else {
			str = str + " "
		}
	}
	return str
}
