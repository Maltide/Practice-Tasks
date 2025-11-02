package problems

import (
	"fmt"
	"strconv"
)

// EvalRPN evaluates an arithmetic expression in Reverse Polish Notation
// Tokens can be integers or operators: "+", "-", "*", "/"
// Division truncates toward zero
// All inputs represent valid expressions
func EvalRPN(tokens []string) int {
	// TODO: implement
	if len(tokens) == 0 {
		return 0
	}

	fmt.Println("token:", tokens)
	stack := []int{}
	result := 0

	for i := range tokens {
		switch tokens[i] {
		case "+":
			fmt.Printf("before last %v + %v (last number)\n", stack[len(stack)-2], stack[len(stack)-1])
			result = stack[len(stack)-2] + stack[len(stack)-1]
			fmt.Println("Result is", result)
			stack = stack[:len(stack)-1]
			fmt.Printf("Cut our stack to <before last number>, stack now is %v\n", stack)
			fmt.Println("Changing last number of stack to our result", result)
			stack[len(stack)-1] = result
			fmt.Printf("stack now is %v\n", stack)
			continue
		case "-":
			fmt.Printf("before last %v + %v (last number)\n", stack[len(stack)-2], stack[len(stack)-1])
			result = stack[len(stack)-2] - stack[len(stack)-1]
			fmt.Println("Result is", result)
			stack = stack[:len(stack)-1]
			fmt.Printf("Cut our stack to <before last number>, stack now is %v\n", stack)
			fmt.Println("Changing last number of stack to our result", result)
			stack[len(stack)-1] = result
			fmt.Printf("stack now is %v\n", stack)
			continue
		case "/":
			fmt.Printf("before last %v + %v (last number)\n", stack[len(stack)-2], stack[len(stack)-1])
			result = stack[len(stack)-2] / stack[len(stack)-1]
			fmt.Println("Result is", result)
			stack = stack[:len(stack)-1]
			fmt.Printf("Cut our stack to <before last number>, stack now is %v\n", stack)
			fmt.Println("Changing last number of stack to our result", result)
			stack[len(stack)-1] = result
			fmt.Printf("stack now is %v\n", stack)
			continue
		case "*":
			result = stack[len(stack)-2] * stack[len(stack)-1]
			fmt.Println("Result is", result)
			stack = stack[:len(stack)-1]
			fmt.Printf("Cut our stack to <before last number>, stack now is %v\n", stack)
			fmt.Println("Changing last number of stack to our result", result)
			stack[len(stack)-1] = result
			fmt.Printf("stack now is %v\n", stack)
			continue
		default:
			digit, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0
			}
			stack = append(stack, digit)
			fmt.Printf("Appending int %v to stack, stack now is %v\n", digit, stack)
			continue
		}
	}
	return result
}
