package main

import "fmt"

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("the type of the variable is : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of the variable is : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("the type of the variable is : %T \n", smallVal)

	var smallFloat float32 = 245.3456
	fmt.Println(smallFloat)
	fmt.Printf("the type of the variable is : %T \n", smallFloat)

}
