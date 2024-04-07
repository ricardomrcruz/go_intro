package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SAVING FILES https://www.youtube.com/watch?v=J88pG3NeRoA

//see doc for os.WriteFile

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64) //checks the data type, if there isnt a float64 type present sends error and resets function
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "s":

		b.save()

		fmt.Println("bill saved succesfully for :", b.name)

	case "t":
		tip, _ := getInput("Enter tip amount (€): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b) //it re executes the function
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)

}

//OUTPUT

// go run  main.go bill.go
// Create a new bill name:
// ricardo
// Created the bill -  ricardo
// Choose option (a - add item, s - save bill, t - add tip):
// a
// Item name:
// spaghetti
// Item price:
// 13.5
// item added -  spaghetti 13.5
// Choose option (a - add item, s - save bill, t - add tip):
// a
// Item name:
// martini
// Item price:
// 4
// item added -  martini 4
// Choose option (a - add item, s - save bill, t - add tip):
// t
// Enter tip amount (€):
// 5
// tip added -  5
// Choose option (a - add item, s - save bill, t - add tip):
// s
// bill saved {ricardo map[martini:4 spaghetti:13.5] 5}
