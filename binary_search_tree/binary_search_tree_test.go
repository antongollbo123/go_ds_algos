package binary_search_tree

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)

	expected := "2 3 4 5 6 7 8 "
	if bst.String() != expected {
		t.Errorf("Expected %s, got %s", expected, bst.String())
	}
}

func TestSearch(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)

	// Test existing value
	if _, found := bst.Search(3); !found {
		t.Errorf("Expected to find value 3")
	}

	// Test non-existing value
	if _, found := bst.Search(10); found {
		t.Errorf("Did not expect to find value 10")
	}
}

func TestRemove(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)

	bst.Remove(3)
	bst.Remove(7)
	bst.Remove(5)

	expected := "2 4 6 8 "
	if bst.String() != expected {
		t.Errorf("Expected %s, got %s", expected, bst.String())
	}
}

func TestString(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)

	expected := "2 3 4 5 6 7 8 "
	if bst.String() != expected {
		t.Errorf("Expected %s, got %s", expected, bst.String())
	}
}

func TestInOrderTraversal(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(2)
	bst.Add(4)
	bst.Add(6)
	bst.Add(8)

	var sb strings.Builder
	bst.inOrderTraversal(&sb)

	expected := "2 3 4 5 6 7 8 "
	if sb.String() != expected {
		t.Errorf("Expected %s, got %s", expected, sb.String())
	}
}

func TestNodeString(t *testing.T) {
	node := Node{Value: 5}

	expected := "5"
	if node.String() != expected {
		t.Errorf("Expected %s, got %s", expected, node.String())
	}
}
