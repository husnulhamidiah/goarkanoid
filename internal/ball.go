package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	rl.Vector2
	radius   float32
	velocity rl.Vector2
	color    rl.Color
}
