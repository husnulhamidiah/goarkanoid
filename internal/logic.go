package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
)

func (g *Game) Update() {

	dt := rl.GetFrameTime()

	if g.State == c.Playing || g.State == c.Spawning {
		g.Ball.Vector2 = rl.Vector2Add(g.Ball.Vector2, rl.Vector2Scale(g.Ball.velocity, dt))
	}

	if g.brickCounter < c.BrickCount && g.State == c.Initial && g.Score == 0 {
		g.revealTimer += dt
		if g.revealTimer >= 0.1 && !g.Bricks[g.brickCounter].active {
			g.Bricks[g.brickCounter].active = true
			g.brickCounter += 1
			g.revealTimer = 0.0
		}
	}

	if g.Ball.Vector2.X >= (float32(rl.GetScreenWidth())-c.BallRadius) || g.Ball.Vector2.X <= c.BallRadius {
		g.Ball.velocity.X *= -1
	}

	if g.Ball.Vector2.Y <= c.BallRadius {
		g.Ball.velocity.Y *= -1
	}

	if g.Ball.Vector2.Y >= (float32(rl.GetScreenHeight())-c.BallRadius) && g.State == c.Playing {
		g.Ball.Vector2 = g.ballOrigPos

		g.Paddle.Rectangle.X = g.paddleOrigPos.X
		g.Paddle.Rectangle.Y = g.paddleOrigPos.Y

		g.Ball.velocity = rl.Vector2{X: float32(c.BallXVelocity), Y: float32(c.BallYVelocity)}

		g.Lives -= 1

		if g.Lives <= 0 {
			g.State = c.GameOver
		} else {
			g.State = c.Initial
		}
	}

	// there's a bug where ball hit the side of the paddle, the ball will stick inside the paddle
	paddleTopEdgeStart := rl.Vector2{X: g.Paddle.Rectangle.X, Y: g.Paddle.Rectangle.Y}
	paddleTopEdgeEnd := rl.Vector2{X: g.Paddle.Rectangle.X + g.Paddle.Rectangle.Width, Y: g.Paddle.Rectangle.Y}
	if rl.CheckCollisionCircleLine(g.Ball.Vector2, c.BallRadius, paddleTopEdgeStart, paddleTopEdgeEnd) && g.State == c.Playing {
		g.Ball.velocity.Y *= -1
	}

	for i := 0; i < c.BrickCount; i++ {
		if g.Bricks[i].active {
			if rl.CheckCollisionCircleRec(g.Ball.Vector2, c.BallRadius, g.Bricks[i].Rectangle) {
				g.Ball.velocity.Y *= -1
				g.Bricks[i].active = false
				g.Score += 10
				g.brickCounter -= 1
				break
			}
		}
	}

	if g.brickCounter == 0 && g.State == c.Playing {
		g.State = c.Spawning
	}

	threshold := int(g.Bricks[c.BrickCount-1].Y) + int(g.Bricks[c.BrickCount-1].Height) + (c.BallRadius * 2)
	if g.brickCounter < c.BrickCount && g.State == c.Spawning && int(g.Ball.Vector2.Y) >= threshold {
		g.revealTimer += dt
		if g.revealTimer >= 0.1 && !g.Bricks[g.brickCounter].active {
			g.Bricks[g.brickCounter].active = true
			g.brickCounter += 1
			g.revealTimer = 0.0
		}
		if g.brickCounter >= c.BrickCount {
			g.State = c.Playing
		}
	}
}
