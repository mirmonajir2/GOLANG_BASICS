package main

import (
	"errors"
	"fmt"
)

//hamming Distance

func hammingDistance(a string, b string) (int, error) {
	dis := 0
	if len(a) != len(b) {
		return 0, errors.New("must have same length")
	}
	for i := range a {
		if a[i] != b[i] {
			dis++
		}
	}
	return dis, nil
}

func main() {
	fmt.Println("Enter String 1: ")
	var s1 string
	fmt.Scanln(&s1)
	fmt.Println("Enter String 2: ")
	var s2 string
	fmt.Scanln(&s2)
	fmt.Println(hammingDistance(s1, s2))
}
