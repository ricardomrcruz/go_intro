package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// User Input lesson18 youtube.com/watch?v=20HlPtQuc_g

//the package bufio gives the ability to read what the user types in the console

func getInput1(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill1() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	c := newBill(name)
	fmt.Println("Created the bill - ", c.name)

	return c

}

// commented for error solving purposes
// func promptOptions1(c bill) {

// 	reader := bufio.NewReader(os.Stdin)
// 	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
// 	fmt.Println(opt)

// }

func main14() {

	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)

}
