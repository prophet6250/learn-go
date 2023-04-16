package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	// fmt.Println("hello world")
	// values()
	// variables()
	// loops()
	// conditions()
	// calculator()
	// arrays()
	maps()
}

func values() {
	// string concatenation works out of the box :)
	fmt.Println("hello" + "world")
	// comma separated values are concatenated in string
	fmt.Println("1+1 =", 68+1)
	// automatic type inference. 7/3 is int, 7.0/3.0 is float or double
	fmt.Println("7/3", 7/3)
	// automatic type inference. 7/3 is int, 7/3.0 is float or double
	fmt.Println("7/3", 7/3.0)
	// boolean operations work just like C-like languages
	fmt.Println(true || false)
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

	// basic ass constant variable printing
	fmt.Println(constantString)

	// can't convert scientific notation containing negative exponents to string
	const chargeOfElectron float64 = 1.602176634 / 1e19
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
	// implicit type conversion
	if 69 == 69.0 {
		fmt.Println("joe mama gay")
	} else {
		fmt.Println("who says I'm gay?\nyou are gay")
	}

	i := 10
	// this will check for values of i, won't allow conditions
	switch i {
	case 0:
		fmt.Println("this is 0")
	case 5:
		fmt.Println("this is 5")
	case 10:
		fmt.Println("this is 10")
	}
	// switch without variable will allow conditions over any variable as well
	switch {
	case i < 5:
		fmt.Println("this is < 5")
	case i < 10:
		fmt.Println("this is < 10")
	case i < 100:
		fmt.Println("this is < 100")
	}
}

func calculator() {
	// just like java, get a reader, then read provide a source
	reader := bufio.NewReader(os.Stdin)
	// I don't want any buffer overflow attacks, F U
	byteOperator, _ := reader.ReadByte()
	// will convert the ASCII representation of char into string
	stringOperator := string(byteOperator)

	var operand1, operand2, result float32
	var err error

	// C-like input scanners available, check io and bufio as well
	fmt.Scanf("%f%f", &operand1, &operand2)

	switch stringOperator {
	case "+":
		result, err = add(operand1, operand2)
	case "-":
		result, err = subtract(operand1, operand2)
	case "*":
		result, err = multiply(operand1, operand2)
	case "/":
		result, err = divide(operand1, operand2)
	default:
		fmt.Println("ERROR: Invalid Operand")
	}

	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		// string formatters work just like C, nice
		fmt.Printf("%f %s %f = %f\n", operand1, stringOperator, operand2, result)
	}
}

func add(a float32, b float32) (float32, error) {
	if valid, err := validateOperands(a, b); !valid {
		return 0, err
	}

	return a + b, nil
}

func subtract(a float32, b float32) (float32, error) {
	if valid, err := validateOperands(a, b); !valid {
		return 0, err
	}

	return a - b, nil
}

func multiply(a float32, b float32) (float32, error) {
	if valid, err := validateOperands(a, b); !valid {
		return 0, err
	}

	return a * b, nil
}

func divide(a float32, b float32) (float32, error) {
	if valid, err := validateOperands(a, b); !valid {
		return 0, err
	}

	divisionResult := a / b

	if math.IsInf(float64(divisionResult), 0) {
		return 0, errors.New("Can't divide by zero!")
	}

	return a / b, nil
}

func validateOperands(x float32, y float32) (bool, error) {
	// check whether a or b are infinite (+ve and -ve alike)
	// good to see infinity checkers already included in the std lib
	if math.IsInf(float64(x), 0) || math.IsInf(float64(y), 0) {
		return false, errors.New("Operands can't be infinite!")
	}

	return true, nil
}

func arrays() {
	fucks := [0]int{}
	fmt.Println("this is how many fucks I've got:", len(fucks))

	// array of size 1, of type int, with first element set to 0
	// can be declared as var increasedFucks int[1]; increasedFucks[0] = 0;
	increasedFucks := [1]int{0}

	fmt.Printf("Increased fucks to %d. Now my fucks contain: %d\n", len(increasedFucks), increasedFucks[0])
}

func maps() {
	var dynamicMap map[string]string = make(map[string]string)
	/*
		// this won't work because these memory references won't exist
		var staticMap map[string]string

		staticMap["hello"] = "world"
		staticMap["foo"] = "bar"
		staticMap["ligma"] = "ligma balls"

		fmt.Println(staticMap)
	*/

	dynamicMap["key1"] = "1"
	dynamicMap["key2"] = "2"
	dynamicMap["key3"] = "3"

	// will return zero value (which is null for string) for nonexistent keys
	valueOfKey, present := dynamicMap["key4"]

	fmt.Println(valueOfKey, present)
	fmt.Println("nonexistent values (should return null):", dynamicMap["key4"])

	delete(dynamicMap, "key1")
	fmt.Println(dynamicMap)

	// won't do anything if the key doesn't exist
	delete(dynamicMap, "key4")
	fmt.Println(dynamicMap)

	// declaration and definition in one line
	var newMap map[string]string = map[string]string{"key": "value"}
	fmt.Println(newMap)
}
