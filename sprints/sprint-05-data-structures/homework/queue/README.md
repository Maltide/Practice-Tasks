# Queue Implementation Tasks

This directory contains queue-based problems and implementations.

## Files

### 1. queue.go
Basic queue implementation with the following operations:
- `Enqueue(item int)` - Add element to rear
- `Dequeue() (int, error)` - Remove and return front element
- `Front() (int, error)` - View front element without removing
- `IsEmpty() bool` - Check if queue is empty
- `Size() int` - Get number of elements

### 2. circular_queue.go
**LeetCode 622: Design Circular Queue**

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO principle and the last position is connected back to the first position to make a circle.

Implementation should support following operations:
- `EnQueue(value int) bool` - Insert element into circular queue
- `DeQueue() bool` - Delete element from circular queue
- `Front() int` - Get front item (-1 if empty)
- `Rear() int` - Get last item (-1 if empty)
- `IsEmpty() bool` - Check whether circular queue is empty
- `IsFull() bool` - Check whether circular queue is full

### 3. stack_using_queues.go
**LeetCode 225: Implement Stack using Queues**

Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:
- `Push(x int)` - Push element x onto stack
- `Pop() int` - Remove and return element on top of stack
- `Top() int` - Return element on top of stack
- `Empty() bool` - Return true if stack is empty

### 4. sliding_window_maximum.go
**LeetCode 239: Sliding Window Maximum**

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.

## Implementation Notes

- Queue operations should be O(1) time complexity
- For circular queue, handle wrap-around correctly
- For sliding window, use deque for O(n) solution
- Follow Go conventions for error handling

## Testing

Run tests with:
```bash
go test ./sprints/sprint-05-data-structures/homework -run TestQueue