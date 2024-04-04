package main

import "fmt"

// Structs & Custom Types
//Receiver Functions

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {

	fs := "Bill breakdown:  \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-20v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-20v ...$%0.2f", "total:", total)

	return fs

}

// ouput
// $ go run  main.go bill.go
// Bill breakdown:
// pie:                 ...$5.99
// cake:                ...$3.99
// total:               ...$9.98

//update tip

// func (b bill) updateTip(tip float64){
// 	b.tip = tip
// }

// func (b bill) addItem(name string, price )
