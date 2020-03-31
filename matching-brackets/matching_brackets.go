package brackets

// Bracket checks if the brackets are balanced.
func Bracket(str string) bool {
	var stack []rune

	for _, c := range str {
		switch c {
		case '{', '[', '(':
			stack = append(stack, c)
		case '}', ']', ')':
			l := len(stack)
			if l == 0 {
				return false
			}
			lastBracket := stack[l-1]
			// Test all the valid conditions for the closing bracket
			if (lastBracket == '{' && c == '}') || (lastBracket == '[' && c == ']') || (lastBracket == '(' && c == ')') {
				stack = stack[:l-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
