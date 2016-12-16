package helper

import "fmt"

var ExportedVar = 1
var secretVar = 2

func ExportedFunc() int {
	fmt.Println("Exported")
	return secretVar
}

func secretFunc() {
	fmt.Println("You can't see me!")
}
