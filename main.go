package main

import (
	"fmt"
	matrix "go_ds_algos/basic_stats"
	"go_ds_algos/binary_search_tree"
)

func main() {

	m := matrix.Matrix{
		Content: [][]int{
			{0, 1, 2, 3},
			{4, 5, 6, 7},
		},
		Dim: 8}

	println(m.Mean())
	println(m.Variance(), m.Std())

	// Create an instance of LinkedList
	n := &binary_search_tree.Node{Value: 1, LeftChild: nil, RightChild: nil}
	n.LeftChild = &binary_search_tree.Node{Value: 0, LeftChild: nil, RightChild: nil}
	n.RightChild = &binary_search_tree.Node{Value: 5, LeftChild: nil, RightChild: nil}
	b := binary_search_tree.BinarySearchTree{Root: n, Len: 1}
	//fmt.Println(b)

	b.Add(2)
	b.Add(4)
	b.Add(10)

	fmt.Println(b)

	b.InvertTree()
	fmt.Println(b)

}
