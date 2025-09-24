package problems

import "math"

func FindSmallestIntInMapMap(m map[int]map[int]int) int {
	// Эта функция находит наименьшее значение в вложенной мапе целых чисел.
	//
	// Аргумент:
	//   - m map[int]map[int]int: вложенная мапа, где ключи и значения являются целыми числами
	//   - Пример: map[int]map[int]int{
	//       1: {1: 10, 2: 5},
	//       2: {1: 3, 2: 8},
	//     }
	//
	// Алгоритм:
	//   - Придумай самостоятельно, используя логику похожую на поиск минимального числа в обычной мапе
	//
	// Что может понадобиться:
	//   - Для прохода по внешней мапе: for outerKey, innerMap := range m { ... }
	//   - Для прохода по внутренней мапе: for innerKey, value := range innerMap { ... }
	//   - Обрати внимание на случаи:
	//     * Пустая внешняя мапа
	//     * Пустые внутренние мапы
	//     * Внутренние мапы с разным количеством элементов
	//
	// Пример для мапы map[int]map[int]int{1: {1: 10, 2: 5}, 2: {1: 3, 2: 8}}:
	//   - Возвращаем 3, так как это наименьшее значение во всей вложенной структуре
	if len(m) == 0 {
		return 0
	}
	var minval int = math.MaxInt

	for _, innerMap := range m {
		for _, value := range innerMap {
			if value < minval {
				minval = value
			}
		}
	}

	return minval
}
