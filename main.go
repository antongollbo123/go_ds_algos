package main

import (
	"fmt"
	"go_ds_algos/binary_search_tree"
)

func main() {
	// Create an instance of LinkedList
	n := &binary_search_tree.Node{Value: 1, LeftChild: nil, RightChild: nil}
	n.LeftChild = &binary_search_tree.Node{Value: 0, LeftChild: nil, RightChild: nil}
	n.RightChild = &binary_search_tree.Node{Value: 5, LeftChild: nil, RightChild: nil}
	b := binary_search_tree.BinarySearchTree{Root: n, Len: 1}
	fmt.Println(b)

	b.Add(2)
	b.Add(4)
	b.Add(10)

	fmt.Println(b)

	node, res := b.Search(10)

	fmt.Println(node, res)

	fmt.Println(b.Search(19))

	b.Remove(10)
	b.Remove(4)

	fmt.Println(b)

}
