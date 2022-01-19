package main

import "fmt"

func main() {
	//if-else
	var input int
	fmt.Println("Enter The Number: ")
	fmt.Scanln(&input) //to take input from command line

	if input%2 == 0 {
		fmt.Println("The Number is even", input)
	} else {
		fmt.Println("The Number is odd", input)
	}

	if x := 8; x > 9 {
		fmt.Println("The Number is greater than 9")
	} else {
		fmt.Println("The number is less than 10")
	}
	// switch case
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
