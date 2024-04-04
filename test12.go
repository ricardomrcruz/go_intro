package main

import "fmt"

//POINTERS

func updateName12(x *string) { //using an asterix in a variable inside a function inside while accepting a pointer as the argument of that function
	*x = "wedge"
}

func main12() {

	name := "tifa"

	// updateName(name)

	fmt.Println("memory address of name is :", &name)

	m := &name //using & before the var name give syou the pointer (when you make a copy of the original variable)
	// fmt.Println("memory address:", m)           //here you can ndisplay the pointer
	// fmt.Println("value at memory address:", *m) //by adding a * before the variable name you can get the value the pointer redirects to

	fmt.Println(name)
	updateName12(m)
	fmt.Println(name) //pointers have a reverse logic to it, by using pointers in the function it makes the variale return to its original value.

}

/*

|---name---|----m-----|
|   0x001  |   0x002  |
|---name---|----m-----|
|  "tifa"  |  p0x001  |
|----------|----------|





*/
