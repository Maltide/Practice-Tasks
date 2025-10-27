package homework

import (
	"errors"
	"time"
)

// Task 7: Timeout Pattern
//
// OBJECTIVE: Process data with timeout constraint
//
// The timeout pattern is crucial for building robust concurrent applications.
// It prevents operations from blocking indefinitely, which can lead to resource
// exhaustion or unresponsive services. This function simulates processing a
// list of data items. The goal is to complete the processing within a specified
// time limit. If processing finishes before the timeout, the function returns
// the result. If the timeout expires before processing is complete, the
// function should return an error, indicating that the operation did not finish
// in time.
//
// Intern's Question:
// "У меня код не проходит один тест - "times_out" у которого на вход я так понимаю 1 миллисекунда
// и этот кейс говорит о том, что времени на выполнение очень мало, т.е. должна первее выйти ошибка
// с истечением времени. У меня этого не происходит, при этом, все остальные тесты проходят.
// У меня создалось впечатление, что горутина просто выполняет инкрементацию быстрее чем заданный
// таймаут и поэтому у меня выдает ошибку тест - что выглядит нелогично для тест-кейса текущего
// мб там надо микросекунду на вход давать тогда?"
//
// Answer to Intern's Question:
// Your observation is correct. The "times_out" test case provides a slice with 100 integers
// and a timeout of 1 millisecond. If your processing logic (e.g., simply iterating and counting)
// is very fast, it might complete before the 1ms timer expires, causing the test to fail because
// it expects a timeout error.
// The test is designed this way to ensure your implementation correctly uses the timeout mechanism.
// The key is not to make processing artificially slow, but to structure your code using `select`
// and `time.After` so that the timeout can interrupt the processing goroutine if necessary.
// The test timeout (1ms) is intentionally short relative to the amount of "work" (processing 100 items),
// to reliably trigger the timeout path in the `select` statement. Using an even shorter timeout like
// 1 microsecond might make it more likely to trigger, but the core issue is the implementation of the
// timeout logic itself.
//
// PACKAGES TO USE:
// - time.After: https://pkg.go.dev/time#After
//   timeout := time.After(5 * time.Second)
//   select { case <-timeout: /* handle timeout */ }
//
// HINT: Use select with time.After and done channel
//
// Definition of Done:
// - The function correctly processes the data and returns the count along with nil error if it finishes before the timeout.
// - The function returns 0 and ErrTimeout if the timeout duration is exceeded before processing completes.
// - The implementation uses a goroutine for processing and correctly employs `select` and `time.After` to manage the timeout.
// - No goroutines are leaked. The processing goroutine should naturally finish or be designed to not block indefinitely.
//
// Example Test Case (from tasks_test.go):
// t.Run("times out", func(t *testing.T) {
//     data := makeRange(1, 100) // Creates a slice [1, 2, ..., 100]
//     _, err := ProcessWithTimeout(data, 1*time.Millisecond)
//     // This test expects err to be non-nil (specifically ErrTimeout)
//     // because processing 100 items should take longer than 1ms,
//     // thus triggering the timeout path in the select statement.
//     if err == nil {
//         t.Error("ProcessWithTimeout() expected timeout error")
//     }
// })

var ErrTimeout = errors.New("processing timeout exceeded")

// ProcessWithTimeout processes data with a timeout
func ProcessWithTimeout(data []int, timeout time.Duration) (int, error) {
	// TODO: Implement processing with timeout

	// 1. Create done channel for completion signal (e.g., chan struct{} or chan int for count)
	if timeout <= 0 {
		return 0, ErrTimeout
	}
	if len(data) == 0 {
		return 0, nil
	}

	count := 0

	// 1. Create done channel for completion signal
	donech := make(chan int, 1)
	timech := time.After(timeout)
	// 2. Launch goroutine to process all data items

	//    Inside the goroutine:
	//    a. Perform the processing (e.g., count items, simulate work)
	//    b. Send a signal (e.g., count or struct{}) to the done channel when finished

	// 3. Use a `select` statement to wait for either:
	//    a. The done channel to receive the completion signal
	//    b. The timeout channel created by `time.After(timeout)`

	// 4. If the done channel receives first, return the result (count) and nil error

	// 5. If the timeout channel receives first, return 0 and ErrTimeout

	// 6. Ensure no goroutines are leaked. The goroutine should finish naturally.
	//    The select statement ensures that once one case is chosen, the function returns,
	//    and the other branch (if still running) is effectively ignored by the function's scope.

	// 3. Use select statement with timeout channel
	go func() {
		for i := 0; i < len(data); i++ {
			count++
		}
		donech <- count
	}()

	select {
	case <-donech:
		return count, nil
	case <-timech:
		return 0, ErrTimeout
	}

	// 4. If done channel receives first, return count
	// 5. If timeout channel receives first, return error
	// 6. Return processed count or timeout error

}
