package homework

import "sync"

// Task 8: Semaphore Pattern
//
// OBJECTIVE: Limit concurrent operations using buffered channel
//
// PACKAGES TO USE:
// - Buffered channel as semaphore:
//   sem := make(chan struct{}, maxConcurrent)
//   sem <- struct{}{}  // acquire
//   <-sem              // release
//
// HINT: Acquire before goroutine, defer release inside goroutine

// ConcurrentDownloader downloads files with max concurrent limit
func ConcurrentDownloader(urls []string, maxConcurrent int) int {
	if urls == nil || maxConcurrent <= 0 {
		return 0
	}
	// TODO: Implement semaphore pattern
	// 1. Create buffered channel as semaphore with size maxConcurrent
	bufferedch := make(chan struct{}, maxConcurrent)

	// 2. Create WaitGroup and success counter with mutex
	var wg sync.WaitGroup
	// 3. For each URL:
	//    a. Send to semaphore channel to acquire slot
	count := 0
	//    b. Launch goroutine
	for range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bufferedch <- struct{}{}
			count++
			<-bufferedch
		}()
	}
	wg.Wait()
	close(bufferedch)
	//    c. In goroutine, defer receive from semaphore to release slot
	//    d. Simulate download and increment success counter
	// 4. Wait for all goroutines to complete
	// 5. Return success count
	return count
}
