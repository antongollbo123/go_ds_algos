package linked_list

import (
	"testing"
)

func TestAdd(t *testing.T) {
	l := &LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)

	expected := "1 2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}
}

func TestGet(t *testing.T) {
	l := &LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)

	node := l.Get(2)
	if node == nil || node.value != 2 {
		t.Errorf("Expected to Get node with value 2, but got %v", node)
	}

	node = l.Get(4)
	if node != nil {
		t.Errorf("Expected to Get nil for non-existent value, but got %v", node)
	}
}

func TestRemove(t *testing.T) {
	l := &LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	l.Remove(4)
	expected := "1 2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}

	l.Remove(1)
	expected = "2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}

	l.Remove(2)
	l.Remove(3)
	expected = ""
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}
}

func TestAddFirst(t *testing.T) {
	// Create a new LinkedList instance
	l := LinkedList{}

	// Test adding to an empty list
	l.AddFirst(1)

	// Check if the head is correctly set
	if l.head == nil {
		t.Errorf("AddFirst did not set the head correctly for an empty list")
	} else if l.head.value != 1 {
		t.Errorf("AddFirst did not set the head value correctly for an empty list. Expected: %d, Got: %d", 1, l.head.value)
	}

	// Test adding to a non-empty list
	l.AddFirst(2)

	// Check if the head is correctly updated
	if l.head.value != 2 {
		t.Errorf("AddFirst did not update the head value correctly for a non-empty list. Expected: %d, Got: %d", 2, l.head.value)
	}

	// Check if the next pointer of the new node points to the previous head
	if l.head.next == nil || l.head.next.value != 1 {
		t.Errorf("AddFirst did not correctly set the next pointer of the new node")
	}

	// Test string representation after adding elements
	expected := "2 1 "
	if l.String() != expected {
		t.Errorf("String representation after AddFirst is incorrect. Expected: %s, Got: %s", expected, l.String())
	}

	// Test adding more elements and checking consistency
	l.AddFirst(3)
	expected = "3 2 1 "
	if l.String() != expected {
		t.Errorf("String representation after multiple AddFirst operations is incorrect. Expected: %s, Got: %s", expected, l.String())
	}
}
