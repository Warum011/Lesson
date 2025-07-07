package main

import (
	"fmt"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var inChan In = in

	for _, stage := range stages {
		thisStageIn := inChan
		thisStageOut := make(Bi)

		go func(stage Stage, in In, out Bi) {
			defer close(out)

			defer func() {
				if r := recover(); r != nil {
				}
			}()

			stageOut := stage(in)

			for {
				select {
				case <-done:
					return
				case v, ok := <-stageOut:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case out <- v:
					}
				}
			}
		}(stage, thisStageIn, thisStageOut)

		inChan = thisStageOut
	}

	return inChan
}

func main() {
	source := make(Bi)
	done := make(Bi)

	go func() {
		defer close(source)
		for i := 1; i <= 5; i++ {
			source <- i
		}
	}()

	stage1 := func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for v := range in {
				out <- v.(int) + 1
			}
		}()
		return out
	}

	stage2 := func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for v := range in {
				out <- v.(int) * 2
			}
		}()
		return out
	}

	stage3 := func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for v := range in {
				time.Sleep(100 * time.Millisecond)
				out <- v
			}
		}()
		return out
	}

	res := ExecutePipeline(source, done, stage1, stage2, stage3)

	for v := range res {
		fmt.Println(v)
	}
}
