package main

import "fmt"

type tree struct {
	head *node
}

type node struct {
	left, right *node
	v           int
}

func newTree() *tree {
	return &tree{}
}

// add - добавление значения в дерево
func addNode(t *tree, v int) {
	if t.head == nil {
		t.head = &node{v: v}
		return
	}

	curr := t.head
	for {
		if v < curr.v {
			if curr.left == nil {
				curr.left = &node{v: v}
				return
			}
			curr = curr.left
		} else if v > curr.v {
			if curr.right == nil {
				curr.right = &node{v: v}
				return
			}
			curr = curr.right
		} else {
			return
		}
	}
}

// remove - удаление значения из дерева
func removeFromTree(t *tree, v int) {
	t.head = removeNode(t.head, v)
}

// values - получение отсортированного слайса значений из дерева
func valuesTree(t *tree) []int {
	var res []int
	var traverse func(node *node)
	traverse = func(node *node) {
		if node != nil {
			traverse(node.left)
			res = append(res, node.v)
			traverse(node.right)
		}
	}
	traverse(t.head)
	return res
}

func removeNode(n *node, v int) *node {
	if n == nil {
		return nil
	}

	if v < n.v {
		n.left = removeNode(n.left, v)
	} else if v > n.v {
		n.right = removeNode(n.right, v)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minRight := n.right
		for minRight.left != nil {
			minRight = minRight.left
		}

		n.v = minRight.v

		n.right = removeNode(n.right, minRight.v)
	}

	return n
}

func main() {
	t := newTree()

	addNode(t, 10)
	addNode(t, 5)
	addNode(t, 15)
	addNode(t, 2)
	addNode(t, 7)
	addNode(t, 12)
	addNode(t, 20)

	fmt.Println(valuesTree(t))

	removeFromTree(t, 2) // без потомков
	fmt.Println(valuesTree(t))

	removeFromTree(t, 12)
	removeFromTree(t, 15) // c одним потомков
	fmt.Println(valuesTree(t))

	removeFromTree(t, 10) // с двумя потомками
	fmt.Println(valuesTree(t))
}
