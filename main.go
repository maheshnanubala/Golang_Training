package main

import (
	"fmt"
	"time"

	"deo.go/calculator"
)

func main() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)

	// time func Using switch **************

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")

	}

	var a, b int
	var operation string
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&b)
	fmt.Println("select operation:")
	fmt.Scan(&operation)

	// operation func also Using switch **********

	switch operation {
	case "Additon":
		fmt.Printf("Additon is: %d \n", calculator.Additon(a, b))
	case "Subtraction":
		fmt.Printf("Subtraction is: %d \n", calculator.Subtraction(a, b))
	case "Multiplying":
		fmt.Printf("Multiplying is: %d \n", calculator.Multiplying(a, b))
	case "Division":
		fmt.Printf("Division is: %d \n", calculator.Division(a, b))
	default:
		fmt.Println("Invalid Operation")
	}

}
