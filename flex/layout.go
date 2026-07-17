package flex

import (
	"fmt"

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
	Root.Node.SetWidth(300)
	Root.Node.SetHeight(600)

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
		Header.Children[0].addChild(newFixedBox(config, fmt.Sprintf("it-%d", i), 25, 25))
	}

	Content = newBox(config, "content")
	Content.Node.SetFlexGrow(1)
	Content.Node.SetPadding(goyoga.EdgeAll, 20)
	Content.Node.SetGap(goyoga.GutterAll, 5)
	Content.Node.SetJustifyContent(goyoga.JustifyCenter)
	Content.Node.SetFlexWrap(goyoga.WrapWrap)
	Content.Node.SetAlignContent(goyoga.AlignFlexStart)
	for i := 1; i <= 16; i++ {
		Content.addChild(newFixedBox(config, fmt.Sprintf("item-%02d", i), 50, 25))
	}

	Footer = newBox(config, "footer")
	Footer.Node.SetPadding(goyoga.EdgeAll, 15)
	Footer.Node.SetJustifyContent(goyoga.JustifyCenter)
	Footer.Node.SetAlignItems(goyoga.AlignCenter)
	Footer.Node.SetFlexDirection(goyoga.FlexDirectionColumn)
	Footer.addChild(newFixedBox(config, "dot", 20, 20))
	Footer.addChild(newFixedBox(config, "bar", 100, 10))

	Root.addChild(Header)
	Root.addChild(Content)
	Root.addChild(Footer)
	Root.Node.CalculateLayout(300, 600, goyoga.DirectionLTR)
}
