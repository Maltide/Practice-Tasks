package stringsprint

// LongestCommonPrefix находит самый длинный общий префикс среди массива строк.
//
// Задача:
//
//	Найти самую длинную строку, которая является префиксом для всех строк
//	в массиве. Если общего префикса нет, вернуть пустую строку.
//
// Сигнатура функции:
//
//	func LongestCommonPrefix(words []string) string
//
// Вход:
//   - words []string: массив строк для поиска общего префикса
//
// Выход:
//   - string: самый длинный общий префикс всех строк
//
// Примеры:
//
//	LongestCommonPrefix([]string{"flower", "flow", "flight"}) → "fl"
//	LongestCommonPrefix([]string{"dog", "racecar", "car"}) → ""
//	LongestCommonPrefix([]string{"interspecies", "interstellar", "interstate"}) → "inters"
//	LongestCommonPrefix([]string{"prefix", "prefix"}) → "prefix"
//	LongestCommonPrefix([]string{"a"}) → "a"
//	LongestCommonPrefix([]string{}) → ""
//	LongestCommonPrefix([]string{"", "abc"}) → ""
//
// Ограничения:
//   - Все строки состоят из строчных английских букв
//   - Массив может быть пустым
//   - Строки могут быть пустыми
//   - Временная сложность: O(S), где S - сумма всех символов
//   - Пространственная сложность: O(1) дополнительной памяти
//
// Особенности:
//   - Если любая строка пустая, общий префикс пустой
//   - Если массив пустой, возвращайте пустую строку
//   - Сравнивайте символы по позициям во всех строках
//   - Остановитесь при первом несовпадении или достижении конца любой строки
//
// Идиоматичный Go:
//   - Используйте range для итерации по символам
//   - Работайте с rune для корректной обработки Unicode
//   - Возвращайте подстроку исходной строки
func LongestCommonPrefix(words []string) string {
	// TODO: Реализуйте функцию
	return ""
}
