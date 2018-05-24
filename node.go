package trie

import (
//"fmt"
)

// Node is a radix trie node (which may also be a leaf).
type Node struct {
	value      string
	children   []*Node
	childCount int
	entry      bool
}

// IsEntry may be called to determine if the current node is
// terminal for an entry. Note that this node may or may not
// also be a leaf. In the case of 'slow' and 'slowly', both
// are entries but only 'slowly' can be a leaf.
func (n *Node) IsEntry() bool {
	return n.entry
}

// IsLeaf may be called to determine if the current node is a
// leaf. Note that a leaf is a terminal node but that a node
// may also be terminal for an entry but not a leaf. In the
// case of 'slow' and 'slowly', 'slow' is NOT a leaf, even
// though it is terminal for the entry 'slow'.
func (n *Node) IsLeaf() bool {
	return n.childCount == 0
}

func (n *Node) makeChildNode(s string, entry bool) *Node {
	//fmt.Printf("makingChildNode: %s\n", s)
	child := makeNode(s, entry)
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

func makeNode(s string, isEntry bool) Node {
	//fmt.Printf("makingNode: %s\n", s)
	return Node{value: s, childCount: 0, entry: isEntry}
}
