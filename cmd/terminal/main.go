package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/smazmi/game-of-life-go/game"
	"github.com/smazmi/game-of-life-go/utils"
)

func main() {
	worldWidth, worldHeight := utils.Size()

	examplePath := filepath.Join("examples", "glidergun.rle")

	world, err := game.LoadRLECentered(
		examplePath,
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
