package main

import (
	"fmt"
	"strings"
)

func main() {

	greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting,"hello!"))
	// fmt.Println(strings.ReplaceAll(greeting,"hello", "hola"))
	// fmt.Println(strings.ToUpper(greeting))
	
	//the original value remains unchanged
	fmt.Println("original string value =", greeting)
}
