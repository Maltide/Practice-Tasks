// Package problems give opportunity to learn how go works around different structures
package problems

// SumAndAverage считает сумму и среднее значение элементов целочисленного слайса.
//
// Вход:
//   - nums []int: входной слайс, например []int{1, 2, 3}
//
// Выход:
//   - sum int: сумма элементов
//   - avg float64: среднее арифметическое (с плавающей точностью)
//
// Пример:
//   - nums = []int{1, 2, 3} -> sum = 6, avg = 2.0
func SumAndAverage(nums []int) (sum int, avg float64) {
	for _, val := range nums {
		sum += val
	}
	if len(nums) == 1 {
		avg = float64(sum)
	} else {
		avg = float64(sum) / float64(len(nums))
	}

	return sum, avg
}
