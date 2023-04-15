package main

import "fmt"

func main() {
	// fmt.Println("hello world")
	values()
}

func values() {
	// fmt.Println("hello" + "world") // string concatenation works out of the box :)
	fmt.Println("1+1 =", 68+1) // comma separated values are concatenated in string
	fmt.Println("7/3", 7/3)    // automatic type inference. 7/3 is int, 7.0/3.0 is float or double
	fmt.Println("7/3", 7/3.0)  // automatic type inference. 7/3 is int, 7/3.0 is float or double
	fmt.Println(true || false) // boolean operations work just like C-like languages
	fmt.Println(!false)
}
