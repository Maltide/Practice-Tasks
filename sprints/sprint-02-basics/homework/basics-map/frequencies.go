// Package problems give opportunity to learn how go works around different structures
package problems

// CountFrequencies строит частоты появления строк в слайсе.
//
// Вход:
//   - items []string: слайс значений, например []string{"a", "b", "a"}
//
// Выход:
//   - map[string]int: мапа, где ключ — элемент, значение — количество его вхождений
//
// Пример:
//   - items = []string{"a", "b", "a"} -> map[string]int{"a": 2, "b": 1}
func CountFrequencies(items []string) map[string]int {
	counter := make(map[string]int)

	for _, val := range items {
		if _, ok := counter[val]; !ok {
			counter[val] = 0
		}
		counter[val]++
	}
	return counter
}
