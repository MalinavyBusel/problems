package _0_remove_duplicates_2

func removeDuplicates(nums []int) int {
	pointer := 0
	counts := 0
	num := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] == num && counts >= 2 {
			counts += 1
			continue
		}
		if nums[i] != num {
			num = nums[i]
			counts = 0
		}
		nums[pointer] = nums[i]
		pointer++
		counts++
	}
	return pointer
}
