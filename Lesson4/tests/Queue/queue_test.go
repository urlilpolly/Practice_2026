package queue

import (
	"reflect"
	"testing"
)

func TestQueueNewInvalidCapacity(t *testing.T) {
	t.Log("Проверяем, что очередь нельзя создать с capacity <= 0")

	if q, err := New[int](0); err == nil || q != nil {
		t.Fatalf("New(0) = (%v, %v), want (nil, error)", q, err)
	} else {
		t.Logf("New(0) вернул ошибку: %v", err)
	}
	if q, err := New[int](-1); err == nil || q != nil {
		t.Fatalf("New(-1) = (%v, %v), want (nil, error)", q, err)
	} else {
		t.Logf("New(-1) вернул ошибку: %v", err)
	}
}

func TestQueueNewAndEmptyQueue(t *testing.T) {
	t.Log("Проверяем создание пустой очереди")

	q, err := New[int](3)
	if err != nil {
		t.Fatalf("New(3) returned unexpected error: %v", err)
	}

	t.Logf("Len() = %d", q.Len())
	t.Logf("Cap() = %d", q.Cap())
	t.Logf("IsEmpty() = %v", q.IsEmpty())
	t.Logf("IsFull() = %v", q.IsFull())
	t.Logf("Values() = %v", q.Values())

	if q.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", q.Len())
	}
	if q.Cap() != 3 {
		t.Fatalf("Cap() = %d, want 3", q.Cap())
	}
	if !q.IsEmpty() {
		t.Fatal("IsEmpty() = false, want true")
	}
	if q.IsFull() {
		t.Fatal("IsFull() = true, want false")
	}
	if q.Values() != nil {
		t.Fatalf("Values() = %v, want nil for empty queue", q.Values())
	}
	if _, err := q.Peek(); err == nil {
		t.Fatal("Peek() error = nil, want error for empty queue")
	} else {
		t.Logf("Peek() у пустой очереди вернул ошибку: %v", err)
	}
	if _, err := q.Dequeue(); err == nil {
		t.Fatal("Dequeue() error = nil, want error for empty queue")
	} else {
		t.Logf("Dequeue() у пустой очереди вернул ошибку: %v", err)
	}
}

func TestQueueEnqueuePeekDequeueAndFull(t *testing.T) {
	t.Log("Проверяем Enqueue, Peek, Dequeue и заполненную очередь")

	q, err := New[string](2)
	if err != nil {
		t.Fatalf("New(2) returned unexpected error: %v", err)
	}

	for _, value := range []string{"first", "second"} {
		t.Logf("Enqueue(%q)", value)
		if err := q.Enqueue(value); err != nil {
			t.Fatalf("Enqueue(%q) returned unexpected error: %v", value, err)
		}
		t.Logf("Values() = %v, Len() = %d", q.Values(), q.Len())
	}

	if !q.IsFull() {
		t.Fatal("IsFull() = false, want true")
	}
	if err := q.Enqueue("third"); err == nil {
		t.Fatal("Enqueue(third) error = nil, want error for full queue")
	} else {
		t.Logf("Enqueue в заполненную очередь вернул ошибку: %v", err)
	}

	peeked, err := q.Peek()
	t.Logf("Peek() = %q, error = %v", peeked, err)
	if err != nil {
		t.Fatalf("Peek() returned unexpected error: %v", err)
	}
	if peeked != "first" {
		t.Fatalf("Peek() = %q, want first", peeked)
	}
	if q.Len() != 2 {
		t.Fatalf("Len() after Peek() = %d, want 2", q.Len())
	}

	assertDequeued(t, q, "first")
	assertDequeued(t, q, "second")
	if !q.IsEmpty() {
		t.Fatal("queue must be empty after two Dequeue() calls")
	}
}

func TestQueueCircularOrder(t *testing.T) {
	t.Log("Проверяем кольцевое поведение очереди: после удаления элементы можно добавлять в начало массива")

	q, err := New[int](3)
	if err != nil {
		t.Fatalf("New(3) returned unexpected error: %v", err)
	}

	for _, value := range []int{1, 2, 3} {
		t.Logf("Enqueue(%d)", value)
		if err := q.Enqueue(value); err != nil {
			t.Fatalf("Enqueue(%d) returned unexpected error: %v", value, err)
		}
	}
	t.Logf("Очередь после добавления 1, 2, 3: %v", q.Values())

	assertDequeued(t, q, 1)
	assertDequeued(t, q, 2)
	t.Logf("Очередь после удаления 1 и 2: %v", q.Values())

	for _, value := range []int{4, 5} {
		t.Logf("Enqueue(%d)", value)
		if err := q.Enqueue(value); err != nil {
			t.Fatalf("Enqueue(%d) returned unexpected error: %v", value, err)
		}
	}

	expectedValues := []int{3, 4, 5}
	t.Logf("Ожидаемый порядок: %v", expectedValues)
	t.Logf("Фактический порядок Values(): %v", q.Values())
	if !reflect.DeepEqual(q.Values(), expectedValues) {
		t.Fatalf("Values() = %v, want %v", q.Values(), expectedValues)
	}

	for _, expected := range expectedValues {
		assertDequeued(t, q, expected)
	}
	if q.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", q.Len())
	}
}

func assertDequeued[T comparable](t *testing.T, q *Queue[T], expected T) {
	t.Helper()

	before := q.Values()
	got, err := q.Dequeue()
	t.Logf("Dequeue() из %v -> получили %v, после удаления Values() = %v", before, got, q.Values())

	if err != nil {
		t.Fatalf("Dequeue() returned unexpected error: %v", err)
	}
	if got != expected {
		t.Fatalf("Dequeue() = %v, want %v", got, expected)
	}
}
