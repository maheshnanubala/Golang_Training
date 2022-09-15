package main

import (
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
		fmt.Println("Result:", add(first, second))
	} else if operation == "diff" {
		fmt.Println("Result:", diff(first, second))
	} else if operation == "multiply" {
		fmt.Println("Result:", multiply(first, second))
	} else if operation == "div" {
		fmt.Println("Result:", div(first, second))
	} else {
		fmt.Println("invalid")
	}

}
