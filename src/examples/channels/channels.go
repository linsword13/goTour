package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2}
	done := make(chan struct{})
	go func() {
		i[0] = 3
		done <- struct{}{}
	}()
	<-done
	fmt.Printf("%v\n", i)
}
