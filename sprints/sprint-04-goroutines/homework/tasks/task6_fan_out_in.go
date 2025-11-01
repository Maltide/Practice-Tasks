package homework

import (
	"sync"
)

// Task 6: Fan-Out/Fan-In Pattern
//
// OBJECTIVE: Distribute work to workers (fan-out), collect results (fan-in)
// PROCESSING RULE: Each number should be tripled (multiplied by 3)
//
// PACKAGES TO USE:
// - sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
// - Channels for distribution and collection
//
// HINT: Create input channel, start workers reading from it, merge outputs
// NOTE: You can use a single channel for simplicity, or separate input/output channels
// for clearer separation of concerns. Both approaches are valid.

// FanOutFanIn distributes work across workers and collects results
func FanOutFanIn(numbers []int, numWorkers int) []int {
	if len(numbers) == 0 || numWorkers <= 0 || numbers == nil {
		return []int{}
	}
	var wg sync.WaitGroup

	// Create a semaphore channel to limit concurrent goroutines to numWorkers
	// This ensures we don't exceed the specified worker limit
	//effective := min(len(numbers), numWorkers) // Use minimum of workers and numbers
	tokench := make(chan int, numWorkers)

	// Buffered channel to collect processed results
	outputch := make(chan int, len(numbers))

	// Slice to store final results (shared across goroutines)
	resultslice := []int{}

	// Launch one goroutine per number (not per worker)
	// Each goroutine will process exactly one number
	for i := 0; i < len(numbers); i++ {
		tokench <- 1 // Acquire semaphore slot
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Process the number (multiply by 3) and send to output channel
			outputch <- numbers[i] * 3
			// Immediately read back the result from the same channel
			// Append to shared result slice (potential race condition)
			<-tokench // Release semaphore slot
		}(i)
	}
	wg.Wait()
	close(outputch)
	close(tokench)
	for val := range outputch {
		resultslice = append(resultslice, val)
	}

	return resultslice
}
