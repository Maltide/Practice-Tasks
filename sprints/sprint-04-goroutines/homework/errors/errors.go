// Package homework содержит примеры для тренировки конкурентности и проверки ошибок.
package homework

import (
	"fmt"
	"sync"
	"time"
)

// PrintNumbers запускает несколько goroutine и печатает числа с синхронизацией через WaitGroup.
func PrintNumbers() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()

	fmt.Println("Done")
}

// PrintSquares печатает квадраты чисел из среза numbers (название переменной примерное — можно иное).
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

// SumChannel суммирует значения, полученные из канала, и возвращает сумму.
func SumChannel(numbers []int) int {
	ch := make(chan int, len(numbers))

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			ch <- num
		}
	}()

	wg.Wait()

	close(ch)

	sum := 0
	for num := range ch {
		sum += num
	}

	return sum
}

// ConcurrentCounter увеличивает счётчик конкурентно n раз и возвращает итоговое значение.
func ConcurrentCounter(n int) int {
	counter := 0

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer wg.Done()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	return counter
}

// ProcessData обрабатывает входные данные и возвращает результат обработки.
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

// MonitorChannel наблюдает за каналом input в течение указанной длительности и возвращает метрики.
func MonitorChannel(input <-chan int, duration time.Duration) []int {
	results := []int{}

	timeout := time.After(duration)

	for {
		select {
		case val, ok := <-input:
			if !ok {
				return results
			}
			results = append(results, val)
		case <-timeout:
			return results
		}
	}
}

// DeadlockExample демонстрирует ситуацию взаимной блокировки для учебных целей.
func DeadlockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		<-ch2
		ch1 <- 1
	}()

	go func() {
		<-ch1
		ch2 <- 2
	}()

	time.Sleep(100 * time.Millisecond)
}
