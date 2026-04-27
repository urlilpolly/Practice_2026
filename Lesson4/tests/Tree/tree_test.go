package bst

import (
	"reflect"
	"testing"
)

func intLess(a, b int) bool {
	return a < b
}

func TestTreeNewAndEmptyTree(t *testing.T) {
	t.Log("Проверяем создание пустого бинарного дерева поиска")

	tree := New[int](intLess)

	t.Logf("Len() = %d", tree.Len())
	t.Logf("Contains(10) = %v", tree.Contains(10))
	t.Logf("Values() = %v", tree.Values())
	t.Logf("Remove(10) = %v", tree.Remove(10))

	if tree.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", tree.Len())
	}
	if tree.Contains(10) {
		t.Fatal("empty tree must not contain values")
	}
	if !reflect.DeepEqual(tree.Values(), []int{}) {
		t.Fatalf("Values() = %v, want empty slice", tree.Values())
	}
	if tree.Remove(10) {
		t.Fatal("Remove() from empty tree must return false")
	}
}

func TestTreeInsertContainsLenAndSortedValues(t *testing.T) {
	t.Log("Проверяем Insert, Contains, Len и то, что Values возвращает элементы по возрастанию")

	tree := New[int](intLess)
	values := []int{5, 3, 7, 2, 4, 6, 8}

	for _, value := range values {
		t.Logf("Insert(%d)", value)
		if !tree.Insert(value) {
			t.Fatalf("Insert(%d) = false, want true", value)
		}
		t.Logf("После вставки Values() = %v, Len() = %d", tree.Values(), tree.Len())
	}

	if tree.Len() != len(values) {
		t.Fatalf("Len() = %d, want %d", tree.Len(), len(values))
	}
	for _, value := range values {
		got := tree.Contains(value)
		t.Logf("Contains(%d) = %v", value, got)
		if !got {
			t.Fatalf("Contains(%d) = false, want true", value)
		}
	}
	if tree.Contains(100) {
		t.Fatal("Contains(100) = true, want false")
	}
	t.Log("Contains(100) = false, значит отсутствующий элемент не найден — это правильно")

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	t.Logf("Ожидаемый Values(): %v", expected)
	t.Logf("Фактический Values(): %v", tree.Values())
	if !reflect.DeepEqual(tree.Values(), expected) {
		t.Fatalf("Values() = %v, want %v", tree.Values(), expected)
	}
}

func TestTreeInsertDuplicate(t *testing.T) {
	t.Log("Проверяем, что дубликаты не добавляются")

	tree := New[int](intLess)

	firstInsert := tree.Insert(5)
	secondInsert := tree.Insert(5)
	t.Logf("Первый Insert(5) = %v", firstInsert)
	t.Logf("Второй Insert(5) = %v", secondInsert)
	t.Logf("Len() = %d", tree.Len())
	t.Logf("Values() = %v", tree.Values())

	if !firstInsert {
		t.Fatal("first Insert(5) = false, want true")
	}
	if secondInsert {
		t.Fatal("second Insert(5) = true, want false")
	}
	if tree.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", tree.Len())
	}
}

func TestTreeRemoveLeafOneChildAndTwoChildren(t *testing.T) {
	t.Log("Проверяем Remove: удаление листа, узла с одним потомком и узла с двумя потомками")

	tree := New[int](intLess)
	for _, value := range []int{5, 3, 7, 2, 4, 6, 8, 1} {
		tree.Insert(value)
	}
	t.Logf("Исходное дерево Values(): %v", tree.Values())

	assertRemovedFromTree(t, tree, 4, []int{1, 2, 3, 5, 6, 7, 8}) // leaf
	assertRemovedFromTree(t, tree, 2, []int{1, 3, 5, 6, 7, 8})    // node with one child
	assertRemovedFromTree(t, tree, 5, []int{1, 3, 6, 7, 8})       // root with two children

	if tree.Remove(100) {
		t.Fatal("Remove(100) = true, want false")
	}
	t.Log("Remove(100) = false, значит отсутствующий элемент не удалился — это правильно")

	expected := []int{1, 3, 6, 7, 8}
	if !reflect.DeepEqual(tree.Values(), expected) {
		t.Fatalf("Values() = %v, want %v", tree.Values(), expected)
	}
	if tree.Len() != len(expected) {
		t.Fatalf("Len() = %d, want %d", tree.Len(), len(expected))
	}
	for _, removed := range []int{2, 4, 5} {
		if tree.Contains(removed) {
			t.Fatalf("tree still contains removed value %d", removed)
		}
	}
}

func assertRemovedFromTree(t *testing.T, tree *Tree[int], value int, expectedValues []int) {
	t.Helper()

	before := tree.Values()
	removed := tree.Remove(value)
	t.Logf("Remove(%d) из %v -> %v", value, before, removed)
	t.Logf("После удаления Values() = %v", tree.Values())

	if !removed {
		t.Fatalf("Remove(%d) = false, want true", value)
	}
	if tree.Contains(value) {
		t.Fatalf("Contains(%d) = true after Remove(%d), want false", value, value)
	}
	if !reflect.DeepEqual(tree.Values(), expectedValues) {
		t.Fatalf("Values() = %v, want %v", tree.Values(), expectedValues)
	}
}
