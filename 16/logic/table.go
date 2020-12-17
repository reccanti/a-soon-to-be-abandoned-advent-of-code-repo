/**
 * A generic structure for managing logic. Because
 * I hate myself
 */
package logic

import (
	"fmt"
	"strings"
)

type Table struct {
	Rows    []interface{}
	Columns []interface{}
	invalid map[int]bool
	solved  map[int]bool
}

func NewTable(rows []interface{}, columns []interface{}) Table {
	return Table{
		Rows:    rows,
		Columns: columns,
		invalid: map[int]bool{},
		solved:  map[int]bool{},
	}
}

/**
 * Generic Getters
 */

func (t Table) Get(r int, c int) (interface{}, interface{}) {
	rowElem := t.Rows[r]
	colElem := t.Columns[c]

	return rowElem, colElem
}

func (t Table) GetRow(r int) []([]interface{}) {
	row := []([]interface{}){}
	for c := 0; c < len(t.Columns); c++ {
		rowElem, colElem := t.Get(r, c)
		row = append(row, []interface{}{rowElem, colElem})
	}
	return row
}

func (t Table) GetColumn(c int) []([]interface{}) {
	col := []([]interface{}){}
	for r := 0; r < len(t.Rows); r++ {
		rowElem, colElem := t.Get(r, c)
		col = append(col, []interface{}{rowElem, colElem})
	}
	return col
}

/**
 * "Invalid" functions
 */

func (t Table) MarkInvalid(r int, c int) {
	index := t.makeIndex(r, c)
	t.invalid[index] = true
}

func (t Table) IsValid(r int, c int) bool {
	index := t.makeIndex(r, c)
	i, hasEntry := t.invalid[index]
	if !hasEntry {
		return true
	}
	return !i
}

func (t Table) GetUnsolvedRow(r int) []([]interface{}) {
	row := []([]interface{}){}
	for c := 0; c < len(t.Columns); c++ {
		if t.IsValid(r, c) && !t.IsSolved(r, c) {
			rowElem, colElem := t.Get(r, c)
			row = append(row, []interface{}{rowElem, colElem})
		}
	}
	return row
}

func (t Table) GetUnsolvedColumn(c int) []([]interface{}) {
	col := []([]interface{}){}
	for r := 0; r < len(t.Rows); r++ {
		if t.IsValid(r, c) && !t.IsSolved(r, c) {
			fmt.Println(r, c, t.solved)
			rowElem, colElem := t.Get(r, c)
			col = append(col, []interface{}{rowElem, colElem})
		}
	}
	return col
}

/**
 * "Solved"
 */

func (t Table) MarkSolved(r int, c int) {
	for i := 0; i < len(t.Rows); i++ {
		t.MarkInvalid(r, i)
	}
	for i := 0; i < len(t.Columns); i++ {
		t.MarkInvalid(i, c)
	}
	index := t.makeIndex(r, c)
	t.invalid[index] = false
	t.solved[index] = true
}

func (t Table) IsSolved(r int, c int) bool {
	index := t.makeIndex(r, c)
	i, hasEntry := t.solved[index]
	if !hasEntry {
		return false
	}
	return i
}

/**
 * utility function, so that we'll always have consistent
 * indexes
 */

func (t Table) makeIndex(r int, c int) int {
	return len(t.Columns)*r + c
}

func (t Table) String() string {
	tableString := ""
	tableString += strings.Repeat("-", len(t.Columns)*4+1) + "\n"
	for r := 0; r < len(t.Rows); r++ {
		for c := 0; c < len(t.Columns); c++ {
			isValid := t.IsValid(r, c)
			isSolved := t.IsSolved(r, c)
			char := " "
			if !isValid {
				char = "X"
			} else if isSolved {
				char = "O"
			}
			tableString += "| " + char + " "
		}
		tableString += "|\n" + strings.Repeat("-", len(t.Columns)*4+1) + "\n"
	}
	return tableString
}
