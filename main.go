package main

import (
	"fmt"
	"go_ds_algos/linked_list"
)

func main() {
	// Create an instance of LinkedList
	l := linked_list.LinkedList{}

	// Adding elements to the linked list
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	fmt.Println("LinkedList after Adding elements:", l)

	// Removing elements from the linked list
	l.Remove(4)
	l.Remove(1)
	fmt.Println("LinkedList after removing elements:", l)

	// Adding element to the beginning of the linked list
	l.AddFirst(1)
	fmt.Println("LinkedList after Adding element to the beginning:", l)
}
