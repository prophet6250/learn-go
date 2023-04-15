package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello world")
	// values()
	// variables()
	// loops()
	conditions()
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

func loops() {
	// basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop equivalent
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
}

func conditions() {
	if 69 == 69.0 { // implicit type conversion
		fmt.Println("joe mama gay")
	} else {
		fmt.Println("who says I'm gay?\nyou are gay")
	}

	i := 10
	switch i { // this will check for values of i, won't allow conditions
	case 0:
		fmt.Println("this is 0")
	case 5:
		fmt.Println("this is 5")
	case 10:
		fmt.Println("this is 10")
	}

	switch { // switch without variable will allow conditions over any variable as well
	case i < 5:
		fmt.Println("this is < 5")
	case i < 10:
		fmt.Println("this is < 10")
	case i < 100:
		fmt.Println("this is < 100")
	}
}
