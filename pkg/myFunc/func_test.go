package myFunc

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	// Определяем набор тестовых случаев с входными данными и ожидаемыми результатами.
	testCases := []struct {
		nums     []int
		k        int
		n        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 1, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 1, 4},
		{[]int{1, 2, 3, 4, 5}, 1, 2, 5},
		{[]int{5, 4, 3, 2, 1}, 3, 2, 3},
		{[]int{7, 1, 3, 8, 4, 5, 9, 6, 2}, 6, 3, 4},
		{[]int{5, 4, 3, 2, 1}, 3, 3, 3},
	}

	// Проходим по каждому тестовому случаю и выполняем тест.
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("FindKthLargest(%v, %d, %d)", tc.nums, tc.k, tc.n), func(t *testing.T) {
			// Вызываем функцию FindKthLargest с входными данными из текущего тестового случая.
			result := FindKthLargest(tc.nums, tc.k, tc.n)
			// Проверяем, соответствует ли результат ожидаемому результату.
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}

		})
	}
}
