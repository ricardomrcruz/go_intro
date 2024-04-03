package main

import (
	"fmt"
	"math"
)

//USING FUNCTIONS


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

//area of a circle
func circleArea(r float64) float64{
	
	return math.Pi * r * r
}
 
func main7() {

	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	//this function calls a function for every strin iterated in the first arg
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
	




}
