package main

import (
	"fmt"
	"time"

	"github.com/smazmi/game-of-life-go/game"
	"github.com/smazmi/game-of-life-go/utils"
)

func main() {
	worldWidth, worldHeight := utils.Size()

	world, err := game.LoadRLECentered(
		"/home/arno/projects/game-of-life-go/examples/copperhead.rle",
		worldWidth,
		worldHeight,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		utils.Clear()
		fmt.Println(world)
		world = world.Next()
		time.Sleep(200 * time.Millisecond)
	}
}
