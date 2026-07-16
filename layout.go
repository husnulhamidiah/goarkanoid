package main

import (
	"fmt"

	goyoga "github.com/husnulhamidiah/goyoga/std"
)

var (
	root    *box
	header  *box
	content *box
	footer  *box
)

func buildLayout() func() {
	config := goyoga.NewConfig()
	config.SetUseWebDefaults(true)

	root = newBox(config, "root")
	root.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	root.node.SetWidth(300)
	root.node.SetHeight(600)

	header = newBox(config, "header")
	header.node.SetPadding(goyoga.EdgeAll, 15)
	header.node.SetJustifyContent(goyoga.JustifySpaceBetween)
	header.node.SetAlignItems(goyoga.AlignCenter)

	header.addChild(newBox(config, "lives"))
	header.children[0].node.SetJustifyContent(goyoga.JustifySpaceBetween)
	header.children[0].node.SetAlignItems(goyoga.AlignCenter)
	header.children[0].node.SetGap(goyoga.GutterAll, 2)

	header.addChild(newFixedBox(config, fmt.Sprintf("top-%d", 2), 50, 25))
	header.addChild(newFixedBox(config, fmt.Sprintf("top-%d", 3), 50, 25))

	for i := 1; i <= 3; i++ {
		header.children[0].addChild(newFixedBox(config, fmt.Sprintf("it-%d", i), 25, 25))
	}

	content = newBox(config, "content")
	content.node.SetFlexGrow(1)
	content.node.SetPadding(goyoga.EdgeAll, 20)
	content.node.SetGap(goyoga.GutterAll, 5)
	content.node.SetJustifyContent(goyoga.JustifyCenter)
	content.node.SetFlexWrap(goyoga.WrapWrap)
	content.node.SetAlignContent(goyoga.AlignFlexStart)
	for i := 1; i <= 16; i++ {
		content.addChild(newFixedBox(config, fmt.Sprintf("item-%02d", i), 50, 25))
	}

	footer = newBox(config, "footer")
	footer.node.SetPadding(goyoga.EdgeAll, 15)
	footer.node.SetJustifyContent(goyoga.JustifyCenter)
	footer.node.SetAlignItems(goyoga.AlignCenter)
	footer.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	footer.addChild(newFixedBox(config, "dot", 20, 20))
	footer.addChild(newFixedBox(config, "bar", 100, 10))

	root.addChild(header)
	root.addChild(content)
	root.addChild(footer)
	root.node.CalculateLayout(300, 600, goyoga.DirectionLTR)

	return func() {
		config.Free()
		root.node.FreeRecursive()
	}
}
