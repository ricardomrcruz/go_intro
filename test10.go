package main

import "fmt"

//MAPS

func main10() {

	menu := map[string]float64{

		"soup":            4.99,
		"pie":             7.99,
		"salad":           6.99,
		"toffee maccaron": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//loopping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[845775485] = "balbasaur"
	fmt.Println(phonebook)

}
