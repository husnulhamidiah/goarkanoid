package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	rl.Rectangle
	velocity rl.Vector2
}
