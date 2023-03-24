package _26

func removeDuplicate(nums []int, val int) int {
	no_val := 0
	for i, _ := range nums {
		if nums[i] != val {
			nums[no_val] = nums[i]
			no_val++
		}

	}
	return no_val
}
