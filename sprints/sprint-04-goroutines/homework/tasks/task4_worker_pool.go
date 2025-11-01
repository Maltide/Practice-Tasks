package homework

import (
	"context"
	"sync"
)

// Task 4: Worker Pool Pattern
//
// OBJECTIVE: Process jobs with fixed number of workers
// PROCESSING RULE: Each job (number) should be doubled (multiplied by 2)
//
// PACKAGES TO USE:
// - sync.WaitGroup: https://pkg.go.dev/sync#WaitGroup
// - context: https://pkg.go.dev/context
//   ctx.Done() returns <-chan struct{} that closes on cancellation
//
// HINT: Create jobs channel, start workers, send jobs, collect results
// NOTE: You can use a single channel for simplicity, or separate input/output channels
// for clearer separation of concerns. Both approaches are valid.

// WorkerPool processes jobs using fixed number of workers
var wg sync.WaitGroup

// WorkerPool receives int from slice and return it doubled through channels in resultslice
func WorkerPool(jobs []int, numWorkers int) []int {
	// TODO: Implement worker pool
	if len(jobs) == 0 || numWorkers <= 0 {
		return []int{}
	}
	tokench := make(chan int, numWorkers)
	resulslice := []int{}
	// 1. Create jobs channel and results channel
	jobch := make(chan int, len(jobs))

	// 2. Start fixed number of worker goroutines
	// 3. Each worker reads from jobs channel, processes job, sends to results
	for i := 0; i < len(jobs); i++ {
		tokench <- 1
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			jobch <- jobs[i] * 2
			<-tokench
		}(i)
	}
	wg.Wait()
	close(jobch)
	// 4. Send all jobs to jobs channel and close it
	// 5. Collect all results from results channel
	for val := range jobch {
		resulslice = append(resulslice, val)
	}
	// 6. Return results slice
	return resulslice
}

// WorkerPoolWithContext adds cancellation support
func WorkerPoolWithContext(ctx context.Context, jobs []int, numWorkers int) ([]int, error) {
	if numWorkers <= 0 || len(jobs) == 0 {
		// пустой вход — пустой результат, без ошибки
		return nil, nil
	}
	// TODO: Implement worker pool with context cancellation
	// 1. Create jobs and results channels
	jobch := make(chan int, len(jobs))
	tokench := make(chan int, numWorkers)
	resulslice := []int{}
	// 2. Start workers that check context cancellation in select
	for _, val := range jobs {
		tokench <- 1
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			jobch <- v
			<-tokench
		}(val)
	}
	wg.Wait()
	close(jobch)
	close(tokench)

	for val := range jobch {
		resulslice = append(resulslice, val)
	}

	select {
	case <-ctx.Done():
		return resulslice, ctx.Err()
	default:
	}
	// 3. Send jobs or stop if context is cancelled
	// 4. Collect results or return error if cancelled
	// 5. Return partial results and context error if cancelled
	return resulslice, nil
}
