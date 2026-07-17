package internal

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/husnulhamidiah/goarkanoid/constant"
	"github.com/husnulhamidiah/goarkanoid/flex"
)

type Game struct {
	Lives int
	Score int

	Hearts []Ball
	Ball   Ball
	Paddle Paddle
	Bricks []Brick

	State c.GameState

	brickCounter int
	revealTimer  float32

	scorePos      rl.Vector2
	ballOrigPos   rl.Vector2
	paddleOrigPos rl.Vector2
	gameOverPos   rl.Vector2
	spawning      bool
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Free() {
	flex.Free()
}

func (g *Game) Init() {
	flex.ComputeFlexLayout()

	g.Lives = 3
	g.Score = 0

	for i := 0; i < g.Lives; i++ {
		g.Hearts = append(g.Hearts, Ball{
			Vector2: rl.Vector2{
				X: flex.Root.GetRelativeComputedLeft(flex.Header.Children[0].Children[i]) + c.LiveCircleRadius,
				Y: flex.Root.GetRelativeComputedTop(flex.Header.Children[0].Children[i]) + c.LiveCircleRadius,
			},
			radius: c.LiveCircleRadius,
			color:  c.MiffyOrange,
		})
	}

	paddedScore := fmt.Sprintf("%04d", g.Score)
	scoreTextWidth := rl.MeasureText(paddedScore, c.HeaderFontSize)
	scoreTextRightEdge := flex.Header.Children[2].Node.GetComputedLeft() + flex.Header.Children[2].Node.GetComputedWidth()
	g.scorePos = rl.Vector2{
		X: scoreTextRightEdge - float32(scoreTextWidth),
		Y: flex.Header.Children[1].Node.GetComputedTop(),
	}

	g.Ball = Ball{
		Vector2: rl.Vector2{
			X: flex.Footer.Children[0].Node.GetComputedLeft() + c.BallRadius,
			Y: flex.Root.GetRelativeComputedTop(flex.Footer.Children[0]) + c.BallRadius,
		},
		radius: c.BallRadius,
		velocity: rl.Vector2{
			X: c.BallXVelocity,
			Y: c.BallYVelocity,
		},
		color: rl.RayWhite,
	}

	g.Paddle = Paddle{
		Rectangle: rl.Rectangle{
			X:      flex.Footer.Children[1].Node.GetComputedLeft(),
			Y:      flex.Root.GetRelativeComputedTop(flex.Footer.Children[1]),
			Width:  c.BarWidth,
			Height: c.BarHeight,
		},
		velocity: rl.Vector2{
			X: c.BarVelocity,
			Y: 0,
		},
	}

	for i := 0; i < c.BrickCount; i++ {
		g.Bricks = append(g.Bricks, Brick{
			Rectangle: rl.Rectangle{
				X:      flex.Content.Children[i].Node.GetComputedLeft(),
				Y:      flex.Root.GetRelativeComputedTop(flex.Content.Children[i]),
				Width:  float32(c.BrickWidth),
				Height: float32(c.BrickHeight),
			},
			active: false,
			color:  c.MiffyYellow,
		})
	}
	g.State = c.Initial

	g.brickCounter = 0
	g.revealTimer = 0.0

	g.ballOrigPos = g.Ball.Vector2
	g.paddleOrigPos = rl.Vector2{
		X: flex.Footer.Children[1].Node.GetComputedLeft(),
		Y: flex.Root.GetRelativeComputedTop(flex.Footer.Children[1]),
	}
	g.gameOverPos = rl.Vector2{
		X: flex.Header.Children[0].Node.GetComputedLeft(),
		Y: flex.Header.Children[0].Node.GetComputedTop(),
	}

	g.spawning = true
}

func (g *Game) Run() {
	for !rl.WindowShouldClose() {
		g.Input()
		g.Update()
		g.Draw()
	}
}
