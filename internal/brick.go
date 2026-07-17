package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	rl.Rectangle
	active bool
	color  rl.Color
}
