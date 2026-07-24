package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
)

func (g *Game) Update() {

	dt := rl.GetFrameTime()

	if g.State == c.Playing {
		g.Ball.Vector2 = rl.Vector2Add(g.Ball.Vector2, rl.Vector2Scale(g.Ball.velocity, dt))
	}

	if g.brickCounter == c.BrickCount {
		g.spawning = false
	}

	if g.brickCounter == 0 && g.State == c.Playing {
		g.spawning = true
	}

	if g.brickCounter < c.BrickCount {
		if g.State == c.Initial && g.Score == 0 {
			g.revealTimer += dt
			if g.revealTimer >= 0.1 && !g.Bricks[g.brickCounter].active {
				g.Bricks[g.brickCounter].active = true
				g.brickCounter += 1
				g.revealTimer = 0.0
			}
			return
		}

		threshold := int(g.Bricks[c.BrickCount-1].Y) + int(g.Bricks[c.BrickCount-1].Height) + (c.BallRadius * 2)
		if g.spawning && g.State == c.Playing && g.Score != 0 && int(g.Ball.Y) >= threshold {
			g.revealTimer += dt
			if g.revealTimer >= 0.1 && !g.Bricks[g.brickCounter].active {
				g.Bricks[g.brickCounter].active = true
				g.brickCounter += 1
				g.revealTimer = 0.0
			}
		}
	}

	if g.Ball.X >= (float32(rl.GetScreenWidth()) - c.BallRadius) {
		g.Ball.X = float32(rl.GetScreenWidth() - c.BallRadius)
		g.Ball.velocity.X *= -1
	}

	if g.Ball.X <= c.BallRadius {
		g.Ball.X = float32(c.BallRadius)
		g.Ball.velocity.X *= -1
	}

	if g.Ball.Y <= c.BallRadius {
		g.Ball.Y = float32(c.BallRadius)
		g.Ball.velocity.Y *= -1
	}

	if g.Ball.Y >= (float32(rl.GetScreenHeight())-c.BallRadius) && g.State == c.Playing {
		g.Ball.Vector2 = g.ballOrigPos

		g.Paddle.X = g.paddleOrigPos.X
		g.Paddle.Y = g.paddleOrigPos.Y

		g.Ball.velocity = rl.Vector2{X: float32(c.BallXVelocity), Y: float32(c.BallYVelocity)}

		g.Lives -= 1

		if g.Lives <= 0 {
			g.State = c.GameOver
			g.PlaySound(g.Sound.GameOver)
		} else {
			g.State = c.Initial
			g.PlaySound(g.Sound.Lose)
		}
	}

	paddleTopEdgeStart := rl.Vector2{X: g.Paddle.X, Y: g.Paddle.Y}
	paddleTopEdgeEnd := rl.Vector2{X: g.Paddle.X + g.Paddle.Width, Y: g.Paddle.Y}
	if rl.CheckCollisionCircleLine(g.Ball.Vector2, c.BallRadius, paddleTopEdgeStart, paddleTopEdgeEnd) && g.State == c.Playing {
		g.Ball.Y = g.Paddle.Y - c.BallRadius
		g.Ball.velocity.Y *= -1

		speed := rl.Vector2Length(g.Ball.velocity)
		g.Ball.velocity.X += g.Paddle.velocity.X * 0.3
		dir := rl.Vector2Normalize(g.Ball.velocity)
		g.Ball.velocity = rl.Vector2Scale(dir, speed)
		g.PlaySound(g.Sound.Bounce)
	}

	for i := 0; i < c.BrickCount; i++ {
		if g.Bricks[i].active {
			if rl.CheckCollisionCircleRec(g.Ball.Vector2, c.BallRadius, g.Bricks[i].Rectangle) {
				g.Ball.velocity.Y *= -1
				g.Bricks[i].active = false
				g.Score += 10
				g.brickCounter -= 1

				g.PlaySound(g.Sound.Pop)
				break
			}
		}
	}
}
