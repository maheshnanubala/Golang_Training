package main

import (
	"calculator_assignment/calculator"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Enter your name :")
	var name string
	fmt.Scanln(&name)

	currentTime := time.Now()
	if currentTime.Hour() < 12 {
		fmt.Println("hello " + name + " Good morning ")
	} else if currentTime.Hour() < 17 {
		fmt.Println("hello " + name + " Good afternoon ")
	} else {
		fmt.Println("hello " + name + " Good evening ")
	}
	// fmt.Println("Greet the person as according to the time :")
	// var greeting string
	// fmt.Scanln(&greeting)
	//fmt.Println("Hello Good " + greeting + " " + name)
	fmt.Println("enter the value of x ")
	var x int
	fmt.Scanln(&x)
	fmt.Println("enter the value of y ")
	var y int
	fmt.Scanln(&y)
	fmt.Println("Please select the number of that operation want to perform ")

	fmt.Println("1: Addition")
	fmt.Println("2: Subtraction")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Division")
	fmt.Print("Enter choice: ")
	var choice int = 0
	fmt.Scanf("%d", &choice)
	//var result int = 0
	switch choice {
	case 1:

		fmt.Printf("The result is :%d", calculator.Addition(x, y))
	case 2:

		fmt.Printf("The result is : %d", calculator.Subtraction(x, y))
	case 3:

		fmt.Printf("The result is : %d", calculator.Multiplication(x, y))
	case 4:

		fmt.Printf("The result is : %d", calculator.Division(x, y))
	default:
		fmt.Println("Invalid value")

	}

}
