package calculater

import "fmt"

func Div() {

	var no1 int64
	var no2 int64

	fmt.Println("Enter First Number :")

	fmt.Scanln(&no1)

	fmt.Println("Enter Second Number :")

	fmt.Scanln(&no2)
	no := no1 / no2

	fmt.Println("DIVISION OF TWO NUMBERS IS :", no)

}
