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
	out := make(Out)
	select {
	case <-done:
		return out
	default:
		out = in
		for i, s := range stages {
			fmt.Println("stage ", i)
			out = s(out)
		}
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
		//result = append(result, s.(string))
		fmt.Println(s)
	}

	for _, r := range result {
		fmt.Println(r)
	}

}
