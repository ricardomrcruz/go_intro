package main

import (
	"fmt"
)

//LOOPS

func main5() {

	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is: ", x)
	// 	x++
	// }

	// for i := 0; i<5; i++ {

	// 	fmt.Println("Value of i is: ", i)

	// }

	names := []string{"yoshiu", "mario", "bowser", "peach", "luigi"}

	// for i := 0; i < len(names); i++ {

	// 	fmt.Println(names[i])

	// }

	// for index, value := range names {

	// 	fmt.Printf("the value at index %v is %v \n", index, value)

	// }

	for _, value := range names {

		fmt.Printf("the value is %v \n", value)
		value = "new string"

	}

	fmt.Println(names)
}
