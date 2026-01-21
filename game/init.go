package game

type CellState int

const (
	Alive CellState = iota
	Dead
)

type World struct {
	height, width int
	cells         [][]CellState
}

func NewWorld(height int, width int) *World {
	cells := make([][]CellState, height)
	for i := range cells {
		cells[i] = make([]CellState, width)
		for j := range cells[i] {
			cells[i][j] = Dead
		}
	}
	return &World{
		height: height,
		width:  width,
		cells:  cells,
	}
}

func (w *World) GetHeight() int {
	return w.height
}

func (w *World) GetWidth() int {
	return w.width
}

func (w *World) inBounds(x, y int) bool {
	return x >= 0 && x < w.width &&
		y >= 0 && y < w.height
}

func (w *World) Get(x, y int) CellState {
	if !w.inBounds(x, y) {
		return Dead
	}
	return w.cells[y][x]
}

func (w *World) Set(x, y int, state CellState) {
	if w.inBounds(x, y) {
		w.cells[y][x] = state
	}
}
