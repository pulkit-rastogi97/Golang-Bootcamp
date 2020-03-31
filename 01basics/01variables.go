package main

import "fmt"

func main() {

	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3.0
	fmt.Printf("%v, %T \n", thorRating, thorRating)

	thorRating1 := 3
	fmt.Printf("%v, %T \n", thorRating1, thorRating1)

	thorRating2 := 3. //Bug could be there as 3. is also float
	fmt.Printf("%v, %T \n", thorRating2, thorRating2)

	var Ironman, CaptAmerica string = "I am Ironman", "I am Capt. America"
	fmt.Println(Ironman, CaptAmerica)

	var defaultValue int
	fmt.Println(defaultValue)

	// var defaultValue string   //Empty line show null value
	// fmt.Println(defaultValue)

	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "I am antman"
	)
	fmt.Println(spiderman, age, powers, antman)

}
