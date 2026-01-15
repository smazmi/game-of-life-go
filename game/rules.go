package game

// ## Rules
//
// ### For Live Cells
//
// - **Underpopulation**
//   - If a live cell has **fewer than 2 live neighbors**, it becomes **dead**.
//
// - **Survival**
//   - If a live cell has **2 or 3 live neighbors**, it **stays alive**.
//
// - **Overpopulation**
//   - If a live cell has **more than 3 live neighbors**, it becomes **dead**.
//
// ### For Dead Cells
//
// - **Reproduction**
//   - If a dead cell has **exactly 3 live neighbors**, it becomes **alive**.
//

func (w *World) Next() *World {
	next := NewWorld(w.height, w.width)
	for y := range next.height {
		for x := range next.width {
			current := w.Get(x, y)
			liveNeighbors := w.LiveNeighbors(x, y)

			var nextState CellState

			switch {
			case current == Alive && (liveNeighbors < 2 || liveNeighbors > 3):
				nextState = Dead
			case current == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
				nextState = Alive
			case current == Dead && liveNeighbors == 3:
				nextState = Alive
			default:
				nextState = Dead
			}

			next.Set(x, y, nextState)
		}
	}
	return next
}
