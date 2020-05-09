package main

import (
	"fmt"
	"strconv"
	"time"
)

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

const (
	sleepPerStage = time.Millisecond * 100
	//fault         = sleepPerStage / 2
)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	multiplex := func(in In, out Bi) {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}

	out := in
	for _, stage := range stages {
		stageIn := make(Bi)
		go multiplex(out, stageIn)
		out = stage(stageIn)
	}

	return out
}

func main() {
	// Stage generator
	g := func(name string, f func(v I) I) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g("Dummy", func(v I) I { return v }),
		g("Multiplier (* 2)", func(v I) I { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v I) I { return v.(int) + 100 }),
		g("Stringifier", func(v I) I { return strconv.Itoa(v.(int)) }),
	}

	/*
		done := make(Bi)
		abortDur := sleepPerStage
		go func() {
			<-time.After(abortDur)
			close(done)
		}()
	*/

	in := make(Bi)
	data := []int{1, 2, 3, 4, 5}
	go func() {
		for _, v := range data {
			in <- v
		}
		close(in)
	}()

	result := make([]string, 0, 10)
	for s := range ExecutePipeline(in, nil, stages...) {
		result = append(result, s.(string))
	}

	for _, r := range result {
		fmt.Println(r)
	}

}
