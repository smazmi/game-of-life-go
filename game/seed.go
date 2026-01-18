package game

import (
	"fmt"
	"math/rand/v2"

	"github.com/tlinden/golsky/rle"
)

func SeedRandom(w *World, probability float64) {
	for y := range w.height {
		for x := range w.width {
			if rand.Float64() < probability {
				w.Set(x, y, Alive)
			}
		}
	}
}

func LoadRLE(path string) (*World, error) {
	data, err := rle.GetRLE(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load rle file: %w", err)
	}

	world := NewWorld(data.Height, data.Width)

	for y, row := range data.Pattern {
		for x, cell := range row {
			if cell == 1 {
				world.Set(x, y, Alive)
			}
		}
	}

	return world, nil
}

func LoadRLECentered(path string, worldWidth, worldHeight int) (*World, error) {
	data, err := rle.GetRLE(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load rle file: %w", err)
	}

	patternHeight := data.Height
	patternWidth := data.Width

	offsetX := (worldWidth - patternWidth) / 2
	offsetY := (worldHeight - patternHeight) / 2
	// clamping
	if offsetX < 0 {
		offsetX = 0
	}
	if offsetY < 0 {
		offsetY = 0
	}

	world := NewWorld(worldHeight, worldWidth)

	for y, row := range data.Pattern {
		for x, cell := range row {
			if cell == 1 {
				world.Set(x+offsetX, y+offsetY, Alive)
			}
		}
	}

	return world, nil
}
