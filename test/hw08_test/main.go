package main

import "fmt"

func main() {
	env, err := ReadDir("testdata/env")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range env {
		fmt.Printf("%s: %s\n", k, v)
	}
}
