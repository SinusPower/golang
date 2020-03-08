package main

import (
	"fmt"
	"os"

	hw02 "github.com/sinuspower/golang/test/hw02_unpack_string"
)

const errorExitCode int = 1

func main() {
	tst := `0m`
	if res, err := hw02.Unpack(tst); err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	} else {
		fmt.Println(res)
	}
}
