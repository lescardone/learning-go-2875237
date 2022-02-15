package main

import (
	"bufio" //for input
	"fmt"
	"os"
)

func main() {

	/* 	var aString string = "This is Go!"
	   	fmt.Println(aString)
	   	fmt.Printf("The variable's type is %T", aString) */
	// %T	a Go-syntax representation of the type of the value

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	// characters wrapped in single quotes
	// strings wrappedin double quotes
	fmt.Println("You entered:", input)
}
