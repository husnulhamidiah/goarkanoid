package internal

import (
	"fmt"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
)

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawRectangle(0, 0, c.WindowWidth, 6*c.WindowHeight/10, c.MiffyBlue)
	rl.DrawRectangle(0, 6*c.WindowHeight/10, c.WindowWidth, 5, c.MiffyBlack)
	rl.DrawRectangle(0, (6*c.WindowHeight/10)+5, c.WindowWidth, c.WindowHeight-(6*c.WindowHeight/10)-5, c.MiffyGreen)

	for i := 0; i < g.Lives; i++ {
		rl.DrawCircle(int32(g.Hearts[i].X), int32(g.Hearts[i].Y), float32(c.LiveCircleRadius), g.Hearts[i].color)
	}

	if g.State == c.GameOver {
		rl.DrawText(strings.ToUpper(c.GameOverText), int32(g.gameOverPos.X), int32(g.gameOverPos.Y), c.HeaderFontSize, rl.RayWhite)
	}
	rl.DrawText(fmt.Sprintf("%04d", g.Score), int32(g.scorePos.X), int32(g.scorePos.Y), c.HeaderFontSize, rl.RayWhite)

	for i := 0; i < c.BrickCount; i++ {
		if !g.Bricks[i].active {
			continue
		}
		rl.DrawRectangleRounded(g.Bricks[i].Rectangle, .3, 1, c.MiffyYellow)
	}

	rl.DrawCircleV(g.Ball.Vector2, c.BallRadius, rl.RayWhite)
	rl.DrawRectangleRounded(g.Paddle.Rectangle, .5, 1, c.MiffyOrange)

	rl.EndDrawing()
}
