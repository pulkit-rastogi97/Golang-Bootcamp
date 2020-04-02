package main

// import (
// 	"fmt"
// )

// func main() {

// 	/*

// 		make v/s new

// 		new - only allocates - no init of memory
// 		make - allocate and init - non zeroed storage

// 	*/

// 	/*

// 		// keys as string and values as int
// 		// inbuilt new keyword is used
// 		// error will throw
// 		var score map[string] int
// 		score["pulkit"] = 89
// 		fmt.Println(score)

// 	*/

// 	score := make(map[string]int)
// 	score["pulkit"] = 90
// 	fmt.Println(score)

// 	score["bhumi"] = 99
// 	score["mridul"] = 79
// 	score["gagan"] = 86
// 	score["shray"] = 35
// 	fmt.Println(score)

// 	getMridulScore := score["mridul"]
// 	fmt.Println(getMridulScore)

// 	delete(score, "shray")
// 	fmt.Println(score)

// 	//for loop k-key and v-value are any variable name
// 	// sorted alphabetically
// 	for k, v := range score {
// 		fmt.Printf("Score of %v  is  %v \n", k, v)
// 	}

// }
