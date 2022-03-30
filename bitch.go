package main


// import other packages
import (
	"fmt"
)

func main(){

	// Variables and Constants
	// string, int, float32, bool

	//var firstName string = "Cyka"

	// := shorthand variables

	//lastName := "Bitch"
	// age := 18

	// Create but don't assign

	var fullName string
	var age int
	var price float32
	var tf bool

	fmt.Println(fullName, age, price, tf)


	fullName = "Cyka Bitch"
	age = 44
	price = 19.99
	tf = true

	fmt.Println(fullName, age, price, tf)

	// Multiple variable declaration

	var name1, name2 string = "Whore", "Suck"
	fmt.Println(name1, name2)

	name1 = "Fuck"
	fmt.Println(name1, name2)

	
	// Constants
	const VULVA string = "Fish"
	fmt.Println(VULVA)

	//PIZZA = "Cheese"
	//fmt.Println(VULVA)

	// Multiple Constants

	const (
		VULVA1 = "Fish"
		VULVA2 = "Cheese"
		MYNUM = 77
	)

	fmt.Println(VULVA2)


	// fmt.Println(firstName)

}