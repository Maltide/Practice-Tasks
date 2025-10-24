package homework

import (
	"fmt"
	"time"
)

// Task 5: Rate Limiter Pattern
//
// OBJECTIVE: Process items with rate limiting
//
// PACKAGES TO USE:
// - time.NewTicker: https://pkg.go.dev/time#NewTicker
//   ticker := time.NewTicker(100 * time.Millisecond)
//   defer ticker.Stop()
//   <-ticker.C  // wait for tick
//
// HINT: Create ticker, wait for tick before processing each item

// RateLimitedProcessor processes items with rate limiting
func RateLimitedProcessor(items []string, maxPerSecond int) <-chan string {
	// TODO: Implement rate limiting
	// 1. Calculate interval between items based on maxPerSecond
	// 2. Create ticker with that interval
	// 3. Create output channel
	// 4. Launch goroutine that waits for ticker before processing each item
	// 5. Send processed items to output channel
	// 6. Close output channel and stop ticker when done
	// 7. Return output channel
	// Return closed channel to prevent hanging in tests
	ch := make(chan string, len(items))
	if len(items) == 0 {
		close(ch)
		return ch
	}
	if maxPerSecond <= 0 {
		for _, val := range items {
			ch <- val
		}
		close(ch)
		return ch
	}
	if len(items) == 1 {
		ch <- items[0]
		close(ch)
		return ch
	}
	ticker := time.NewTicker(time.Second / time.Duration(maxPerSecond))
	fmt.Printf("ticker = %v\n", ticker)

	go func() {

		fmt.Println("Starting for loop")
		for i := 0; i < len(items); i++ {
			fmt.Println("Trying to read signal from ticker.C")
			<-ticker.C
			fmt.Println("Signal was found")
			ch <- items[i]
			fmt.Printf("current item %v was write in ch\n", items[i])
		}
		ticker.Stop()
		fmt.Println("ticker stopped")
		close(ch)
		fmt.Println("channel closed")
	}()

	return ch
}
