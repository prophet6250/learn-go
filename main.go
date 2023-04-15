package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello world")
	// values()
	variables()
}

func values() {
	// fmt.Println("hello" + "world") // string concatenation works out of the box :)
	fmt.Println("1+1 =", 68+1) // comma separated values are concatenated in string
	fmt.Println("7/3", 7/3)    // automatic type inference. 7/3 is int, 7.0/3.0 is float or double
	fmt.Println("7/3", 7/3.0)  // automatic type inference. 7/3 is int, 7/3.0 is float or double
	fmt.Println(true || false) // boolean operations work just like C-like languages
	fmt.Println(!false)
}

const constantString string = "constant string"

func variables() {
	var a = "some string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	another_string := "another string"
	fmt.Println(another_string)

	fmt.Println(constantString) // basic ass constant variable printing

	const chargeOfElectron float64 = 1.602176634 / 1e19 // can't convert scientific notation containing negative exponents to string
	fmt.Println(chargeOfElectron)
}
