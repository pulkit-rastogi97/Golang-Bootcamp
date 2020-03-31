package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var name string
	var userRating string

	//Front End

	//Take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter your full name")
	name, _ = reader.ReadString('\n')

	//Take rating from user and convert
	// it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Pizza center between 1 and 5 :: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back end
	// fmt.Printf("Hello %v , \n Thanks for rating our Pizza center with %v star rating. \n\n Your rating was recorded in our system at %v ",name, myNumRating, time.Now().Format(time.Kitchen))

	//time.Stamp
	fmt.Printf("Hello %v , \n Thanks for rating our Pizza center with %v star rating. \n\n Your rating was recorded in our system at %v \n\n", name, myNumRating, time.Now().Format(time.Stamp))

	if myNumRating == 5 {
		fmt.Println("Bonus for tema for 5 star service")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("We are always improving")
	} else if myNumRating < 3 {
		fmt.Println("Need serious work on our side")
	}
}
