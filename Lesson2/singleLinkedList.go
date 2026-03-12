package main

import (
	"errors"
	"fmt"
)

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// add - добавление значения в связный список
func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}

	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) (any, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New("out of range")
	}

	curr := l.first
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr.v, nil
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) error {
	if idx < 0 || idx >= l.size {
		return errors.New("out of range")
	}

	if idx == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		curr := l.first
		for i := 0; i < idx-1; i++ {
			curr = curr.next
		}
		curr.next = curr.next.next

		if idx == l.size-1 {
			l.last = curr
		}
	}
	l.size--
	return nil
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	res := make([]any, 0, l.size)

	curr := l.first
	for curr != nil {
		res = append(res, curr.v)
		curr = curr.next
	}

	return res
}

func main() {
	l := newSinglyLinkedList()

	add(l, "dog")
	add(l, "cat")
	add(l, "cow")

	fmt.Println(values(l))

	val, _ := get(l, 1)
	fmt.Println("Get index 1:", val)

	remove(l, 1)
	fmt.Println("values without index 1", values(l))

	add(l, "horse")
	fmt.Println(values(l))
}
