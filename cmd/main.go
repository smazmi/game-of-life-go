package main

import (
	"fmt"

	"github.com/smazmi/game-of-life-go/game"
)

func main() {
	world := game.NewWorld(5, 10)

	world.Set(1, 2, game.Alive)
	world.Set(3, 4, game.Alive)
	world.Set(3, 5, game.Alive)
	world.Set(3, 3, game.Alive)
	world.Set(3, 2, game.Alive)

	fmt.Println(world)

	world = world.Next()
	fmt.Println(world)
}
