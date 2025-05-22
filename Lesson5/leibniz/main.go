package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func leibnizLine(ctx context.Context, id int, n int, wg *sync.WaitGroup, resCh chan<- float64) {
	defer wg.Done()
	var sum float64

	for i := id; ; i += n {
		select {
		case <-ctx.Done():
			resCh <- sum
			return
		default:
			term := math.Pow(-1, float64(i)) / float64(2*i+1)
			sum += term
		}
	}
}

func main() {

	n := flag.Int("input", 1, "number of goroutines:")
	flag.Parse()

	if *n < 1 {
		fmt.Println("the number of goroutines cannot be less than 1")
		return
	}

	resCh := make(chan float64, *n)
	var wg sync.WaitGroup
	wg.Add(*n)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < *n; i++ {
		go leibnizLine(ctx, i, *n, &wg, resCh)

	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("counting ^_^, To stop ctrl+c")

	select {
	case <-sigChan:
		cancel()
	}

	wg.Wait()
	close(resCh)

	var totalSum float64
	for step := range resCh {
		totalSum += step
	}

	result := totalSum * 4
	fmt.Printf("the calculated value of Pi is: %.10f", result)

}
