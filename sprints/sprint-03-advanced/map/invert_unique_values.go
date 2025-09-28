package mapsprint

// InvertUnique инвертирует map, меняя местами ключи и значения.
//
// Задача:
//
//	Создать новый map, где ключи исходного map становятся значениями,
//	а значения становятся ключами. Функция предполагает, что все значения
//	в исходном map уникальны.
//
// Сигнатура функции:
//
//	func InvertUnique(m map[string]int) map[int]string
//
// Вход:
//   - m map[string]int: исходный map с уникальными значениями
//
// Выход:
//   - map[int]string: инвертированный map
//
// Примеры:
//
//	InvertUnique(map[string]int{"a": 1, "b": 2, "c": 3}) → map[int]string{1: "a", 2: "b", 3: "c"}
//	InvertUnique(map[string]int{"hello": 42, "world": 13}) → map[int]string{42: "hello", 13: "world"}
//	InvertUnique(map[string]int{}) → map[int]string{} (пустой map)
//	InvertUnique(map[string]int{"single": 100}) → map[int]string{100: "single"}
//	InvertUnique(map[string]int{"zero": 0, "negative": -5}) → map[int]string{0: "zero", -5: "negative"}
//
// Ограничения:
//   - Все значения в исходном map уникальны (предусловие)
//   - Map может быть пустым
//   - Значения могут быть отрицательными или нулевыми
//   - Временная сложность: O(n), где n - количество элементов в map
//   - Пространственная сложность: O(n) для нового map
//
// Особенности:
//   - Функция предполагает уникальность значений (не проверяет это)
//   - Обработайте пустой map
//   - Порядок элементов в map не гарантирован
//   - Создайте новый map, не изменяя исходный
//
// Идиоматичный Go:
//   - Используйте make для создания нового map
//   - Итерируйтесь по исходному map с помощью range
//   - Возвращайте пустой map для пустого входа
func InvertUnique(m map[string]int) map[int]string {
	// TODO: Реализуйте функцию
	return nil
}
