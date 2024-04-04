package main

import "fmt"

//PASS BY VALUE

func updateName11(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main11() {
	//group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// updateName(name)
	//when passing a function for the second time, golang allocates a new block of memory thats why this function prints tifa and not wedge.
	//because its a copy, it only changed the value of the variable in the copy.

	name = updateName11(name)

	fmt.Println(name)

	//group B types -> slices, functions, maps
	menu := map[string]float64{

		"icecream": 4.99,
		"pie":      7.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
