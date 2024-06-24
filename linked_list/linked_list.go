package linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (l *LinkedList) add(value int) {
	node := new(Node)
	node.value = value

	if l.head == nil {
		l.head = node
	} else {
		iter := l.head
		for ; iter.next != nil; iter = iter.next {

		}
		iter.next = node
	}
}

func (l LinkedList) get(value int) *Node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l *LinkedList) remove(value int) {

	var previous *Node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
			} else {
				previous.next = current.next
				return
			}
		}
		previous = current
	}

}

func (l LinkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s ", iterator))
	}
	return sb.String()
}

func main() {
	l := LinkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	fmt.Println(l)

	l.remove(4)
	l.remove(1)

	fmt.Println(l)

}
