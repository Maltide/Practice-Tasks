// Package problems give opportunity to learn how go works around different structures
package problems

import "strings"

// IsPalindrome проверяет, является ли строка палиндромом.
//
// Вход:
//   - s string: произвольная строка, например "A man a plan a canal Panama"
//
// Выход:
//   - bool: true, если строка читается одинаково слева направо и справа налево после
//     приведения к одному регистру и удаления пробелов; иначе false.
//
// Пример:
//   - s = "abba" -> true
//   - s = "abc"  -> false
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.Join(strings.Fields(s), ""))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
