package problems

import (
	"testing"
)

// BubbleSortStability helper function for testing sort stability
func BubbleSortStability(values []int, indices []int) {
	n := len(indices)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			// Сравниваем значения, а не индексы
			if values[indices[j]] > values[indices[j+1]] {
				indices[j], indices[j+1] = indices[j+1], indices[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

// InsertionSortStability helper function for testing sort stability
func InsertionSortStability(values []int, indices []int) {
	// Перемещаемся по массиву, начиная со второго элемента
	for i := 1; i < len(indices); i++ {
		keyIndex := indices[i] // Индекс текущего элемента для вставки
		j := i - 1             // Индекс последнего элемента отсортированной части

		// Перемещаем элементы отсортированной части, которые больше key,
		// на одну позицию вправо
		for j >= 0 && values[indices[j]] > values[keyIndex] {
			indices[j+1] = indices[j]
			j--
		}

		// Вставляем keyIndex на правильную позицию
		indices[j+1] = keyIndex
	}
}

// MergeSortStability helper function for testing sort stability
func MergeSortStability(values []int, indices []int) []int {
	// Базовый случай: массивы с 0 или 1 элементом уже отсортированы
	if len(indices) <= 1 {
		result := make([]int, len(indices))
		copy(result, indices)
		return result
	}

	// Рекурсивно сортируем левую и правую половины
	mid := len(indices) / 2
	left := MergeSortStability(values, indices[:mid])
	right := MergeSortStability(values, indices[mid:])

	// Объединяем отсортированные половины
	return mergeStability(values, left, right)
}

// mergeStability helper function for merging two sorted index slices
func mergeStability(values []int, left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Сравниваем значения по индексам и добавляем индекс с меньшим значением в результат
	for i < len(left) && j < len(right) {
		// Для стабильности используем <=, чтобы сохранить порядок равных элементов
		if values[left[i]] <= values[right[j]] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся индексы из левого слайса (если есть)
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Добавляем оставшиеся индексы из правого слайса (если есть)
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// TestStabilityBubbleSort tests bubble sort stability
func TestStabilityBubbleSort(t *testing.T) {
	// Создаем тестовые данные с парами (значение, оригинальный индекс)
	// Равные значения имеют разные индексы, что позволяет проверить стабильность
	original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
	expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

	// Извлекаем значения для сортировки
	values := make([]int, len(original))
	for i, pair := range original {
		values[i] = pair[0]
	}

	// Создаем индексы для отслеживания стабильности
	indices := make([]int, len(original))
	for i := range indices {
		indices[i] = i
	}

	// Сортируем индексы с помощью пузырьковой сортировки, сравнивая значения
	BubbleSortStability(values, indices)

	// Создаем отсортированный массив пар
	sorted := make([][2]int, len(original))
	for i, idx := range indices {
		sorted[i] = original[idx]
	}

	// Проверяем, что результат соответствует ожидаемому
	for i, pair := range sorted {
		if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
			t.Errorf("BubbleSort stability test failed. Expected %v, got %v", expected[i], pair)
		}
	}
}

// TestStabilityInsertionSort tests insertion sort stability
func TestStabilityInsertionSort(t *testing.T) {
	// Создаем тестовые данные с парами (значение, оригинальный индекс)
	// Равные значения имеют разные индексы, что позволяет проверить стабильность
	original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
	expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

	// Извлекаем значения для сортировки
	values := make([]int, len(original))
	for i, pair := range original {
		values[i] = pair[0]
	}

	// Создаем индексы для отслеживания стабильности
	indices := make([]int, len(original))
	for i := range indices {
		indices[i] = i
	}

	// Сортируем индексы с помощью сортировки вставками, сравнивая значения
	InsertionSortStability(values, indices)

	// Создаем отсортированный массив пар
	sorted := make([][2]int, len(original))
	for i, idx := range indices {
		sorted[i] = original[idx]
	}

	// Проверяем, что результат соответствует ожидаемому
	for i, pair := range sorted {
		if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
			t.Errorf("InsertionSort stability test failed. Expected %v, got %v", expected[i], pair)
		}
	}
}

// TestStabilityMergeSort tests merge sort stability
func TestStabilityMergeSort(t *testing.T) {
	// Создаем тестовые данные с парами (значение, оригинальный индекс)
	// Равные значения имеют разные индексы, что позволяет проверить стабильность
	original := [][2]int{{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4}, {3, 5}}
	expected := [][2]int{{1, 1}, {1, 4}, {2, 3}, {3, 0}, {3, 2}, {3, 5}}

	// Извлекаем значения для сортировки
	values := make([]int, len(original))
	for i, pair := range original {
		values[i] = pair[0]
	}

	// Создаем индексы для отслеживания стабильности
	indices := make([]int, len(original))
	for i := range indices {
		indices[i] = i
	}

	// Сортируем индексы с помощью сортировки слиянием, сравнивая значения
	sortedIndices := MergeSortStability(values, indices)

	// Создаем отсортированный массив пар
	sorted := make([][2]int, len(original))
	for i, idx := range sortedIndices {
		sorted[i] = original[idx]
	}

	// Проверяем, что результат соответствует ожидаемому
	for i, pair := range sorted {
		if pair[0] != expected[i][0] || pair[1] != expected[i][1] {
			t.Errorf("MergeSort stability test failed. Expected %v, got %v", expected[i], pair)
		}
	}
}

// RunStabilityTests runs all stability tests
func RunStabilityTests(t *testing.T) {
	t.Run("BubbleSort Stability", TestStabilityBubbleSort)
	t.Run("InsertionSort Stability", TestStabilityInsertionSort)
	t.Run("MergeSort Stability", TestStabilityMergeSort)
}
