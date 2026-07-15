package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	goyoga "github.com/husnulhamidiah/goyoga/std"
)

type brick struct {
	rl.Rectangle
	active bool
}

type box struct {
	node     *goyoga.Node
	label    string
	children []*box
}

type circle struct {
	rl.Vector2
	radius float32
}

func newBox(config *goyoga.Config, label string) *box {
	return &box{node: goyoga.NewNodeWithConfig(config), label: label}
}

func fixedBox(config *goyoga.Config, label string, width, height float32) *box {
	b := newBox(config, label)
	b.node.SetWidth(width)
	b.node.SetHeight(height)
	return b
}

func (b *box) addChild(child *box) {
	b.node.InsertChild(child.node, uint(len(b.children)))
	b.children = append(b.children, child)
}

func (b *box) getRelativeComputedLeft(target *box) (left float32) {
	node := target.node
	for node != nil && node != b.node {
		left += node.GetComputedLeft()
		node = node.GetOwner()
	}
	return
}

func (b *box) getRelativeComputedTop(target *box) (top float32) {
	node := target.node
	for node != nil && node != b.node {
		top += node.GetComputedTop()
		node = node.GetOwner()
	}
	return
}

func main() {
	config := goyoga.NewConfig()
	defer config.Free()
	config.SetUseWebDefaults(true)

	root := newBox(config, "root")
	defer root.node.FreeRecursive()
	root.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	root.node.SetWidth(300)
	root.node.SetHeight(600)

	header := newBox(config, "header")
	header.node.SetPadding(goyoga.EdgeAll, 15)
	header.node.SetJustifyContent(goyoga.JustifySpaceBetween)
	header.node.SetAlignItems(goyoga.AlignCenter)

	header.addChild(newBox(config, "lives"))
	header.children[0].node.SetJustifyContent(goyoga.JustifySpaceBetween)
	header.children[0].node.SetAlignItems(goyoga.AlignCenter)
	header.children[0].node.SetGap(goyoga.GutterAll, 2)

	header.addChild(fixedBox(config, fmt.Sprintf("top-%d", 2), 50, 25))
	header.addChild(fixedBox(config, fmt.Sprintf("top-%d", 3), 50, 25))

	for i := 1; i <= 3; i++ {
		header.children[0].addChild(fixedBox(config, fmt.Sprintf("it-%d", i), 25, 25))
	}

	content := newBox(config, "content")
	content.node.SetFlexGrow(1)
	content.node.SetPadding(goyoga.EdgeAll, 20)
	content.node.SetGap(goyoga.GutterAll, 5)
	content.node.SetJustifyContent(goyoga.JustifyCenter)
	content.node.SetFlexWrap(goyoga.WrapWrap)
	content.node.SetAlignContent(goyoga.AlignFlexStart)
	for i := 1; i <= 16; i++ {
		content.addChild(fixedBox(config, fmt.Sprintf("item-%02d", i), 50, 25))
	}

	footer := newBox(config, "footer")
	footer.node.SetPadding(goyoga.EdgeAll, 15)
	footer.node.SetJustifyContent(goyoga.JustifyCenter)
	footer.node.SetAlignItems(goyoga.AlignCenter)
	footer.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	footer.addChild(fixedBox(config, "dot", 20, 20))
	footer.addChild(fixedBox(config, "bar", 100, 10))

	root.addChild(header)
	root.addChild(content)
	root.addChild(footer)
	root.node.CalculateLayout(300, 600, goyoga.DirectionLTR)

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(300, 600, "Brick Breaker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	lifeCount := 3
	var lifes []*circle
	for i := 0; i < lifeCount; i++ {
		x := float32(root.getRelativeComputedLeft(header.children[0].children[i]) + 12.5)
		y := float32(root.getRelativeComputedTop(header.children[0].children[i]) + 12.5)
		lifes = append(lifes, &circle{
			rl.Vector2{X: x, Y: y},
			12.5,
		})
	}

	var bricks []*brick
	for i := 0; i < 16; i++ {
		rect := rl.Rectangle{
			X:      float32(content.children[i].node.GetComputedLeft()),
			Y:      float32(root.getRelativeComputedTop(content.children[i])),
			Width:  float32(50),
			Height: float32(25),
		}
		bricks = append(bricks, &brick{rect, true})
	}
	brickCounter := 0

	ballPos := rl.Vector2{
		X: footer.children[0].node.GetComputedLeft() + 10,
		Y: root.getRelativeComputedTop(footer.children[0]) + 10,
	}
	ballVel := rl.Vector2{X: float32(200), Y: float32(-200)}

	bar := rl.Rectangle{
		X:      float32(footer.children[1].node.GetComputedLeft()),
		Y:      float32(root.getRelativeComputedTop(footer.children[1])),
		Width:  float32(100),
		Height: float32(10),
	}

	score := 0
	start := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(0, 0, 300, 350, rl.Color{0, 85, 153, 255})
		rl.DrawRectangle(0, 350, 300, 5, rl.Color{32, 34, 33, 255})
		rl.DrawRectangle(0, 355, 300, 800-350-5, rl.Color{0, 113, 42, 255})

		func() {
			for i := 0; i < lifeCount; i++ {
				rl.DrawCircle(int32(lifes[i].X), int32(lifes[i].Y), float32(12.5), rl.Color{242, 101, 34, 255})
			}
		}()

		func() {
			paddedScore := fmt.Sprintf("%04d", score)
			scoreTextWidth := rl.MeasureText(paddedScore, 25)
			scoreTextRightEdge := header.children[2].node.GetComputedLeft() + header.children[2].node.GetComputedWidth()

			textX := int32(scoreTextRightEdge) - scoreTextWidth
			textY := header.children[1].node.GetComputedTop()
			rl.DrawText(paddedScore, int32(textX), int32(textY), 25, rl.RayWhite)
		}()

		func() {
			for i := 0; i < 16; i++ {
				if bricks[i].active {
					if rl.CheckCollisionCircleRec(ballPos, 10, bricks[i].Rectangle) {
						ballVel.Y *= -1
						bricks[i].active = false
						score += 1
						brickCounter += 1
						break
					}
				}
			}
			for i := 0; i < 16; i++ {
				if bricks[i].active {
					rl.DrawRectangleRounded(bricks[i].Rectangle, .3, 1, rl.Color{255, 200, 11, 255})
				}
			}

		}()

		func() {
			if brickCounter == 16 && int(ballPos.Y) >= rl.GetScreenWidth()-25 {
				for i := range bricks {
					bricks[i].active = true
				}
				brickCounter = 0
				ballVel = rl.Vector2Scale(ballVel, 1.1)
			}
		}()

		func() {
			rl.DrawCircleV(ballPos, 10, rl.RayWhite)
			if start {
				dt := rl.GetFrameTime()
				ballPos = rl.Vector2Add(ballPos, rl.Vector2Scale(ballVel, dt))
			}
		}()

		func() {
			if ballPos.X >= (float32(rl.GetScreenWidth())-10) || ballPos.X <= 10 {
				ballVel.X *= -1
			}

			if ballPos.Y <= 10 {
				ballVel.Y *= -1
			}
			if ballPos.Y >= (float32(rl.GetScreenHeight()) - 10) {
				if !start {
					return
				}

				ballPos = rl.Vector2{
					X: footer.children[0].node.GetComputedLeft() + 10,
					Y: root.getRelativeComputedTop(footer.children[0]) + 10,
				}
				bar.X = float32(footer.children[1].node.GetComputedLeft())
				bar.Y = float32(root.getRelativeComputedTop(footer.children[1]))

				start = false
				lifeCount -= 1

				if lifeCount <= 0 {
					score = 0
				}
			}
		}()

		func() {
			rl.DrawRectangleRounded(bar, .5, 1, rl.Color{242, 101, 34, 255})
			if rl.CheckCollisionCircleRec(ballPos, 10, bar) && start {
				ballVel.Y *= -1
			}
		}()

		func() {
			if rl.IsKeyPressed(rl.KeySpace) {
				start = !start
			}

			dt := rl.GetFrameTime()
			if rl.IsKeyDown(rl.KeyRight) {
				bar.X += 400 * dt
				if !start {
					ballPos.X += 400 * dt
				}
			}

			if rl.IsKeyDown(rl.KeyLeft) {
				bar.X -= 400 * dt
				if !start {
					ballPos.X -= 400 * dt
				}
			}
		}()

		rl.EndDrawing()
	}
}
