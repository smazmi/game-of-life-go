package main

import (
	"fmt"
	"time"

	"github.com/smazmi/game-of-life-go/game"
	"github.com/smazmi/game-of-life-go/utils"
)

func main() {
	x, y := utils.Size()
	world := game.NewWorld(y, x)

	game.SeedRandom(world, 0.5)

	for {
		utils.Clear()
		fmt.Println(world)
		world = world.Next()
		time.Sleep(200 * time.Millisecond)
	}
}
