package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func leibnizLine(id, n int, done <-chan struct{}, wg *sync.WaitGroup, resCh chan<- float64) {
	defer wg.Done()

	var sum float64
	for i := id; ; i += n {
		select {
		case <-done:
			resCh <- sum
			return
		default:
			term := math.Pow(-1, float64(i)) / float64(2*i+1)
			sum += term
		}
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "Число горутин для вычисления ряда:")
	flag.Parse()

	if n < 1 {
		fmt.Println("Число горутин должно быть не меньше 1")
		os.Exit(1)
	}

	done := make(chan struct{})
	resCh := make(chan float64, n)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go leibnizLine(i, n, done, &wg, resCh)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Вычисление начинается... Для остановки нажмите Ctrl+C")

	select {
	case sig := <-sigChan:
		fmt.Printf("\nПолучен сигнал %s, завершаем вычисления...\n", sig)
		close(done)
	}

	wg.Wait()
	close(resCh)

	var totalSum float64
	for part := range resCh {
		totalSum += part
	}

	piApprox := totalSum * 4
	fmt.Printf("Приближенное значение числа π: %.15f\n", piApprox)
	fmt.Println("Программа завершена.")

	time.Sleep(100 * time.Millisecond)
}
