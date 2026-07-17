package flex

import (
	goyoga "github.com/husnulhamidiah/goyoga/std"
)

type box struct {
	Node     *goyoga.Node
	Label    string
	Children []*box
}

func newBox(config *goyoga.Config, label string) *box {
	return &box{Node: goyoga.NewNodeWithConfig(config), Label: label}
}

func newFixedBox(config *goyoga.Config, label string, width, height float32) *box {
	b := newBox(config, label)
	b.Node.SetWidth(width)
	b.Node.SetHeight(height)
	return b
}

func (b *box) addChild(child *box) {
	b.Node.InsertChild(child.Node, uint(len(b.Children)))
	b.Children = append(b.Children, child)
}

func (b *box) GetRelativeComputedLeft(target *box) (left float32) {
	node := target.Node
	for node != nil && node != b.Node {
		left += node.GetComputedLeft()
		node = node.GetOwner()
	}
	return
}

func (b *box) GetRelativeComputedTop(target *box) (top float32) {
	node := target.Node
	for node != nil && node != b.Node {
		top += node.GetComputedTop()
		node = node.GetOwner()
	}
	return
}
