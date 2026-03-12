package main

import (
	"errors"
	"fmt"
)

type queue struct {
	s         []any // слайс в котором хранятся значения
	low, high int   // индексы верхней и нижней границы очереди
	size      int   // размер очереди
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

// push - добавление в очередь значения
func push(q *queue, v any) error {
	if (q.high+1)%q.size == q.low {
		return errors.New("Queue is full")
	}

	if q.low == -1 {
		q.low = 0
	}

	q.high = (q.high + 1) % q.size
	q.s[q.high] = v

	return nil
}

// pop - получения значения из очереди и его удаление
func pop(q *queue) (any, error) {
	if q.low == -1 {
		return nil, errors.New("queue is empty")
	}

	v := q.s[q.low]
	q.s[q.low] = nil

	if q.low == q.high {
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}
	return v, nil
}

func main() {
	q := newQueue(3)

	push(q, "one")
	push(q, "two")
	push(q, "three")

	if err := push(q, "four"); err != nil {
		fmt.Println(err)
	}

	val1, _ := pop(q)
	fmt.Println("Pop:", val1)

	push(q, "five")

	for i := 0; i < 3; i++ {
		val, err := pop(q)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Popped:", val)
		}
	}
}
