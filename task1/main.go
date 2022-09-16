package main

import (
	"fmt"
	"task/calculator"
	"time"
)

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your  Name: ")

	// var then variable name then variable type
	var name string

	// Taking input from user
	fmt.Scanln(&name)
	currenttime := time.Now()
	//fmt.Println(currenttime)
	if currenttime.Hour() < 12 {

		fmt.Println("Good Morning", name)
	} else if currenttime.Hour() < 18 {

		fmt.Println("Good Afternoon", name)
	} else if currenttime.Hour() < 21 {

		fmt.Println("Good Evening", name)
	}
	var operation string
	var number1, number2 int

	fmt.Println("Select Operation:")

	fmt.Println("1.Add")
	fmt.Println("2.Sub")
	fmt.Println("3.Mul")
	fmt.Println("4.Div")
	fmt.Scan(&operation)

	fmt.Print("Please enter First number: ")
	fmt.Scan(&number1)
	fmt.Print("Please enter Second number: ")
	fmt.Scan(&number2)

	switch operation {
	case "add":

		fmt.Println("output:", calculator.Add(number1, number2))

	case "sub":
		fmt.Println("output:", calculator.Sub(number1, number2))
	case "mul":
		fmt.Println("output:", calculator.Mul(number1, number2))
	case "div":
		fmt.Println("output:", calculator.Div(number1, number2))

	default:
		fmt.Println("Invalid Operation")
	}

}
