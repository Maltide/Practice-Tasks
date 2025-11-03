package homework

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

// Tests for functions in errors.go
// Run with: go test -race -v errors_test.go errors.go

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	if cerr := w.Close(); cerr != nil {
		log.Printf("captureOutput: close writer: %v", cerr)
	}

	os.Stdout = old

	var buf bytes.Buffer
	if _, cerr := io.Copy(&buf, r); cerr != nil {
		log.Printf("captureOutput: copy: %v", cerr)
	}
	return buf.String()
}

func TestPrintNumbers(t *testing.T) {
	t.Run("prints numbers and done", func(t *testing.T) {
		output := captureOutput(func() {
			PrintNumbers()
		})

		// Check that "Done" is printed
		if !strings.Contains(output, "Done") {
			t.Error("Expected 'Done' to be printed")
		}

		// Check that numbers 1-5 are printed (order may vary due to concurrency)
		for i := 1; i <= 5; i++ {
			expected := string(rune('0' + i))
			if !strings.Contains(output, expected) {
				t.Errorf("Expected number %d to be printed", i)
			}
		}

		// Count lines to ensure we have 6 lines (5 numbers + "Done")
		lines := strings.Split(strings.TrimSpace(output), "\n")
		if len(lines) != 6 {
			t.Errorf("Expected 6 lines of output, got %d", len(lines))
		}
	})
}

func TestPrintSquares(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected map[string]bool // Expected strings that should appear
	}{
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: map[string]bool{},
		},
		{
			name:    "single number",
			numbers: []int{3},
			expected: map[string]bool{
				"3 squared is 9": true,
			},
		},
		{
			name:    "multiple numbers",
			numbers: []int{2, 4},
			expected: map[string]bool{
				"2 squared is 4":  true,
				"4 squared is 16": true,
			},
		},
		{
			name:    "negative numbers",
			numbers: []int{-2, 3},
			expected: map[string]bool{
				"-2 squared is 4": true,
				"3 squared is 9":  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				PrintSquares(tt.numbers)
			})

			// Check that all expected strings appear in output
			for expectedStr := range tt.expected {
				if !strings.Contains(output, expectedStr) {
					t.Errorf("Expected '%s' to be in output, got: %s", expectedStr, output)
				}
			}

			// Count non-empty lines
			lines := strings.Split(strings.TrimSpace(output), "\n")
			nonEmptyLines := 0
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					nonEmptyLines++
				}
			}

			if len(tt.numbers) > 0 && nonEmptyLines != len(tt.numbers) {
				t.Errorf("Expected %d lines of output, got %d", len(tt.numbers), nonEmptyLines)
			}
		})
	}
}

func TestSumChannel(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "single number",
			numbers:  []int{5},
			expected: 5,
		},
		{
			name:     "multiple numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test with timeout to avoid potential hanging
			done := make(chan int, 1)
			go func() {
				result := SumChannel(tt.numbers)
				done <- result
			}()

			select {
			case result := <-done:
				if result != tt.expected {
					t.Errorf("SumChannel() = %d, want %d", result, tt.expected)
				}
			case <-time.After(100 * time.Millisecond):
				t.Error("SumChannel() timed out")
			}
		})
	}
}

func TestConcurrentCounter(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "zero iterations",
			n:        0,
			expected: 0,
		},
		{
			name:     "single iteration",
			n:        1,
			expected: 1,
		},
		{
			name:     "multiple iterations",
			n:        100,
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConcurrentCounter(tt.n)
			if result != tt.expected {
				t.Errorf("ConcurrentCounter() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestProcessData(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "empty slice",
			data:     []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			data:     []int{5},
			expected: []int{10}, // Only first element is processed
		},
		{
			name:     "multiple elements",
			data:     []int{1, 2, 3, 4},
			expected: []int{2}, // Only first element is processed (1*2=2)
		},
		{
			name:     "negative numbers",
			data:     []int{-3, 5, -1},
			expected: []int{-6}, // Only first element is processed (-3*2=-6)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProcessData(tt.data)

			if len(result) != len(tt.expected) {
				t.Errorf("ProcessData() returned %d elements, want %d", len(result), len(tt.expected))
				return
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("ProcessData()[%d] = %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestMonitorChannel(t *testing.T) {
	t.Run("receives values before timeout", func(t *testing.T) {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		result := MonitorChannel(ch, 100*time.Millisecond)
		expected := []int{1, 2, 3}

		if len(result) != len(expected) {
			t.Errorf("MonitorChannel() returned %d values, want %d", len(result), len(expected))
			return
		}

		for i, v := range result {
			if v != expected[i] {
				t.Errorf("MonitorChannel()[%d] = %d, want %d", i, v, expected[i])
			}
		}
	})

	t.Run("timeout with no values", func(t *testing.T) {
		ch := make(chan int)

		start := time.Now()
		result := MonitorChannel(ch, 50*time.Millisecond)
		duration := time.Since(start)

		if len(result) != 0 {
			t.Errorf("MonitorChannel() returned %d values, want 0", len(result))
		}

		if duration < 45*time.Millisecond {
			t.Errorf("MonitorChannel() returned too quickly: %v", duration)
		}
	})

	t.Run("timeout with some values", func(t *testing.T) {
		ch := make(chan int, 2)
		ch <- 10
		ch <- 20

		result := MonitorChannel(ch, 50*time.Millisecond)

		if len(result) < 2 {
			t.Errorf("MonitorChannel() returned %d values, want at least 2", len(result))
		}

		if len(result) >= 2 {
			if result[0] != 10 || result[1] != 20 {
				t.Errorf("MonitorChannel() = %v, want [10, 20, ...]", result)
			}
		}
	})
}

func TestDeadlockExample(t *testing.T) {
	t.Run("completes without hanging", func(_ *testing.T) {
		start := time.Now()
		DeadlockExample()
		duration := time.Since(start)

		// Should take approximately 100ms due to the sleep
		if duration < 90*time.Millisecond || duration > 200*time.Millisecond {
			t.Errorf("DeadlockExample() took %v, expected around 100ms", duration)
		}
	})

	t.Run("demonstrates deadlock scenario", func(_ *testing.T) {
		// Test that the function completes and doesn't actually deadlock
		DeadlockExample()
		// If we reach here, the function completed successfully
	})
}

// Benchmark tests
func BenchmarkConcurrentCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentCounter(10)
	}
}

func BenchmarkProcessData(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		ProcessData(data)
	}
}

func BenchmarkMonitorChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 5)
		for j := 0; j < 5; j++ {
			ch <- j
		}
		close(ch)
		MonitorChannel(ch, 10*time.Millisecond)
	}
}
