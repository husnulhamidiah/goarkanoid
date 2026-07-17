package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
)

func (g *Game) Input() {
	dt := rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeySpace) {
		if g.State == c.Initial {
			g.State = c.Playing
		}

		if g.State == c.GameOver {
			g.Init()
			for i := 0; i < c.BrickCount; i++ {
				g.Bricks[i].active = false
			}
		}
	}

	if rl.IsKeyDown(rl.KeyRight) && g.State != c.GameOver {
		g.Paddle.X += c.BarVelocity * dt
		if g.State == c.Initial {
			g.Ball.X += c.BarVelocity * dt
		}
	}

	if rl.IsKeyDown(rl.KeyLeft) && g.State != c.GameOver {
		g.Paddle.X -= c.BarVelocity * dt
		if g.State == c.Initial {
			g.Ball.X -= c.BarVelocity * dt
		}
	}
}
