package main

import (
	"fmt"
	"os"
	"strings"
)

func createBill() bill {

	reader := bufio.newReader(os.Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

}

func main() {

	mybill := createBill()

}
