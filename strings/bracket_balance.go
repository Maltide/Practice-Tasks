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
	// 	var bstr []byte = []byte(s)
	// 	runebraces := []rune{'(', ')', '[', ']', '{', '}'}
	// 	bytebrace := make(map[byte]bool)
	// 	for i := range runebraces {
	// 		bytebrace[i] = byte(runebraces[i])
	// 	}

	// 	for i := range bstr {
	// 		switch bstr[i] {
	// 		case '(':
	// 			if bstr[i+1] == ')' {
	// 				continue
	// 			} else {
	// 				return false
	// 			}
	// 		case '[':
	// 			if bstr[i+1] == ']' {
	// 				continue
	// 			} else {
	// 				return false
	// 			}
	// 		case '{':
	// 			if bstr[i+1] == '}' {
	// 				continue
	// 			} else {
	// 				return false
	// 			}
	// 		default:
	// 			return true
	// 		}
	// 	}
	return true
}
