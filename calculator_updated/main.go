package main

import (
	"calculator_updated/calculator"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Enter your Name:")
	var name string
	fmt.Scanln(&name)
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	//time := time.Now()
	if currentTime.Hour() < 12 {
		fmt.Println("Good morning", name)
	} else if currentTime.Hour() < 18 {
		fmt.Println("Good afternoon", name)
	} else {
		fmt.Println("Good evening", name)
	}

	fmt.Println("Enter first number:")
	var first int8
	fmt.Scanln(&first)
	fmt.Println("Enter second number:")
	var second int8
	fmt.Scanln(&second)
	fmt.Println("select operation:")
	var operation string
	fmt.Scanln(&operation)
	if operation == "add" {
		fmt.Println("Result:", calculator.Add(first, second))
	} else if operation == "diff" {
		fmt.Println("Result:", calculator.Diff(first, second))
	} else if operation == "multiply" {
		fmt.Println("Result:", calculator.Multiply(first, second))
	} else if operation == "div" {
		fmt.Println("Result:", calculator.Div(first, second))
	} else {
		fmt.Println("invalid")
	}

}
