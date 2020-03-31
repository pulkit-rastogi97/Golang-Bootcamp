package main

import (
	"fmt"
)

func main() {

	var isLoggedIn bool = true
	var balance int = 10

	// "After saving () of if will go"
	//
	// if (isLoggedIn && balance > 15) {
	// 	fmt.Println("Show cart page")
	// } else {
	// 	fmt.Println("Show user another page")
	// }

	if isLoggedIn && balance > 5 {
		fmt.Println("Show cart page")
		var len, err = fmt.Println("Show cart page")
		if err == nil {
			fmt.Printf("length is %v, %T ", len, len)
			fmt.Println(`I am a backtick`)
		}

	} else {
		fmt.Println("Show user another page")
	}

}
