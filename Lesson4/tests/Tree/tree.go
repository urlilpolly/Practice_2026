package bst

import "errors"

type Tree[T any] struct {
	root *node[T]
	size int
	less func(a, b T) bool
}

type node[T any] struct {
	left  *node[T]
	right *node[T]
	value T
}

func New[T any](less func(a, b T) bool) *Tree[T] {
	if less == nil {
		errors.New("Фунцкия не должна возвращать nil")
	}

	return &Tree[T]{less: less}
}

func (t *Tree[T]) Len() int {
	return t.size
}

func (t *Tree[T]) Insert(value T) bool {
	if t.root == nil {
		t.root = &node[T]{value: value}
		t.size = 1
		return true
	}

	current := t.root
	for {
		switch {
		case t.less(value, current.value):
			if current.left == nil {
				current.left = &node[T]{value: value}
				t.size++
				return true
			}
			current = current.left
		case t.less(current.value, value):
			if current.right == nil {
				current.right = &node[T]{value: value}
				t.size++
				return true
			}
			current = current.right
		default:
			return false
		}
	}
}

func (t *Tree[T]) Contains(value T) bool {
	current := t.root
	for current != nil {
		switch {
		case t.less(value, current.value):
			current = current.left
		case t.less(current.value, value):
			current = current.right
		default:
			return true
		}
	}

	return false
}

func (t *Tree[T]) Remove(value T) bool {
	var removed bool
	t.root, removed = t.removeNode(t.root, value)
	if removed {
		t.size--
	}
	return removed
}

func (t *Tree[T]) removeNode(current *node[T], value T) (*node[T], bool) {
	if current == nil {
		return nil, false
	}

	switch {
	case t.less(value, current.value):
		var removed bool
		current.left, removed = t.removeNode(current.left, value)
		return current, removed
	case t.less(current.value, value):
		var removed bool
		current.right, removed = t.removeNode(current.right, value)
		return current, removed
	default:
		if current.left == nil {
			return current.right, true
		}
		if current.right == nil {
			return current.left, true
		}

		successor := current.right
		for successor.left != nil {
			successor = successor.left
		}

		current.value = successor.value
		var removed bool
		current.right, removed = t.removeNode(current.right, successor.value)
		return current, removed
	}
}

func (t *Tree[T]) Values() []T {
	result := make([]T, 0, t.size)

	var traverse func(*node[T])
	traverse = func(n *node[T]) {
		if n == nil {
			return
		}

		traverse(n.left)
		result = append(result, n.value)
		traverse(n.right)
	}

	traverse(t.root)
	return result
}
