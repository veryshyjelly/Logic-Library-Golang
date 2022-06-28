package logic

type Sentence interface {
	Evaluate(map[string]bool) bool
	Formula() string
	Symbols() map[string]bool
	Eq(interface{}) bool
}

func Parenthesize(s string) string {
	/*Parenthesizes an expression if not already parenthesized*/
	if (len(s) == 0) || isAlpha(s) || (s[0] == '(' && s[len(s)-1] == ')' && balanced(s[1:len(s)-1])) {
		return s
	}
	return "(" + s + ")"
}

func isAlpha(s string) bool {
	for _, v := range s {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
			return false
		}
	}
	return true
}

func balanced(s string) bool {
	/*Checks if a string has balanced parenthesis*/
	count := 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			if count <= 0 {
				return false
			}
			count--
		}
	}
	return count == 0
}
