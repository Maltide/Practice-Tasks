# Stack Implementation Tasks

This directory contains stack-based problems and implementations.

## Files

### 1. stack.go
Basic stack implementation with the following operations:
- `Push(item int)` - Add element to top
- `Pop() (int, error)` - Remove and return top element
- `Peek() (int, error)` - View top element without removing
- `IsEmpty() bool` - Check if stack is empty
- `Size() int` - Get number of elements

### 2. valid_parentheses.go
**LeetCode 20: Valid Parentheses**

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
- Open brackets must be closed by the same type of brackets
- Open brackets must be closed in the correct order

### 3. min_stack.go
**LeetCode 155: Min Stack**

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:
- `Push(val int)` - Push element onto stack
- `Pop()` - Remove element on top of stack
- `Top() int` - Get the top element
- `GetMin() int` - Retrieve the minimum element

### 4. rpn_calculator.go
**LeetCode 150: Evaluate Reverse Polish Notation**

Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, and /. Each operand may be an integer or another expression.

## Implementation Notes

- All stack operations should be O(1) time complexity
- Use slices as the underlying data structure
- Handle edge cases (empty stack operations)
- Follow Go conventions for error handling

## Testing

Run tests with:
```bash
go test ./sprints/sprint-05-data-structures/homework -run TestStack