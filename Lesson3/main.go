package main

import (
	"Practice2026/Lesson3/binaryTree"
	"Practice2026/Lesson3/linkedList"
	"Practice2026/Lesson3/queue"
	"fmt"
)

func main() {
	// Список
	list := linkedlist.New[string]()
	list.Append("dog")
	list.Append("cat")
	list.Append("cow")

	fmt.Println("list:", list.Values())
	value, _ := list.Get(1)
	fmt.Println("list[1]:", value)
	removed, _ := list.Remove(1)
	fmt.Println("removed:", removed)
	fmt.Println("list:", list.Values())

	// Очередь
	q, _ := queue.New[string](3)
	_ = q.Enqueue("one")
	_ = q.Enqueue("two")
	_ = q.Enqueue("three")
	fmt.Println("\nqueue:", q.Values())
	dequeued, _ := q.Dequeue()
	fmt.Println("dequeued:", dequeued)
	fmt.Println("queue:", q.Values())

	// Дерево
	tree := bst.New[int](func(a, b int) bool { return a < b })
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(20)
	fmt.Println("\ntree:", tree.Values())
	tree.Remove(10)
	fmt.Println("tree after remove:", tree.Values())
}
