package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	//the length of the arrays is fixed
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshiu", "mario", "bowser", "peach"}
	names[1] = "luigi" //to change the value of a element of an array

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores)

}
