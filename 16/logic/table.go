/**
 * A generic structure for managing logic. Because
 * I hate myself
 */
package logic

import (
	"strings"
)

type Table struct {
	Rows    []interface{}
	Columns []interface{}
	invalid map[int]bool
	solved  map[int]bool
}

type Cell struct {
	Row         int
	RowValue    interface{}
	Column      int
	ColumnValue interface{}
}

func NewTable(rows []interface{}, columns []interface{}) Table {
	return Table{
		Rows:    rows,
		Columns: columns,
		invalid: map[int]bool{},
		solved:  map[int]bool{},
	}
}

func NewCell(row int, rowValue interface{}, column int, columnValue interface{}) Cell {
	return Cell{
		Row:         row,
		RowValue:    rowValue,
		Column:      column,
		ColumnValue: columnValue,
	}
}

/**
 * Generic Getters
 */

func (t Table) Get(r int, c int) Cell {
	rowElem := t.Rows[r]
	colElem := t.Columns[c]

	return NewCell(r, rowElem, c, colElem)
}

func (t Table) GetRow(r int) []Cell {
	row := []Cell{}
	for c := 0; c < len(t.Columns); c++ {
		cell := t.Get(r, c)
		row = append(row, cell)
	}
	return row
}

func (t Table) GetColumn(c int) []Cell {
	col := []Cell{}
	for r := 0; r < len(t.Rows); r++ {
		cell := t.Get(r, c)
		col = append(col, cell)
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

func (t Table) GetUnsolvedRow(r int) []Cell {
	row := []Cell{}
	for c := 0; c < len(t.Columns); c++ {
		if t.IsValid(r, c) && !t.IsSolved(r, c) {
			cell := t.Get(r, c)
			row = append(row, cell)
		}
	}
	return row
}

func (t Table) GetUnsolvedColumn(c int) []Cell {
	col := []Cell{}
	for r := 0; r < len(t.Rows); r++ {
		if t.IsValid(r, c) && !t.IsSolved(r, c) {
			cell := t.Get(r, c)
			col = append(col, cell)
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

func (t Table) GetSolution() ([]Cell, bool) {
	// the number of "solved" entries should be
	// equal to the number of rows or columns
	// in our table
	solvedEntries := []Cell{}
	for i := 0; i < len(t.Rows); i++ {
		for j := 0; j < len(t.Columns); j++ {
			if t.IsSolved(i, j) {
				cell := t.Get(i, j)
				solvedEntries = append(solvedEntries, cell)
			}
		}
	}
	if len(solvedEntries) == len(t.Rows) {
		return solvedEntries, true
	} else {
		return solvedEntries, false
	}
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
