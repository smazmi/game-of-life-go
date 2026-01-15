package game

func (w *World) String() string {
	var grid []rune
	for y := range w.height {
		for x := range w.width {
			switch w.Get(x, y) {
			case Alive:
				grid = append(grid, '@')
			case Dead:
				grid = append(grid, '.')
			}
		}
		grid = append(grid, '\n')
	}
	return string(grid)
}
