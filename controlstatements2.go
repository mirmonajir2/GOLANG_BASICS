package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
	    for infinite execution
	   for true{
	   		fmt.Println()
	   	}
	*/
	// implement do-while loop using for
	j := 0
	for {
		fmt.Println(j)
		if j == 0 {
			break
		}
	}
}
