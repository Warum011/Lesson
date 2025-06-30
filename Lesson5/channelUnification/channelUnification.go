package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(wg *sync.WaitGroup, c <-chan interface{}) {
			defer wg.Done()
			select {
			case <-c:
				cancel()
			case <-ctx.Done():
			}
		}(&wg, ch)
	}
	wg.Wait()
	close(out)

	return out
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
