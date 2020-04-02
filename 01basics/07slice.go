package main

// import(
// 	"fmt"
// 	"sort"
// )


// func main(){
	

// 	// no. will be array
// 	//without no. will be slice
// 	var things = [] string { "maleta", "ropa", "reloj"}
// 	fmt.Println(things)


// 	things = append(things, "bolso")
// 	fmt.Println(things)

// 	// remove one element from left
// 	// things = append(things[1:])
// 	// fmt.Println(things)


// 	/*

// 	// default
// 	things = append(things[1:len(things)])
// 	fmt.Println(things)

// 	*/


// 	// remove one element from both ends
// 	things = append(things[1:len(things)-1])
// 	fmt.Println(things)

// 	/*

// 	//Index must be non-negative

// 	things = append(things[1:-1])
// 	fmt.Println(things)

// 	*/


// 	// slice creation
// 	heros := make([]string, 3, 3)
// 	heros[0] = "thor"
// 	heros[1] = "ironman"
// 	heros[2] = "spiderman"
// 	fmt.Println(heros)

// 	//Can increase the size on runtime. reserves more memory
// 	heros = append(heros, "deadpool")
// 	fmt.Println(heros)
	
// 	//Cap allowed : doubled the size after inserting deadpool
// 	fmt.Println(cap(heros))

// 	myInts := [] int{4, 7, 3, 2, 8}

// 	isSorted := sort.IntsAreSorted(myInts)
// 	fmt.Println("Are ints sorted :: ", isSorted)


// 	// Will sort the slice of integers in increasing order
// 	sort.Ints(myInts)
// 	fmt.Println(myInts)


// }