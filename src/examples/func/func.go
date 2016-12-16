package main

import "fmt"

func age(initAge int) func() {
	i := initAge
	return func() {
		i++
		fmt.Printf("%v years old\n", i)
	}
}

func main() {
	celBirthday := age(18)
	for i := 0; i < 5; i++ {
		celBirthday()
	}
}
