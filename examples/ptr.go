package main

import "fmt"

func foo() *int {
	i := 1
	return &i
}

func main() {
	fmt.Printf("%v\n", *foo())
}
