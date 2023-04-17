package _46_permutations

func permute(nums []int) (res [][]int) {
	current := []int{}

	var backtrack func([]int, []int)
	backtrack = func(nums []int, current []int) {
		if len(nums) == 0 {
			res = append(res, current)
			return
		}
		for i, _ := range nums {
			addedNum := nums[i]
			remainingNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
			backtrack(remainingNums, append(current, addedNum))
		}
	}
	backtrack(nums, current)
	return res
}
