package queue

import "fmt"

type Queue[T any] struct {
	data      []T
	low, high int
	size      int
	capacity  int
}

func New[T any](capacity int) (*Queue[T], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("Вместимость должна быть больше 0")
	}

	return &Queue[T]{
		data:     make([]T, capacity),
		low:      -1,
		high:     -1,
		size:     0,
		capacity: capacity,
	}, nil
}

func (q *Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) Cap() int {
	return q.capacity
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue[T]) Enqueue(value T) error {
	if q.IsFull() {
		return fmt.Errorf("Очередь заполнена")
	}

	if q.IsEmpty() {
		q.low = 0
		q.high = 0
		q.data[q.high] = value
		q.size = 1
		return nil
	}
	q.high = (q.high + 1) % q.capacity
	q.data[q.high] = value
	q.size++

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, fmt.Errorf("Очередь пуста")
	}

	value := q.data[q.low]
	q.data[q.low] = zero

	if q.size == 1 {
		q.low = -1
		q.high = -1
		q.size = 0
		return value, nil
	}

	q.low = (q.low + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, fmt.Errorf("Очередь пуста")
	}

	return q.data[q.low], nil
}

func (q *Queue[T]) Values() []T {
	if q.IsEmpty() {
		return nil
	}
	result := make([]T, 0, q.size)
	index := q.low

	for i := 0; i < q.size; i++ {
		result = append(result, q.data[index])
		index = (index + 1) % q.capacity
	}
	return result
}
