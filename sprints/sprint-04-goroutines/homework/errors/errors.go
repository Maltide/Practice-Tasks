package homework

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumbers() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	fmt.Println("Done")
}

func PrintSquares(numbers []int) {
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d squared is %d\n", num, num*num)
		}()
	}

	wg.Wait()
}

func SumChannel(numbers []int) int {
	ch := make(chan int)

	go func() {
		for _, num := range numbers {
			ch <- num
		}
	}()

	sum := 0
	for num := range ch {
		sum += num
	}

	return sum
}

func ConcurrentCounter(n int) int {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	return counter
}

func ProcessData(data []int) []int {
	results := make(chan int)

	go func() {
		for _, d := range data {
			results <- d * 2
		}
		close(results)
	}()

	output := []int{}
	if len(data) > 0 {
		output = append(output, <-results)
	}

	return output
}

func MonitorChannel(input <-chan int, duration time.Duration) []int {
	results := []int{}
	timeout := time.After(duration)

	for {
		select {
		case val := <-input:
			results = append(results, val)
		case <-timeout:
			return results
		}
	}
}

func DeadlockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		<-ch2
	}()

	go func() {
		ch2 <- 2
		<-ch1
	}()

	time.Sleep(100 * time.Millisecond)
}
