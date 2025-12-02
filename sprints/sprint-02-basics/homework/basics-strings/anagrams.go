// Package problems give opportunity to learn how go works around different structures
package problems

import "strings"

// AreAnagrams определяет, являются ли две строки анаграммами.
//
// Вход:
//   - a string: первая строка, например "listen"
//   - b string: вторая строка, например "silent"
//
// Выход:
//   - bool: true, если строки состоят из одинакового набора символов после
//     приведения к одному регистру и удаления пробелов; иначе false.
//
// Пример:
//   - a = "listen", b = "silent" -> true
//   - a = "rat",    b = "car"    -> false
func AreAnagrams(a, b string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	counter := make(map[rune]int)

	a, b = strings.ToLower(strings.Join(strings.Fields(a), "")),
		strings.ToLower(strings.Join(strings.Fields(b), ""))

	for _, val := range a {
		if _, ok := counter[val]; !ok {
			counter[val] = 0
		}
		counter[val]++
	}

	for _, val := range b {
		if _, ok := counter[val]; !ok {
			return false
		}
		counter[val]--
		if counter[val] < 0 {
			return false
		}
	}
	for _, total := range counter {
		if total != 0 {
			return false
		}
	}
	return true
}
