package main

import (
	"fmt"
)

//PACKAGE SCOPES (Test 9)

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {

	fmt.Println("hello", n)

}

func showScore() {
	fmt.Println("you scored these many points:", score)
}
