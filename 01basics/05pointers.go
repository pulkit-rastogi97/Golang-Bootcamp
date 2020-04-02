package main

// import (
// 	"fmt"
// )

// func main() {

// 	//Pointer decalaration
// 	 var p *int

// 	// "eRROR"
// 	//
// 	// fmt.Println("P is having a value: ", *p)

// 	if p != nil {
// 		fmt.Println(" P is having a value : ", *p)
// 	} else {
// 		fmt.Println("P is not having any value")
// 	}

// 	// "Assigning a vlue to p"
// 	//var score int = 34
// 	 //p := &score

// 	var lifebooster float64 = 99.2
// 	// var lifeboosterRef = &lifebooster
// 	//
// 	// OR
// 	//
// 	lifeboosterRef := &lifebooster

// 	/*
// 		Here you can see a visual change,
// 		we haven't printed the lifebooster variable
// 		but we are not having an error.

// 		Because error occurs only when we not use the variable
// 		but here in this case, we are using the variable somehow
// 		because we are assigning the memory address of the variable
// 		of lifebooster to lifeboosterRef (pointer)
// 	*/

// 	// fmt.Println(lifebooster)
// 	fmt.Println(lifeboosterRef)

// 	//value stored at pointer's pointing to memory
// 	fmt.Println(*lifeboosterRef)

// 	// After changing the value of the referenced variable
// 	lifebooster = lifebooster * 2.2

// 	fmt.Println(*lifeboosterRef)

// }
