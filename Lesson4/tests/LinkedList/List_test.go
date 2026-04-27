package linkedlist

import (
	"reflect"
	"testing"
)

func TestListNewAndEmptyList(t *testing.T) {
	t.Log("Проверяем создание пустого связного списка")

	list := New[int]()

	t.Logf("Len() = %d", list.Len())
	t.Logf("IsEmpty() = %v", list.IsEmpty())
	t.Logf("Values() = %v", list.Values())

	if list.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", list.Len())
	}
	if !list.IsEmpty() {
		t.Fatal("IsEmpty() = false, want true")
	}
	if !reflect.DeepEqual(list.Values(), []int{}) {
		t.Fatalf("Values() = %v, want empty slice", list.Values())
	}
}

func TestListAppendGetAndValues(t *testing.T) {
	t.Log("Проверяем Append, Get, Values и Len")

	list := New[string]()
	for _, value := range []string{"a", "b", "c"} {
		t.Logf("Append(%q)", value)
		list.Append(value)
	}

	t.Logf("После добавления Values() = %v", list.Values())
	t.Logf("После добавления Len() = %d", list.Len())

	if list.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", list.Len())
	}
	if list.IsEmpty() {
		t.Fatal("IsEmpty() = true, want false")
	}

	expectedValues := []string{"a", "b", "c"}
	if !reflect.DeepEqual(list.Values(), expectedValues) {
		t.Fatalf("Values() = %v, want %v", list.Values(), expectedValues)
	}

	for index, expected := range expectedValues {
		got, err := list.Get(index)
		t.Logf("Get(%d) = %q, error = %v", index, got, err)
		if err != nil {
			t.Fatalf("Get(%d) returned unexpected error: %v", index, err)
		}
		if got != expected {
			t.Fatalf("Get(%d) = %q, want %q", index, got, expected)
		}
	}
}

func TestListGetInvalidIndex(t *testing.T) {
	t.Log("Проверяем, что Get возвращает ошибку для неверных индексов")

	list := New[int]()
	list.Append(10)

	if value, err := list.Get(-1); err == nil {
		t.Fatalf("Get(-1) = (%d, nil), want error", value)
	} else {
		t.Logf("Get(-1) вернул ошибку: %v", err)
	}
	if value, err := list.Get(1); err == nil {
		t.Fatalf("Get(1) = (%d, nil), want error", value)
	} else {
		t.Logf("Get(1) вернул ошибку: %v", err)
	}
}

func TestListRemoveFirstMiddleLastAndOnlyElement(t *testing.T) {
	t.Log("Проверяем Remove: удаление первого, среднего, последнего и единственного элемента")

	list := New[int]()
	for _, value := range []int{10, 20, 30, 40} {
		list.Append(value)
	}
	t.Logf("Исходный список: %v", list.Values())

	assertRemoved(t, list, 0, 10, []int{20, 30, 40})
	assertRemoved(t, list, 1, 30, []int{20, 40})
	assertRemoved(t, list, 1, 40, []int{20})
	assertRemoved(t, list, 0, 20, []int{})

	if !list.IsEmpty() {
		t.Fatal("list must be empty after removing all elements")
	}
	if list.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", list.Len())
	}
}

func TestListRemoveInvalidIndex(t *testing.T) {
	t.Log("Проверяем, что Remove возвращает ошибку для неверных индексов")

	list := New[int]()
	list.Append(10)

	if value, err := list.Remove(-1); err == nil {
		t.Fatalf("Remove(-1) = (%d, nil), want error", value)
	} else {
		t.Logf("Remove(-1) вернул ошибку: %v", err)
	}
	if value, err := list.Remove(1); err == nil {
		t.Fatalf("Remove(1) = (%d, nil), want error", value)
	} else {
		t.Logf("Remove(1) вернул ошибку: %v", err)
	}
}

func assertRemoved(t *testing.T, list *List[int], index int, expectedRemoved int, expectedValues []int) {
	t.Helper()

	before := list.Values()
	got, err := list.Remove(index)
	t.Logf("Remove(%d) из %v -> удалено %d, после удаления Values() = %v", index, before, got, list.Values())

	if err != nil {
		t.Fatalf("Remove(%d) returned unexpected error: %v", index, err)
	}
	if got != expectedRemoved {
		t.Fatalf("Remove(%d) = %d, want %d", index, got, expectedRemoved)
	}
	if !reflect.DeepEqual(list.Values(), expectedValues) {
		t.Fatalf("Values() = %v, want %v", list.Values(), expectedValues)
	}
	if list.Len() != len(expectedValues) {
		t.Fatalf("Len() = %d, want %d", list.Len(), len(expectedValues))
	}
}
