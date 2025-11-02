package problems

// IsValidParentheses checks if a string contains valid parentheses
// Valid parentheses must be closed in the correct order
// Supported pairs: '()', '[]', '{}'
func IsValidParentheses(s string) bool {
	// TODO: implement
	if len(s) == 0 {
		return false
	}
	stack := []rune{}
	for _, val := range s {
		switch val {
		case '(', '[', '{':
			stack = append(stack, rune(val))
			continue
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
				continue
			}
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
				continue
			}
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
				continue
			}
		}
	}
	return true
}
