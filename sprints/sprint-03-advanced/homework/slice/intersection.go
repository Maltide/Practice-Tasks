package problems

// IntersectionUnique находит пересечение двух слайсов целых чисел, возвращая уникальные значения.
//
// Задача:
//
//	Найти все элементы, которые присутствуют в обоих слайсах. Каждый элемент должен
//	появиться в результате только один раз, даже если он повторяется в исходных слайсах.
//
// Сигнатура функции:
//
//	func IntersectionUnique(a, b []int) []int
//
// Вход:
//   - a []int: первый слайс целых чисел
//   - b []int: второй слайс целых чисел
//
// Выход:
//   - []int: слайс уникальных элементов, присутствующих в обоих входных слайсах
//
// Примеры:
//
//	IntersectionUnique([]int{1, 2, 2, 3}, []int{2, 3, 4}) → []int{2, 3}
//	IntersectionUnique([]int{1, 5}, []int{2, 3}) → []int{} (или nil)
//	IntersectionUnique([]int{}, []int{1, 2}) → []int{} (или nil)
//	IntersectionUnique([]int{1, 1, 1}, []int{1, 1}) → []int{1}
//
// Ограничения:
//   - Порядок элементов в результате не важен
//   - Пустые слайсы возвращают пустой результат
//   - Временная сложность: желательно O(n + m), где n и m - длины слайсов
//   - Пространственная сложность: O(min(n, m)) для хранения множества
//
// Особенности:
//   - Обработайте случаи с пустыми слайсами
//   - Учтите дубликаты в исходных данных
//   - Рассмотрите использование map для эффективного поиска
//
// Идиоматичный Go:
//   - Используйте map[int]bool или map[int]struct{} для создания множества
//   - Возвращайте nil для пустого результата или []int{}   // в итоге выяснилось, что слайс и пустые есть, инциализированные, и при этом nil-овые, т.е. стоило бы написать в этой строке "И" вместо "ИЛИ"
func IntersectionUnique(a, b []int) []int {
	var resultslice = []int{}

	// if a == nil || b == nil { насколько нужна данная проверка если в цикле далее все равно все автоскпинется из-за отсутсвия слаайсов как таковых и попадет на финальную проверку в конце кода???
	// 	return nil
	// }

	biggestslice := make(map[int]bool)
	duplicate := make(map[int]bool)

	if len(a) <= len(b) {
		for _, val := range b {
			biggestslice[val] = true
		}
		for _, val := range a {
			if duplicate[val] {
				continue
			}
			if biggestslice[val] {
				resultslice = append(resultslice, val)
				duplicate[val] = true
			}
		}
	} else {
		for _, val := range a {
			biggestslice[val] = true
		}
		for _, val := range b {
			if duplicate[val] {
				continue
			}
			if biggestslice[val] {
				resultslice = append(resultslice, val)
				duplicate[val] = true
			}
		}
	}

	if len(resultslice) == 0 {
		return nil
	}

	return resultslice
}
