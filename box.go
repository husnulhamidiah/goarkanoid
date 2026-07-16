package main

import (
	goyoga "github.com/husnulhamidiah/goyoga/std"
)

type box struct {
	node     *goyoga.Node
	label    string
	children []*box
}

func newBox(config *goyoga.Config, label string) *box {
	return &box{node: goyoga.NewNodeWithConfig(config), label: label}
}

func newFixedBox(config *goyoga.Config, label string, width, height float32) *box {
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
