package game

type offset struct {
	dx int // column
	dy int // row
}

var neighborOffsets = []offset{
	{-1, 0},  // left
	{1, 0},   // right
	{0, -1},  // top
	{0, 1},   // bottom
	{-1, -1}, // top-left
	{1, -1},  // top-right
	{-1, 1},  // bottom-left
	{1, 1},   // bottom-right
}

func (w *World) LiveNeighbors(x, y int) int {
	count := 0
	for _, offset := range neighborOffsets {
		nx := x + offset.dx
		ny := y + offset.dy
		if w.Get(nx, ny) == Alive {
			count++
		}
	}
	return count
}
