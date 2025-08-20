package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
	head *Element
	size int
}

type Element struct {
	value int
	tail  *Element
}

func New(elements []int) *List {
	list := new(List)

	for i := 0; i < len(elements); i++ {
		list.Push(elements[i])
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.size++

	l.head = &Element{
		value: element,
		tail:  l.head,
	}
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return -1, errors.New("cannot Pop() when list is empty")
	}

	head := l.head

	l.head = head.tail
	l.size--

	return head.value, nil
}

func (l *List) Array() []int {
	ints, idx := make([]int, l.size), l.size-1

	for elem := l.head; elem != nil; elem = elem.tail {
		ints[idx] = elem.value
		idx--
	}

	return ints
}

func (l *List) Reverse() *List {
	list := new(List)

	for elem := l.head; elem != nil; elem = elem.tail {
		list.Push(elem.value)
	}

	return list
}
