package main

// import other packages
import (
	"fmt"
)

func addIt(num1, num2 int){
	return (num1 + num2)
}

func gunFunction(gunWeapon string, bullet int){
	fmt.Println("This is my Function!")
	fmt.Println("Hi", gunWeapon, "Weapon", bullet, "ammunition")
}

func main(){

	// Functions
	myFunction("Ak47", 200)

	fmt.Println(addIt(2,3))



}