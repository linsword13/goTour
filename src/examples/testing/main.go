package main

import (
	"./helper"
	"fmt"
)

func main() {
	int := helper.ExportedFunc()
	fmt.Println("we get", int, helper.ExportedVar)
}
