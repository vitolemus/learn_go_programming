package main

import "fmt"

func main() {

	fruit := "apple"

	if fruit == "cambur" {
		fmt.Println("You can eat a", fruit)
	} else {
		fmt.Println("You can't eat any fruit... Sorry")
	}

	number := 18

	switch {
	case number > 18:
		fmt.Println("You are an adult")
	case number < 18:
		fmt.Println("You are a kid")
	case number == 18:
		fmt.Println("You are of age")

	}

	switch vocal := "O"; vocal {
	case "A":
		fmt.Println("your vocal is \"A\"")
	case "E":
		fmt.Println("your vocal is \"E\"")
	case "I":
		fmt.Println("your vocal is \"I\"")
	case "O":
		fmt.Println("your vocal is \"O\"")
	case "U":
		fmt.Println("your vocal is \"U\"")

	}

}
