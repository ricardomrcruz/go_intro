package main

import (
	"fmt"
	//  "strings"
	"sort"
)

//THE STANDARD LIBRARY


func main4() {

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting,"hello!"))
	// fmt.Println(strings.ReplaceAll(greeting,"hello", "hola"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "re")) //position coutin from 0
	// fmt.Println(strings.Split(greeting, " "))
	
	//the original value remains unchanged
	// fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 36)
	fmt.Println(index)

	names := []string{"yoshiu", "mario", "bowser", "peach", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
