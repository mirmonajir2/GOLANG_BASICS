package main

import "fmt"

/*
 As we all know that there are two types of datatype available Primitive and Non-Primitive

 Primitive for Golang's are int,float(32,64),bool,complex,byte,string,rune

 Non-primitive for Golang's are struct,array,map,chan,interface,slice(dynamic array same as arraylist), reflect, pointer

 N.T- Golang is a functional language.
*/
func main() {
	var num int
	num = 100
	fmt.Println(num)

	num1 := 19 //automatically declare datatype and variable and assign the value
	fmt.Println(num1)

	var num2 float64
	num2 = 2.4
	fmt.Println(num2)

	name := "Mir Monajir"
	fmt.Printf(name)
}
