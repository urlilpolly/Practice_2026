package linkedlist

import "fmt"

type List[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

type node[T any] struct {
	value T
	next  *node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Append(value T) {
	newNode := &node[T]{value: value}

	if l.size == 0 {
		l.first = newNode
		l.last = newNode
		l.size = 1
		return
	}

	l.last.next = newNode
	l.last = newNode
	l.size++
}

func (l *List[T]) Get(index int) (T, error) {
	var zero T

	if index < 0 || index >= l.size {
		return zero, fmt.Errorf("%d вышел за границы", index)
	}

	current := l.first
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

func (l *List[T]) Remove(index int) (T, error) {
	var zero T

	if index < 0 || index >= l.size {
		return zero, fmt.Errorf("%d вышел за границы", index)
	}

	var removed *node[T]
	if index == 0 {
		removed = l.first
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		prev := l.first
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		removed = prev.next
		prev.next = removed.next
		if index == l.size-1 {
			l.last = prev
		}
	}

	l.size--
	return removed.value, nil
}

func (l *List[T]) Values() []T {
	result := make([]T, 0, l.size)

	for current := l.first; current != nil; current = current.next {
		result = append(result, current.value)
	}

	return result
}
