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

func (b *BinarySearchTree) Remove(value int) {
	b.Root = b.RemoveByNode(b.Root, value)
	if b.Root != nil {
		b.Len--
	}
}

func (b *BinarySearchTree) RemoveByNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.LeftChild = b.RemoveByNode(root.LeftChild, value)
	} else if value > root.Value {
		root.RightChild = b.RemoveByNode(root.RightChild, value)
	} else {
		// Case 1: No child or only one child
		if root.LeftChild == nil {
			return root.RightChild
		} else if root.RightChild == nil {
			return root.LeftChild
		}

		// Case 2: Node with two children
		temp := b.minValueNode(root.RightChild)
		root.Value = temp.Value
		root.RightChild = b.RemoveByNode(root.RightChild, temp.Value)
	}

	return root
}

func (b *BinarySearchTree) minValueNode(node *Node) *Node {
	current := node
	for current.LeftChild != nil {
		current = current.LeftChild
	}
	return current
}

func (b BinarySearchTree) inOrderTraversalByNode(sb *strings.Builder, root *Node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.LeftChild)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.RightChild)
}

func (b *BinarySearchTree) InvertTree() {
	b.Root = b.invertTreeByNode(b.Root)
}

func (b *BinarySearchTree) invertTreeByNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.LeftChild, root.RightChild = root.RightChild, root.LeftChild
	b.invertTreeByNode(root.LeftChild)
	b.invertTreeByNode(root.RightChild)
	return root
}
