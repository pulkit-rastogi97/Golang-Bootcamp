package main

import (
	"fmt"
)

func main() {
	superman()
	result := multiplyMe(3, 6)
	fmt.Println(result)
	newResult := newMultiplyMe(7, 7)
	fmt.Println(newResult)
	
	sum, myLength, myName := adder(3, 6, 7, 8, 23, 12)
	fmt.Println(sum, myLength, myName)

}

func superman() {
	fmt.Println("I am superman")
}

func multiplyMe(v1 int, v2 int) int {
	return v1 * v2
	// will throw an error missing return statement
	// fmt.Println("Never goin to execute")
}

// will also work
func newMultiplyMe(v1, v2 int) int {
	return v1 * v2
}

// no function overloading in golang

//only specific type of multiple values
// can return multiples values from a single function
func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just For Fun"
	return sum, length, name
}
