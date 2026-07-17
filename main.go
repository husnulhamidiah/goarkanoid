package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/husnulhamidiah/goarkanoid/internal"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(300, 600, "Brick Breaker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := internal.NewGame()
	game.Init()
	defer game.Free()

	game.Run()
}
