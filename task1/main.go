package main

import (
	"fmt"
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

	var output = 0
	switch operation {
	case "1":
		output = number1 + number2

	case "2":
		output = number1 - number2
	case "3":
		output = number1 * number2
	case "4":
		output = number1 / number2

	default:
		fmt.Println("Invalid Operation")
	}
	fmt.Println("The result is:", output)

}
