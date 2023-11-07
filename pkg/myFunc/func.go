package myFunc

import (
	"sort"
)

func FindKthLargest(nums []int, k, n int) int {
	if n == 1 {
		// Если n равно 1, то это сортировка встроенной функцией sort
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	} else if n == 2 {
		// Если n равно 2, то это сортировка вставками
		insertSort(nums)
	} else if n == 3 {
		// Если n равно 3, то это сортировка пузырьком
		bubbleSort(nums)
	} else {
		// В случае некорректного значения n возвращаем -1 или другое значение по умолчанию
		return -1
	}

	// Возвращаем -1, если значение k некорректное
	if k <= 0 || k > len(nums) {
		return -1
	}

	// Возвращаем k-й наибольший элемент
	return nums[k-1]
}
