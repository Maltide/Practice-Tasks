package stringsprint

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RLDecode выполняет декодирование строки из формата Run-Length Encoding (RLE).
//
// Задача:
//
//	Восстановить исходную строку из RLE-кодированного формата. Строка содержит
//	символы, за которыми могут следовать числа, указывающие количество повторений.
//	Если после символа нет числа, символ встречается один раз.
//
// Сигнатура функции:
//
//	func RLDecode(s string) (string, error)
//
// Вход:
//   - s string: RLE-кодированная строка
//
// Выход:
//   - string: декодированная исходная строка
//   - error: ошибка, если формат некорректен
//
// Примеры:
//
//	RLDecode("a2bc5a3") → ("aabcccccaaa", nil)
//	RLDecode("abcdef") → ("abcdef", nil) (нет чисел)
//	RLDecode("a2b2c2") → ("aabbcc", nil)
//	RLDecode("") → ("", nil) (пустая строка)
//	RLDecode("a") → ("a", nil) (один символ)
//	RLDecode("a4") → ("aaaa", nil)
//	RLDecode("a0") → ("", nil) (ноль повторений)
//	RLDecode("a12b3") → ("aaaaaaaaaaabbb", nil) (многозначные числа)
//	RLDecode("2a") → ("", error) (число без символа)
//	RLDecode("a-1") → ("", error) (отрицательное число)
//
// Ограничения:
//   - Входная строка содержит буквы и цифры
//   - Числа могут быть многозначными (например, "a12")
//   - Числа всегда неотрицательные
//   - Временная сложность: O(n + m), где n - длина входа, m - длина результата
//   - Пространственная сложность: O(m) для результирующей строки
//
// Особенности:
//   - Парсите символы и следующие за ними числа
//   - Если после символа нет числа, считайте количество равным 1
//   - Обработайте многозначные числа
//   - Возвращайте ошибку для некорректных форматов
//   - Обработайте случай с нулевым количеством повторений
//
// Идиоматичный Go:
//   - Используйте strings.Builder для построения результата
//   - Используйте strconv.Atoi для парсинга чисел
//   - Возвращайте описательные ошибки с fmt.Errorf
//   - Используйте unicode.IsDigit для проверки цифр
func RLDecode(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var builder strings.Builder
	count := ""
	current := rune(s[0])

	if unicode.IsDigit(current) {
		return "", fmt.Errorf("число без символа")
	} else if current == '-' {
		return "", fmt.Errorf("неправильный формат данных")
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '-' {
			return "", fmt.Errorf("отрицательное значение")
		}

		if unicode.IsDigit(rune(s[i])) {
			count += string(s[i])
			continue
		}

		if len(count) == 0 {
			builder.WriteRune(current)
			current = rune(s[i])
			count = ""
			continue
		}

		number, err := strconv.Atoi(count)
		if err != nil {
			return "", err
		}

		for range number {
			builder.WriteRune(current)
		}
		current = rune(s[i])
		count = ""

	}

	if len(count) == 0 {
		builder.WriteRune(current)
	} else {
		number, err := strconv.Atoi(count)
		if err != nil {
			return "", err
		}

		for range number {
			builder.WriteRune(current)
		}

	}

	return builder.String(), nil
}
