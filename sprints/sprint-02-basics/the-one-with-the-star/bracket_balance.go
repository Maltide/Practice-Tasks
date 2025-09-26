package problems

// IsBracketSequenceBalanced проверяет корректность расстановки скобок в строке.
//
// Вход:
//   - s string: строка, содержащая только символы '(', ')', '[', ']', '{', '}'
//
// Выход:
//   - bool: true, если каждая открывающая скобка закрывается соответствующей
//     скобкой в верном порядке; иначе false.
//
// Пример:
//   - s = "([]){}" -> true
//   - s = "([)]"   -> false
func IsBracketSequenceBalanced(s string) bool {

	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, byte(char))

		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]

			if pairs[byte(char)] != top {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
