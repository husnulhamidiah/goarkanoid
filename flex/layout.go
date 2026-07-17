package flex

import (
	"fmt"

	c "github.com/husnulhamidiah/goarkanoid/constant"
	goyoga "github.com/husnulhamidiah/goyoga/std"
)

var (
	config  *goyoga.Config
	Root    *box
	Header  *box
	Content *box
	Footer  *box
)

func Free() {
	config.Free()
	Root.Node.FreeRecursive()
}

func ComputeFlexLayout() {
	config = goyoga.NewConfig()
	config.SetUseWebDefaults(true)

	Root = newBox(config, "root")
	Root.Node.SetFlexDirection(goyoga.FlexDirectionColumn)
	Root.Node.SetWidth(c.WindowWidth)
	Root.Node.SetHeight(c.WindowHeight)

	Header = newBox(config, "header")
	Header.Node.SetPadding(goyoga.EdgeAll, 15)
	Header.Node.SetJustifyContent(goyoga.JustifySpaceBetween)
	Header.Node.SetAlignItems(goyoga.AlignCenter)

	Header.addChild(newBox(config, "lives"))
	Header.Children[0].Node.SetJustifyContent(goyoga.JustifySpaceBetween)
	Header.Children[0].Node.SetAlignItems(goyoga.AlignCenter)
	Header.Children[0].Node.SetGap(goyoga.GutterAll, 2)

	Header.addChild(newFixedBox(config, fmt.Sprintf("top-%d", 2), 50, 25))
	Header.addChild(newFixedBox(config, fmt.Sprintf("top-%d", 3), 50, 25))

	for i := 1; i <= 3; i++ {
		Header.Children[0].addChild(newFixedBox(config, fmt.Sprintf("it-%d", i), c.LiveCircleRadius*2, c.LiveCircleRadius*2))
	}

	Content = newBox(config, "content")
	Content.Node.SetFlexGrow(1)
	Content.Node.SetPadding(goyoga.EdgeAll, 20)
	Content.Node.SetGap(goyoga.GutterAll, 5)
	Content.Node.SetJustifyContent(goyoga.JustifyCenter)
	Content.Node.SetFlexWrap(goyoga.WrapWrap)
	Content.Node.SetAlignContent(goyoga.AlignFlexStart)
	for i := 1; i <= c.BrickCount; i++ {
		Content.addChild(newFixedBox(config, fmt.Sprintf("item-%02d", i), c.BrickWidth, c.BrickHeight))
	}

	Footer = newBox(config, "footer")
	Footer.Node.SetPadding(goyoga.EdgeAll, 15)
	Footer.Node.SetJustifyContent(goyoga.JustifyCenter)
	Footer.Node.SetAlignItems(goyoga.AlignCenter)
	Footer.Node.SetFlexDirection(goyoga.FlexDirectionColumn)
	Footer.addChild(newFixedBox(config, "dot", c.BallRadius*2, c.BallRadius*2))
	Footer.addChild(newFixedBox(config, "bar", c.BarWidth, c.BarHeight))

	Root.addChild(Header)
	Root.addChild(Content)
	Root.addChild(Footer)
	Root.Node.CalculateLayout(c.WindowWidth, c.WindowHeight, goyoga.DirectionLTR)
}
