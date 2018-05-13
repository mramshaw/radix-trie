package trie

import (
//"fmt"
)

// Node is a radix trie node (which may also be a leaf).
type Node struct {
	value      string
	children   []*Node
	childCount int
}

// IsLeaf may be called to determine if the current node is a leaf.
func (n *Node) IsLeaf() bool {
	return n.childCount == 0
}

func (n *Node) makeChildNode(s string) *Node {
	//fmt.Printf("makingChildNode: %s\n", s)
	child := makeNode(s)
	n.childCount++
	if n.children == nil {
		n.children = []*Node{&child}
	} else {
		n.children = append(n.children, &child)
	}
	return &child
}

func (n *Node) setChildNode(newNode *Node) bool {
	//fmt.Printf("settingChildNode: %v\n", newNode)
	n.childCount = 1
	n.children = []*Node{newNode}
	return true
}

func makeNode(s string) Node {
	//fmt.Printf("makingNode: %s\n", s)
	return Node{value: s, childCount: 0}
}
