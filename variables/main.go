package main

import "fmt"

func main() {
	var myVarInt int
	myVarInt = 25
	fmt.Println("This is my first Variable:", myVarInt)

	myVarStr := "Hello... Hello..!"
	fmt.Println("This is my String:", myVarStr)

	myBool := true
	fmt.Println("Go is Awesome... That's", myBool)

	fmt.Println("My variable address is", &myVarInt)
	fmt.Println("My variable address is", &myVarStr)
	fmt.Println("My variable address is", &myBool)

}
