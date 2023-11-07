package myFunc

// insertSort выполняет сортировку среза вставками в порядке убывания.
func insertSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] < key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// bubbleSort выполняет сортировку среза пузырьком в порядке убывания.
func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
