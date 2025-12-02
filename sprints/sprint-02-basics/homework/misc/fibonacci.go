// Package problems give opportunity to learn how go works around different structures
package problems

// Fibonacci генерирует первые n чисел последовательности Фибоначчи начиная с 0 и 1.
//
// Вход:
//   - n int: длина требуемой последовательности (n >= 1)
//
// Выход:
//   - []int: слайс длины n с числами Фибоначчи, например при n = 5 -> []int{0, 1, 1, 2, 3}
//
// Пример:
//   - n = 5 -> []int{0, 1, 1, 2, 3}
func Fibonacci(n int) []int {
	Fibonaccislice := []int{}
	num := 0
	for i := range n {
		if i == 0 {
			Fibonaccislice = append(Fibonaccislice, 0)
			continue
		}
		if i == 1 {
			Fibonaccislice = append(Fibonaccislice, 1)
			continue
		}
		num = Fibonaccislice[i-1] + Fibonaccislice[i-2]
		Fibonaccislice = append(Fibonaccislice, num)
	}

	return Fibonaccislice
}
