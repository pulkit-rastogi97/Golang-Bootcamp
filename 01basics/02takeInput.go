package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var myString string
	fmt.Scanln(&myString)
	fmt.Println(myString)


	//"When we don't want to use the second value"
	
	var name string = "pulkit"
	var a, _ = fmt.Println(name)
	fmt.Println(a)


	//"How to take user input"
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name :: ")
	myName, _ := reader.ReadString('\n')
	fmt.Println(myName)


	//"Error of not being able to use 2 as type string"
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating :: ")
	myRating, _ := reader.ReadString('\n')
	fmt.Println(myRating + 2)


	//'cannot use error value as type int in argument to strconv.ParseFloat'
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating :: ")
	myRating, _ := strconv.ParseFloat(reader.ReadString('\n'))
	fmt.Println(myRating + 2)


	//"How to take input and converting it to other int or float form"
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating :: ")
	myRating, _ := reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(myNumRating + 2)


	
}
