package problems

// MaxSlidingWindow returns the max sliding window for an array of integers
// Given an array of integers nums, there is a sliding window of size k
// which is moving from the very left of the array to the very right
// Return the max sliding window
func MaxSlidingWindow(nums []int, k int) []int {
	// TODO: implement
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return nums
	}

	maxval := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums)-k+1; i++ {
		curMax := nums[i]
		for j := i; j < i+k; j++ {
			if nums[j] > curMax {
				curMax = nums[j]
			}
		}
		maxval = append(maxval, curMax)
	}

	return maxval
}
