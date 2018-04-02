package trie

import (
//"fmt"
)

type node struct {
	value      string
	children   []*node
	childCount int
}

func (n *node) IsLeaf() bool {
	return n.childCount == 0
}

func (n *node) makeChildNode(s string) *node {
	//fmt.Printf("makingChildNode: %s\n", s)
	child := makeNode(s)
	n.childCount++
	if n.children == nil {
		n.children = []*node{&child}
	} else {
		n.children = append(n.children, &child)
	}
	return &child
}

func (n *node) setChildNode(newNode *node) bool {
	//fmt.Printf("settingChildNode: %v\n", newNode)
	n.childCount = 1
	n.children = []*node{newNode}
	return true
}

func makeNode(s string) node {
	//fmt.Printf("makingNode: %s\n", s)
	return node{value: s, childCount: 0}
}
