package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Brick struct {
	rl.Rectangle
	Active bool
}

var (
	windowWidth = 420

	gap         = 3
	brickWidth  = 66
	brickHeight = 37

	hudTopPadding   = 10
	hudLeftPadding  = 10
	hudCircleRadius = 20

	speed   = 300
	started = false

	lives = 3

	brickCount = 0
)

func main() {

	bricksColor := [4]rl.Color{rl.Blue, rl.Red, rl.Orange, rl.Green}
	bricks := [4][5]*Brick{}

	rl.InitWindow(420, 820, "Brick Breaker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	boardOffset := 0

	circlePos := rl.Vector2{X: 10, Y: 800}
	vel := rl.Vector2{X: float32(speed), Y: float32(-speed)}

	score := "0000"

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		circleX := hudLeftPadding + hudCircleRadius
		circleY := hudCircleRadius + hudTopPadding
		for i := 0; i < lives; i++ {
			rl.DrawCircle(int32(circleX+i*(2*hudCircleRadius+gap)), int32(circleY), float32(hudCircleRadius), rl.DarkGray)
		}

		fontSize := 50
		f := rl.GetFontDefault()
		sz := rl.MeasureTextEx(f, score, float32(fontSize), 1)

		textX := (420 - sz.X) - float32(hudLeftPadding)
		textY := (50-sz.Y)/2 + float32(hudTopPadding)

		rl.DrawTextEx(f, score, rl.Vector2{X: float32(textX), Y: float32(textY)}, 50, 1, rl.DarkGray)

		for rowidx, row := range bricks {
			for colidx, _ := range row {

				x := gap + colidx*(brickWidth+gap)
				y := gap + rowidx*(brickHeight+gap)

				brickLeftPadding := (windowWidth - (brickWidth * 5) - (gap * 5)) / 2
				x = brickLeftPadding + x

				brickTopPadding := 65
				y = brickTopPadding + y

				rect := rl.Rectangle{
					X:      float32(x),
					Y:      float32(y),
					Width:  float32(brickWidth),
					Height: float32(brickHeight),
				}

				if bricks[rowidx][colidx] == nil {
					bricks[rowidx][colidx] = &Brick{rect, true}
				}

				if bricks[rowidx][colidx].Active {
					if rl.CheckCollisionCircleRec(circlePos, float32(hudCircleRadius), bricks[rowidx][colidx].Rectangle) {
						bricks[rowidx][colidx].Active = false

						brickCount++

						vel.Y *= -1
						vel.X *= -1

						scoreInt, _ := strconv.Atoi(score)
						scoreInt += 100
						score = strconv.Itoa(scoreInt)
					}
					rl.DrawRectangleRounded(bricks[rowidx][colidx].Rectangle, .5, 1, bricksColor[rowidx])
				}
			}
		}

		if brickCount >= 15 && circlePos.Y > 600 {
			brickCount = 0
			for rowidx, row := range bricks {
				for colidx, _ := range row {
					bricks[rowidx][colidx].Active = true
				}
			}
		}

		boardLength := 66 * 2
		boardY := rl.GetScreenHeight() - 37
		boardXStart := rl.Vector2{X: 0 + float32(boardOffset), Y: float32(boardY)}
		boardXEnd := rl.Vector2{X: boardXStart.X + float32(boardLength), Y: float32(boardY)}

		rl.DrawLineEx(boardXStart, boardXEnd, float32(15), rl.DarkGray)

		if rl.CheckCollisionCircleLine(circlePos, float32(hudCircleRadius), boardXStart, boardXEnd) {
			vel.Y *= -1
		}

		dt := rl.GetFrameTime()
		//circlePos.X = circlePos.X + vel.X*dt
		//circlePos.Y = circlePos.Y + vel.Y*dt
		//rl.DrawCircleV(circlePos, float32(hudCircleRadius), rl.Black)

		if circlePos.X-float32(hudCircleRadius) < 0 {
			circlePos.X = float32(hudCircleRadius)
			vel.X *= -1
		}
		if circlePos.X+float32(hudCircleRadius) > 420 {
			circlePos.X = float32(420 - hudCircleRadius)
			vel.X *= -1.0
		}
		if circlePos.Y-float32(hudCircleRadius) < 0 {
			circlePos.Y = float32(hudCircleRadius)
			vel.Y *= -1.0
		}
		if circlePos.Y+float32(hudCircleRadius) > 820 {
			circlePos.Y = float32(820 - hudCircleRadius)
			//vel.Y *= -1

			// stick to bottom
			vel.Y = 0
			vel.X = 0

			started = false
			circlePos.X = float32(boardXStart.X + float32(boardLength/2))
			circlePos.Y = float32(boardY - 28)

			lives -= 1

			if lives == 0 {
				score = score + " Game Over"
			}
		}

		if !started {
			circlePos.X = float32(boardXStart.X + float32(boardLength/2))
			circlePos.Y = float32(boardY - 28)
		} else {
			circlePos.X = circlePos.X + vel.X*dt
			circlePos.Y = circlePos.Y + vel.Y*dt
		}
		rl.DrawCircleV(circlePos, float32(hudCircleRadius), rl.Black)

		if rl.IsKeyPressed(rl.KeySpace) {
			if score == "Game Over" {
				for rowidx, row := range bricks {
					for colidx, _ := range row {
						bricks[rowidx][colidx].Active = true
					}
				}
				lives = 3
				score = "0000"
			}

			started = true
			vel.Y = float32(speed)
			vel.X = float32(-speed)
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			if int(boardXEnd.X)+50 > windowWidth {
				boardOffset += int(float32(windowWidth) - boardXEnd.X)
			} else if boardXEnd.X > 420 {
				boardOffset += 0
			} else {
				boardOffset += 50
			}
		}
		if rl.IsKeyPressedRepeat(rl.KeyRight) {
			if int(boardXEnd.X)+75 > windowWidth {
				boardOffset += int(float32(windowWidth) - boardXEnd.X)
			} else if boardXEnd.X > 420 {
				boardOffset += 0
			} else {
				boardOffset += 75
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			if int(boardXStart.X)-50 <= 0 {
				boardOffset = 0
			} else if boardXStart.X < 0 {
				boardOffset -= 0
			} else {
				boardOffset -= 50
			}
		}
		if rl.IsKeyPressedRepeat(rl.KeyLeft) {
			if int(boardXStart.X)-75 <= 0 {
				boardOffset = 0
			} else if boardXStart.X < 0 {
				boardOffset -= 0
			} else {
				boardOffset -= 75
			}
		}

		rl.EndDrawing()
	}
}
