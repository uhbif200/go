package main

import "fmt"

func ValidParentheses(parens string) bool {
	// Your code goes here
	var (
		open  int
		close int
	)

	for _, paren := range parens {
		if string(paren) == "(" {
			open += 1
		} else if string(paren) == ")" {
			close += 1
		}
		if close > open {
			return false
		}
	}
	return open == close
}

func main() {
	fmt.Printf("ValidParentheses(\"((()))\"): %v\n", ValidParentheses("((()))"))
	fmt.Printf("ValidParentheses(\"(()))\"): %v\n", ValidParentheses("(()))"))
	fmt.Printf("ValidParentheses(\")(\"): %v\n", ValidParentheses(")("))
}
