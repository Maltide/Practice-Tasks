package homework

import "fmt"

// Task 3: Pipeline Pattern
//
// OBJECTIVE: Create 3-stage pipeline: generate → square → filter even
//
// PACKAGES TO USE:
// - Channels: https://go.dev/tour/concurrency/2
//   out := make(chan int)
//   go func() { for i := 0; i < n; i++ { out <- i }; close(out) }()
//   for val := range ch { /* process */ }
//
// HINT: Each stage returns <-chan int, chain them together

// ProcessPipeline creates a 3-stage pipeline
func ProcessPipeline(n int) <-chan int {
	if n <= 0 || n == 1 {
		out := make(chan int)
		close(out)
		return out
	}
	// TODO: Chain the three stages together
	// 1. Call generate function to create first stage
	out := generate(n)
	// 2. Pass its output channel to square function
	out = square(out)
	// 3. Pass square output channel to filterEven function
	out = filterEven(out)
	// 4. Return the final output channel
	return out
}

func generate(n int) <-chan int {
	// TODO: Generate numbers from 1 to n
	// 1. Create output channel
	out := make(chan int, n)
	// 2. Launch goroutine that loops from 1 to n
	go func() {
		for i := 1; i <= n; i++ {
			// 3. Send each number to output channel
			out <- i
		}
		// 4. Close channel when loop completes
		close(out)
	}()
	// 5. Return output channel immediately
	return out
}

func square(in <-chan int) <-chan int {
	// TODO: Square each number from input channel
	out := make(chan int, len(in))
	// 1. Create output channel

	// 2. Launch goroutine that ranges over input channel
	go func() {
		for val := range in {
			// 3. For each number, calculate its square
			// 4. Send squared number to output channel
			out <- val * val
		}
		// 5. Close output channel when input is exhausted
		close(out)
	}()
	// 6. Return output channel immediately
	return out
}

func filterEven(in <-chan int) <-chan int {
	// TODO: Filter and pass only even numbers
	// 1. Create output channel
	out := make(chan int, len(in))
	// 2. Launch goroutine that ranges over input channel
	go func() {
		for val := range in {
			// 3. For each number, check if it's even
			if val%2 == 0 {
				fmt.Printf("even val № %v\n", val)
				// 4. Send only even numbers to output channel
				out <- val
				fmt.Printf("send ")
			}
		}
		// 5. Close output channel when input is exhausted
		close(out)
	}()
	// 6. Return output channel immediately

	return out
}
