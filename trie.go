package trie

import (
	//  "fmt"
	"strings"
)

// Trie is a radix trie implementation.
type Trie struct {
	child []*node
	count int
}

// NewTrie is used to create a new Trie.
func NewTrie() *Trie {
	return &Trie{}
}

// Count returns the number of nodes in the Trie.
func (t *Trie) Count() int {
	return t.count
}

func (t *Trie) isEmpty() bool {
	return t.count == 0
}

// Insert is used to add a new term to the Trie.
func (t *Trie) Insert(s string) bool {

	// remove leading & trailing whitespace
	trimmed := strings.TrimSpace(s)

	// Sanity check (should catch empty strings too)
	if len(trimmed) < 2 {
		return false
	}

	if t.count == 0 {
		t.makeRuneNode(s)
		return true
	}

	// Check for duplicate nodes
	if t.findNode(trimmed) != nil {
		return false
	}

	for _, c := range t.child {
		//fmt.Println()
		if c.value == s[:1] {
			//fmt.Printf("findNode child matched: %[01]v %[01]T\n", c)
			beheaded := s[1:]
			return t.insertRuneNode(c, c, beheaded)
		}
	}

	t.makeRuneNode(s)
	return true
}

func (t *Trie) insertRuneNode(parent *node, n *node, s string) bool {

	for _, c := range n.children {
		index := t.findRuneMatch(c.value, s)
		//fmt.Printf("insertRuneMatch: '%s' '%s' %d len(s) = %d\n", c.value, s, index, len(s))
		lenC := len(c.value)
		if index > 0 {
			if index == lenC {
				return t.insertRuneNode(parent, c, s[index:])
			}
			if index < lenC {
				child := c.makeChildNode(c.value[index:])
				child.children = c.children
				child.childCount = c.childCount
				c.setChildNode(child)
				//fmt.Printf("c.value: %s\n", c.value)
				c.value = c.value[:index]
				//fmt.Printf("c.value: %s\n", c.value)
			}
			//fmt.Printf("making child node: %s\n", s[index:])
			c.makeChildNode(s[index:])
			t.count++
			return true
		}
	}

	// No match in the children so attach to the parent node
	//fmt.Printf("parented.value2: %s\n", s)
	parent.makeChildNode(s)
	t.count++
	return true
}

func (t *Trie) makeRuneNode(s string) {
	rootRune := makeNode(s[:1])
	rootChild := makeNode(s[1:])
	rootRune.children = []*node{&rootChild}
	rootRune.childCount = 1
	t.child = []*node{&rootRune}
	t.count++
}

// Find is used to search for a specific term in the Trie.
func (t *Trie) Find(s string) bool {

	// Remove leading & trailing whitespace
	trimmed := strings.TrimSpace(s)

	// Sanity check (should catch empty strings too)
	if len(s) < 2 {
		return false
	}

	return t.findNode(trimmed) != nil
}

func (t *Trie) findNode(s string) *node {

	for _, c := range t.child {
		if c.value == s[:1] {
			//fmt.Println()
			beheaded := s[1:]
			return t.findRuneNode(c, beheaded)
		}
	}

	// Runaway check
	return nil
}

func (t *Trie) findRuneNode(n *node, s string) *node {

	for _, c := range n.children {
		index := t.findRuneMatch(c.value, s)
		//fmt.Printf("findRuneMatch: '%s' '%s' %d len(s) = %d\n", c.value, s, index, len(s))
		lenC := len(c.value)
		if index > 0 {
			if index == lenC && index == len(s) {
				return c
			}
			if c.value == s {
				return c
			}
			if c.childCount > 0 {
				return t.findRuneNode(c, s[index:])
			}
		}
	}
	return nil
}

func (t *Trie) findRuneMatch(v string, s string) int {

	res := 0
	sLen := len(s)
	runes := []rune(s)

	for i, c := range v {

		if i > sLen {
			return res
		}
		if c != runes[i] {
			return res
		}
		res++
	}
	return res
}
