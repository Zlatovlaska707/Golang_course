package myFunc

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	// Определяем набор тестовых случаев с входными данными и ожидаемыми результатами.
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 4}, []int{5, 4, 4, 3, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	// Проходим по каждому тестовому случаю и выполняем тест.
	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		insertSort(input)

		// Проверяем, соответствует ли результат ожидаемому результату.
		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("InsertSort(%v) => %v, expected %v", tc.input, input, tc.expected)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	// Определяем набор тестовых случаев с входными данными и ожидаемыми результатами.
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 4}, []int{5, 4, 4, 3, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	// Проходим по каждому тестовому случаю и выполняем тест.
	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		bubbleSort(input)

		// Проверяем, соответствует ли результат ожидаемому результату.
		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("BubbleSort(%v) => %v, expected %v", tc.input, input, tc.expected)
		}
	}
}
