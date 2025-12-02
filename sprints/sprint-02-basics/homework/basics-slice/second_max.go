// Package problems give opportunity to learn how go works around different structures
package problems

import "math"

// SecondLargest находит второе по величине отличающееся значение в слайсе.
//
// Вход:
//   - nums []int: слайс как минимум с двумя различными значениями, например []int{5, 1, 5, 2}
//
// Выход:
//   - int: значение, которое является вторым максимумом
//
// Пример:
//   - nums = []int{5, 1, 5, 2} -> 2
//   - nums = []int{1, 2, 3}     -> 2
func SecondLargest(nums []int) int {
	firstmax := nums[0]

	secondlargest := math.MinInt

	for i := 1; i < len(nums); i++ {
		if nums[i] > firstmax {
			secondlargest = firstmax
			firstmax = nums[i]
		} else if nums[i] < firstmax && nums[i] > secondlargest {
			secondlargest = nums[i]
		}
	}
	return secondlargest
}
