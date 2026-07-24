package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Sound struct {
	Pop      rl.Sound
	Bounce   rl.Sound
	Lose     rl.Sound
	GameOver rl.Sound
}

func (g *Game) PlaySound(sound rl.Sound) {
	if !g.muted {
		rl.PlaySound(sound)
	}
}
