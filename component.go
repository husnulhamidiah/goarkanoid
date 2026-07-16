package main

import (
	"fmt"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	LiveCircleRadius = 12.5
	HeaderFontSize   = 25
	GameOverText     = "Game Over"
	BrickCount       = 16
	BrickWidth       = 50
	BrickHeight      = 25
	BallRadius       = 10
	BallXVelocity    = 200
	BallYVelocity    = -200
	BarWidth         = 100
	BarHeight        = 10
	BarVelocity      = 400
)

var (
	MiffyBlue   = rl.Color{G: 85, B: 153, A: 255}
	MiffyBlack  = rl.Color{R: 32, G: 34, B: 33, A: 255}
	MiffyGreen  = rl.Color{G: 113, B: 42, A: 255}
	MiffyOrange = rl.Color{R: 242, G: 101, B: 34, A: 255}
	MiffyYellow = rl.Color{R: 255, G: 200, B: 11, A: 255}
)

type brick struct {
	rl.Rectangle
	active bool
}

type circle struct {
	rl.Vector2
	radius float32
}

func render() {

	lifeCount := 3
	score := 0
	gameState := "initial"
	brickCounter := 0
	start := false
	revealTimer := 0.0

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(300, 600, "Brick Breaker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var lives []*circle
	for i := 0; i < lifeCount; i++ {
		lifeCirclePos := rl.Vector2{
			X: root.getRelativeComputedLeft(header.children[0].children[i]) + LiveCircleRadius,
			Y: root.getRelativeComputedTop(header.children[0].children[i]) + LiveCircleRadius,
		}
		lives = append(lives, &circle{lifeCirclePos, LiveCircleRadius})
	}

	paddedScore := fmt.Sprintf("%04d", score)
	scoreTextWidth := rl.MeasureText(paddedScore, HeaderFontSize)
	scoreTextRightEdge := header.children[2].node.GetComputedLeft() + header.children[2].node.GetComputedWidth()

	gameOverTextPos := rl.Vector2{
		X: header.children[0].node.GetComputedLeft(),
		Y: header.children[0].node.GetComputedTop(),
	}

	scoreTextPos := rl.Vector2{
		X: scoreTextRightEdge - float32(scoreTextWidth),
		Y: header.children[1].node.GetComputedTop(),
	}

	var bricks []*brick
	for i := 0; i < BrickCount; i++ {
		rect := rl.Rectangle{
			X:      content.children[i].node.GetComputedLeft(),
			Y:      root.getRelativeComputedTop(content.children[i]),
			Width:  float32(BrickWidth),
			Height: float32(BrickHeight),
		}
		bricks = append(bricks, &brick{rect, false})
	}

	ballPos := rl.Vector2{
		X: footer.children[0].node.GetComputedLeft() + BallRadius,
		Y: root.getRelativeComputedTop(footer.children[0]) + BallRadius,
	}
	ballVel := rl.Vector2{X: float32(BallXVelocity), Y: float32(BallYVelocity)}

	bar := rl.Rectangle{
		X:      footer.children[1].node.GetComputedLeft(),
		Y:      root.getRelativeComputedTop(footer.children[1]),
		Width:  float32(BarWidth),
		Height: float32(BarHeight),
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		if rl.IsKeyPressed(rl.KeySpace) {
			if gameState == "initial" || gameState == "playing" {
				start = true
				gameState = "playing"
			}

			if gameState == "gameover" {
				lifeCount = 3
				score = 0
				brickCounter = 0
				start = false
				gameState = "initial"
				for i := 0; i < BrickCount; i++ {
					bricks[i].active = false
				}
			}
		}

		if rl.IsKeyDown(rl.KeyRight) && gameState != "gameover" {
			bar.X += BarVelocity * dt
			if !start {
				ballPos.X += BarVelocity * dt
			}
		}

		if rl.IsKeyDown(rl.KeyLeft) && gameState != "gameover" {
			bar.X -= BarVelocity * dt
			if !start {
				ballPos.X -= BarVelocity * dt
			}
		}

		if start {
			ballPos = rl.Vector2Add(ballPos, rl.Vector2Scale(ballVel, dt))
		}

		if ballPos.X >= (float32(rl.GetScreenWidth())-BallRadius) || ballPos.X <= BallRadius {
			ballVel.X *= -1
		}

		if ballPos.Y <= BallRadius {
			ballVel.Y *= -1
		}

		if ballPos.Y >= (float32(rl.GetScreenHeight())-BallRadius) && start {
			ballPos = rl.Vector2{
				X: footer.children[0].node.GetComputedLeft() + BallRadius,
				Y: root.getRelativeComputedTop(footer.children[0]) + BallRadius,
			}
			bar.X = footer.children[1].node.GetComputedLeft()
			bar.Y = root.getRelativeComputedTop(footer.children[1])
			ballVel = rl.Vector2{X: float32(BallXVelocity), Y: float32(BallYVelocity)}

			start = false
			lifeCount -= 1
		}

		if lifeCount <= 0 && gameState == "playing" {
			score = 0
			gameState = "gameover"
		}

		if brickCounter == 0 && gameState == "playing" {
			gameState = "spawning"
		}

		for i := 0; i < BrickCount; i++ {
			if bricks[i].active {
				if rl.CheckCollisionCircleRec(ballPos, BallRadius, bricks[i].Rectangle) {
					ballVel.Y *= -1
					bricks[i].active = false
					score += 10
					brickCounter -= 1
					break
				}
			}
		}

		if rl.CheckCollisionCircleRec(ballPos, BallRadius, bar) && start {
			ballVel.Y *= -1
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(0, 0, 300, 350, MiffyBlue)
		rl.DrawRectangle(0, 350, 300, 5, MiffyBlack)
		rl.DrawRectangle(0, 355, 300, 800-350-5, MiffyGreen)

		for i := 0; i < lifeCount; i++ {
			rl.DrawCircle(int32(lives[i].X), int32(lives[i].Y), float32(LiveCircleRadius), MiffyOrange)
		}

		if gameState == "gameover" {
			rl.DrawText(strings.ToUpper(GameOverText), int32(gameOverTextPos.X), int32(gameOverTextPos.Y), HeaderFontSize, rl.RayWhite)
		}
		rl.DrawText(fmt.Sprintf("%04d", score), int32(scoreTextPos.X), int32(scoreTextPos.Y), HeaderFontSize, rl.RayWhite)

		if brickCounter < BrickCount && gameState == "initial" {
			revealTimer += float64(dt)
			if revealTimer >= 0.1 && !bricks[brickCounter].active {
				bricks[brickCounter].active = true
				brickCounter += 1
				revealTimer = 0.0
			}
		}

		threshold := int(bricks[len(bricks)-1].Y) + int(bricks[len(bricks)-1].Height) + (BallRadius * 2)
		if brickCounter < BrickCount && gameState == "spawning" && int(ballPos.Y) >= threshold {
			revealTimer += float64(dt)
			if revealTimer >= 0.1 && !bricks[brickCounter].active {
				bricks[brickCounter].active = true
				brickCounter += 1
				revealTimer = 0.0
			}
			if brickCounter >= BrickCount {
				gameState = "playing"
			}
		}

		for i := 0; i < BrickCount; i++ {
			if !bricks[i].active {
				continue
			}
			rl.DrawRectangleRounded(bricks[i].Rectangle, .3, 1, MiffyYellow)
		}

		rl.DrawCircleV(ballPos, BallRadius, rl.RayWhite)
		rl.DrawRectangleRounded(bar, .5, 1, MiffyOrange)

		rl.EndDrawing()
	}
}
