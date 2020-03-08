package main

import (
	"fmt"
	"os"

	hw02 "github.com/sinuspower/golang/test/hw02_unpack_string"
)

const errorExitCode int = 1

func main() {
	tst := `a4b\55c2d5e0`
	if res, err := hw02.Unpack(tst); err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	} else {
		fmt.Println(res)
	}
}
