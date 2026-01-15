package game

import "math/rand/v2"

func SeedRandom(w *World, probability float64) {
	for y := range w.height {
		for x := range w.width {
			if rand.Float64() < probability {
				w.Set(x, y, Alive)
			}
		}
	}
}
