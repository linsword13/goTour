package main

import (
	"fmt"
	"time"
)

func main() {
	i := []int{1, 2}
	go func() { i[0] = 3 }()
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", i)
}
