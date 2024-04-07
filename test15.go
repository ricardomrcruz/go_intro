package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SWITCH CASES LESSON19

func getInput15(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill15() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func promptOptions15(b bill) {

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "s":
		fmt.Println("you saved the bill")
	case "t":
		tip, _ := getInput("Enter tip amount (€): ", reader)

		fmt.Println(tip)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b) //it re executes the function
	}
}

func main15() {

	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)

}
