package main

import (
	"fmt"
	"go_workspace/calculater"
	"time"
)

func main() {

	fmt.Println("Enter Your Name :")

	var name string

	fmt.Scanln(&name)

	curntTime := time.Now()
	fmt.Println(curntTime)

	if curntTime.Hour() < 12 {
		fmt.Printf(" Hello Good <Morning>, <%s>!", name)
	} else if curntTime.Hour() < 18 {
		fmt.Printf(" Hello Good <Afternoon>, <%s>!", name)
	} else if curntTime.Hour() < 24 {
		fmt.Printf(" Hello Good <Evening>, <%s>!", name)
	} else {
		fmt.Printf(" Welcome , <%s>!", name)
	}

	fmt.Println()

	fmt.Println("Calculater : Select The Number to process the Operation of TWO Numbers:")
	fmt.Println("1. ADD")
	fmt.Println("2. SUB")
	fmt.Println("3. MUL")
	fmt.Println("4. DIV")

	var no int

	fmt.Scanln(&no)

	if 1 == no {
		calculater.Add()
	} else if 2 == no {
		calculater.Sub()
	} else if 3 == no {
		calculater.Mul()
	} else if 4 == no {
		calculater.Div()
	} else {
		fmt.Println("You Enter Invalid Number, Which is not present in above option's")
	}

}
