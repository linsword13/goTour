package main

import "C"

//export GoSquare
func GoSquare(n int) int {
	return n * n
}

func main() {}
