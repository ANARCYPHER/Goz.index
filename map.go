package main

// import other packages
import (
	"fmt"
)


func main(){

	// Maps var weapon = map [key_type]value_type{...}

	var weapon = map[string]string{
		"John":"ak47",
		"Mary":"ak74",
		"Tim":"hk"}

	fmt.Println(pizza)
	fmt.Println(pizza["John"])

        var weapon = map[string]string{
		"John":"ak47",
		"Mary":"ak74",
		"Tim":"hk"}

	fmt.Println(weapon)
	fmt.Println(weapon["John"])


	// Shorthand
	// toppings := map[string]string{"John":"Pepperoni"}

	// Change an Item
	weapon["John"] = "Uzi"
	fmt.Println(weapon["John"])

	// Delete Items
	delete(weapon, "Tim")
	fmt.Println(weapon)
	
	// Iterate Over Maps
	for key, value := range weapon {
		fmt.Println(key, value)
	}	


}