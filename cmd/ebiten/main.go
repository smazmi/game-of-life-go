package main

import (
	"image/color"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/smazmi/game-of-life-go/game"
)

const (
	cellSize       = 10
	worldWidth     = 80
	worldHeight    = 60
	updatesPerStep = 5
)

type Game struct {
	world *game.World
	tick  int
}

func (g *Game) Update() error {
	g.tick++
	if g.tick%updatesPerStep == 0 {
		g.world = g.world.Next()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for y := range g.world.GetHeight() {
		for x := range g.world.GetWidth() {
			switch g.world.Get(x, y) {
			case game.Alive:
				px := x * cellSize
				py := y * cellSize
				vector.FillRect(screen, float32(px), float32(py), float32(cellSize), float32(cellSize), color.White, true)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return worldWidth * cellSize, worldHeight * cellSize
}

func main() {
	examplePath := filepath.Join("examples", "glidergun.rle")

	w, err := game.LoadRLECentered(
		examplePath,
		worldWidth,
		worldHeight,
	)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(worldWidth*cellSize, worldHeight*cellSize)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(&Game{world: w}); err != nil {
		log.Fatal(err)
	}
}
