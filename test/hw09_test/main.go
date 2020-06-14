package main

import "fmt"

//go:generate echo "Test generate"
//go:generate go run .
func main() {
	fmt.Println("All works")
}
