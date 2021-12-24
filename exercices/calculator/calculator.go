package calculator

import "fmt"

func Calculator() {

	fmt.Println("Welcome to calculator")
	fmt.Println("********************MAIN MENU*************************")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("******************************************************")
	var choice int

	fmt.Scan(&choice)
	var a, b int

	fmt.Println("Enter value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter value of b: ")
	fmt.Scan(&b)
	if choice == 1 {
		// choice 1 activates this part --> addition
		ans := a + b
		fmt.Println("Answer = ", ans)
	} else if choice == 2 {
		// choice 2 activates this part --> subtraction
		ans := a - b
		fmt.Println("Answer = ", ans)
	} else if choice == 3 {
		// choice 3 activates this part --> multiplication
		ans := a * b
		fmt.Println("Answer = ", ans)
	} else {
		// choice 4 activates this part --> division
		// remember not to enter second value as 0
		// as that would raise a DivideByZero error
		// or may display infinity
		ans := a / b
		fmt.Println("Answer = ", ans)
	}
	fmt.Println("******************************************************")
	fmt.Println("Thank you for using calculator! Have a nice day ahead. ^-^")
	fmt.Println("******************************************************")
}
