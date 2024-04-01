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

func main() {

	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

}
