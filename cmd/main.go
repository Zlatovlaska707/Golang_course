package main

import (
	"fmt"
	"task_2/pkg/myFunc"
)

func main() {
	var number int
	fmt.Print("Выберите сортировку(1-быстрая; 2-вставками; 3-пузырьком): ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	result1 := myFunc.FindKthLargest(nums1, k1, number)
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums1, k1, result1)

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	result2 := myFunc.FindKthLargest(nums2, k2, number)
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums2, k2, result2)
}
