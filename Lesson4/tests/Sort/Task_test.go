package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "one element",
			input: []int{10},
		},
		{
			name:  "already sorted",
			input: []int{-5, -1, 0, 3, 9},
		},
		{
			name:  "reverse order",
			input: []int{9, 7, 5, 3, 1},
		},
		{
			name:  "with duplicates",
			input: []int{4, 2, 4, 1, 2, 1},
		},
		{
			name:  "negative and positive values",
			input: []int{3, -10, 0, 5, -1},
		},
		{
			name:  "task array",
			input: []int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append([]int(nil), tt.input...)
			expected := append([]int(nil), tt.input...)
			sort.Ints(expected)

			t.Logf("Проверяем сортировку: %s", tt.name)
			t.Logf("Входной массив: %v", tt.input)

			InsertionSort(got)

			t.Logf("Что получила программа после InsertionSort: %v", got)
			t.Logf("Что ожидали получить: %v", expected)

			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("InsertionSort() работает неправильно: got %v, want %v", got, expected)
			}
		})
	}
}
