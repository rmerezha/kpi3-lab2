package lab2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// PrefixToInfix converts a prefix mathematical expression into infix form.
// The input parameter `input` is a string representing the mathematical expression in prefix form,
// and the function returns a string representing the same expression in infix form.
//
// Parameters:
//   - input: A prefix mathematical expression in the form of a string that needs to be converted to infix form.
//     It is important that the expression is properly formatted and does not contain any extra spaces or errors.
//
// Returns:
//   - string: The infix expression equivalent to the input prefix expression.
//   - error: An error if the input expression is incorrect or cannot be converted.
//     If the conversion is successful, the error will be nil.
//
// Errors may occur if:
//   - The expression is not properly formatted (e.g., missing operands or operators).
//   - The input string is empty or contains invalid characters.
func PrefixToInfix(input string) (string, error) {
	tokens := strings.Split(input, " ")

	var stack []string
	for i := len(tokens) - 1; i >= 0; i-- {
		tok := tokens[i]

		if isOperator(tok) {
			if len(stack) < 2 {
				return "", errors.New("error: invalid expression")
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var expr string
			if i == 0 || tok == "*" || tok == "/" {
				expr = fmt.Sprintf("%s %s %s", op1, tok, op2)
			} else {
				expr = fmt.Sprintf("(%s %s %s)", op1, tok, op2)
			}
			stack = append(stack, expr)
		} else {
			_, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return "", err
			}
			stack = append(stack, tok)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("error: invalid expression")
	}

	return stack[0], nil
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}
