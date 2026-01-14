package main

import "fmt"

type CellState string

const (
	Alive CellState = "@"
	Dead  CellState = "."
)

func (c CellState) Valid() bool {
	switch c {
	case "@", ".":
		return true
	default:
		return false
	}
}

type Grid struct {
	height, width int
	cells         [][]CellState
}

func (g *Grid) Print() {
	for y := range g.height {
		for x := range g.width {
			fmt.Print(g.cells[y][x])
		}
		fmt.Println()
	}
}

func gridConstructor(height int, width int) *Grid {
	cells := make([][]CellState, height)
	for i := range cells {
		cells[i] = make([]CellState, width)
		for j := range cells[i] {
			cells[i][j] = Dead
		}
	}
	return &Grid{
		height: height,
		width:  width,
		cells:  cells,
	}
}

func main() {
	g := gridConstructor(5, 10)
	g.Print()
}
