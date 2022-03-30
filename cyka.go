package main

// import other packages
import (
	"fmt"
)

func main(){

	// Arrays
	var firstNames = [3] string {"Suck", "Fuck", "Pussy"}
	fmt.Println(firstNames[1])

	// Infer length
	var nummies = [] int {1,2,3,4,5}
	fmt.Println(nummies[0])

	// shorthand
	lastNames := [2] string {"Vulva", "Pussy"}
	fmt.Println(lastNames[0])

	// shorthand with inferred
	ourNames := [] string {"Vulva", "Ass"}
	fmt.Println(ourNames)

	// Change item in array
	ourNames[1] = "Cum"
	fmt.Println(ourNames)

	// Default Assignment 
	var ourNummies = [5] int {1,2}
	fmt.Println(ourNummies)

	// Assign in certain positions
	var moreNummies = [10] int {0:41, 4:99}
	fmt.Println(moreNummies)

	// Slices
	var toppings = [5] string {"Vulva", "Pussy", "Ass", "Whore", "Butt"}
	fmt.Println(toppings)

	toppingSlice := toppings[0:2]
	fmt.Println(toppingSlice)

	// modify a slice
	toppingSlice[1] = "Pussy"
	fmt.Println(toppingSlice)

	// Add item to slice
	toppingSlice = append(toppingSlice, "Cum")
	fmt.Println(toppingSlice)

	// add two slices together
	otherSlice := toppings[3:4]
	fmt.Println(otherSlice)

	otherSlice = append(otherSlice, toppingSlice...)
	fmt.Println(otherSlice)


}