package linked_list

import (
	"testing"
)

func TestAdd(t *testing.T) {
	l := &LinkedList{}
	l.add(1)
	l.add(2)
	l.add(3)

	expected := "1 2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}
}

func TestGet(t *testing.T) {
	l := &LinkedList{}
	l.add(1)
	l.add(2)
	l.add(3)

	node := l.get(2)
	if node == nil || node.value != 2 {
		t.Errorf("Expected to get node with value 2, but got %v", node)
	}

	node = l.get(4)
	if node != nil {
		t.Errorf("Expected to get nil for non-existent value, but got %v", node)
	}
}

func TestRemove(t *testing.T) {
	l := &LinkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)

	l.remove(4)
	expected := "1 2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}

	l.remove(1)
	expected = "2 3 "
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}

	l.remove(2)
	l.remove(3)
	expected = ""
	if l.String() != expected {
		t.Errorf("Expected %s, got %s", expected, l.String())
	}
}
