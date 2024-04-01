package main

import (
	"fmt"
)

func main() {

	age := 25

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less tha 45")
	}

	names := []string{"yoshiu", "mario", "bowser", "peach", "luigi"}

	for index, value := range names {
		if index == 1 {
			fmt.Println(" continuing at pos", index)
			continue
		}

		fmt.Printf("the value at %v is %v", index, value)

	}

}
