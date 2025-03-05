// puzzle and puzzle content structures
// puzzle creation


package puzzle

import (
	"fmt"
	"strconv"
)

// Puzzle data objects ======================================================================================
type CellValueList struct {
	v [9]bool
}

type Cell struct {
	OriginalFlag   bool // may be nil
	Value          *int // may be nil
	PossibleValues CellValueList
	RejectedValues CellValueList
}

func (c Cell) ToString() string {
	if c.Value == nil {
		return " "
	}
	return strconv.Itoa(*c.Value + 1)
}

type CellContainer struct {
	cells [9]*Cell
}

type Puzzle struct {
	rows   [9]CellContainer
	cols   [9]CellContainer
	blocks [9]CellContainer
}

func (p Puzzle) GetCell(index int) *Cell {
	coords := CoordinatesFromIndex(index)
	cellPtr := p.rows[coords.row].cells[coords.rowIndex()]
	return cellPtr
}

// assumes symbols are 1 to 9 (for now)
func (p Puzzle) GetSymbol(index int) string {
	intPtr := p.GetCell(index).Value
	if intPtr == nil {
		return " "
	}
	return strconv.Itoa(*intPtr + 1)
}

// / initialize a puzzle from an initial value struct
func CreatePuzzle(piv PuzzleInitialValues) Puzzle {
	var p Puzzle
	allFalse := CellValueList{v: [...]bool{false, false, false, false, false, false, false, false, false}}
	allTrue := CellValueList{v: [...]bool{true, true, true, true, true, true, true, true, true}}

	for index, val := range piv.V {
		// create cell
		cell := Cell{OriginalFlag: false, Value: nil, PossibleValues: allTrue, RejectedValues: allFalse}
		if val != nil {
			// fmt.Printf("Creating cell index: %2d with address: %d and value %d\n", index, val, *val)
			cell.OriginalFlag = true
			cell.Value = val
			cell.PossibleValues = allFalse
			cell.PossibleValues.v[*val] = true
			cell.RejectedValues = allTrue
			cell.RejectedValues.v[*val] = false
		}

		// place cell pointer in appropriate row, col and block
		coords := CoordinatesFromIndex(index)
		p.rows[coords.row].cells[coords.rowIndex()] = &cell
		p.cols[coords.col].cells[coords.colIndex()] = &cell
		p.blocks[coords.block].cells[coords.blockIndex()] = &cell

	}
	return p
}

type PuzzleCoordinates struct {
	row   int
	col   int
	block int
}

func (pc PuzzleCoordinates) rowIndex() int {
	return pc.col
}

func (pc PuzzleCoordinates) colIndex() int {
	return pc.row
}

func (pc PuzzleCoordinates) blockIndex() int {
	return pc.row%3*3 + pc.col%3
}

// converts an index value into puzzle coordinates
// assumes the puzzle is indexed from 0 to 80 by column then row
func CoordinatesFromIndex(index int) PuzzleCoordinates {
	if index < 0 || index > 80 {
		panic("Invalid index: " + strconv.Itoa(index) + ". Index must be between 0 and 80")
	}

	var pc PuzzleCoordinates
	pc.row = index / 9
	pc.col = index - (9 * pc.row)
	pc.block = pc.row/3 + pc.col/3

	return pc
}

// converts an index value into puzzle coordinates
// assumes the puzzle is indexed from 0 to 80 by column then row
func GetIndexFromRowCol(row int, col int) int {
	return row*9 + col
}

type PuzzleInitialValues struct {
	V [81]*int
}

func (piv PuzzleInitialValues) ToString() string {
	s := ""
	for row := 0; row < 9; row++ {
		s += fmt.Sprintf("%d: ", row)
		for col := 0; col < 9; col++ {
			index := GetIndexFromRowCol(row, col)
			if piv.V[index] == nil {
				s += fmt.Sprintf("- ")
			} else {
				s += fmt.Sprintf("%d ", *piv.V[index])
			}
		}
		s += "\n"
	}
	return s
}
