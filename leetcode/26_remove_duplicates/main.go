package _26

func removeDuplicates(nums []int) int {
	var unique_count int
	for i := 1; i < len(nums); i++ {
		if nums[unique_count] != nums[i] {
			unique_count++
			nums[unique_count] = nums[i]
		}
	}
	return unique_count + 1
}
