package binary_search_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

type BinarySearchTree struct {
	Root *Node
	Len  int
}

func (n Node) String() string {
	return strconv.Itoa(n.Value)
}

func (b BinarySearchTree) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b BinarySearchTree) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.Root)
}

func (b *BinarySearchTree) Add(value int) {
	b.Root = b.AddByNode(b.Root, value)
	b.Len++
}

func (b *BinarySearchTree) AddByNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.LeftChild = b.AddByNode(root.LeftChild, value)
	}

	if value > root.Value {
		root.RightChild = b.AddByNode(root.RightChild, value)
	}
	return root
}

func (b BinarySearchTree) Search(value int) (*Node, bool) {

	return b.SearchByNode(b.Root, value)

}

func (b BinarySearchTree) SearchByNode(root *Node, value int) (*Node, bool) {

	if root == nil {
		return nil, false
	}

	if root.Value == value {
		return root, true
	} else if root.Value > value {
		return b.SearchByNode(root.LeftChild, value)
	} else {
		return b.SearchByNode(root.RightChild, value)
	}
}

func (b BinarySearchTree) inOrderTraversalByNode(sb *strings.Builder, root *Node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.LeftChild)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.RightChild)
}
