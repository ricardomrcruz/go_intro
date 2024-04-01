package main

import (
	"fmt"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleNames(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}
}

func main() {

	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	//this function calls a function for every strin iterated in the first arg 
	cycleNames([]string{"cloud","tifa", "barret"}, sayGreeting)
}
