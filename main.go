package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
	"github.com/husnulhamidiah/goarkanoid/internal"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(c.WindowWidth, c.WindowHeight, "Brick Breaker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := internal.NewGame()
	game.Init()
	defer game.Free()

	game.Run()
}
