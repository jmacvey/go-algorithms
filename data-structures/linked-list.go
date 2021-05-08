package ds

import "errors"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func MakeLinkedList() LinkedList { 
	return LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *LinkedList) IsEmpty() bool { 
	return l.size == 0
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList) Get(index int) (int, error) {
	if index >= l.size {
		return 0, errors.New("invalid index")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data, nil
}

func (l *LinkedList) Add(value int) {
	newTail := Node{data: value, next: nil}
	if l.tail != nil {
		l.tail.next = &newTail
	} 
	l.tail = &newTail
	l.size++

	if (l.size == 1) {
		l.head = l.tail
	}
}

func (l *LinkedList) Remove(index int) (int, error) {
	if index >= l.size {
		return 0, errors.New("invalid index")
	}

	var v int
	if (index == 0) {
		v = l.head.data
		l.head = l.head.next
	} else {
		prev := l.head
		for i := 1; i < index; i++ {
			prev = prev.next
		}

		if (prev.next == l.tail) {
			l.tail = prev
		}

		v = prev.next.data
		prev.next = prev.next.next
	}

	l.size--
	if (l.size <= 1) {
		l.tail = l.head
	}

	return v, nil
}

func (l *LinkedList) Insert(value int, index int) error {
	if index >= l.size {
		return errors.New("invalid index")
	}

	newValue := Node{data:value, next: nil}

	if (index == 0) {
		newValue.next = l.head
		l.head = &newValue
	} else {
		previous := l.head
		for i := 1; i < index; i++ {
			previous = previous.next
		}

		newValue.next = previous.next
		previous.next = &newValue
	}
	
	l.size++
	return nil
}

func (l *LinkedList) Reverse() {
	if (l.size > 1) {
		previous := l.head
		next := l.head.next
		for next != nil {
			tmp := next.next
			next.next = previous
			previous = next
			next = tmp
		}

		tmp := l.head
		l.head = l.tail
		l.tail = tmp
		l.tail.next = nil
	}
}
