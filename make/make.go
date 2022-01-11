package main

import "fmt"

func main() {
	//Creando un Slice vacío
	letras := make([]string, 0, 3)

	fmt.Println(letras, len(letras), cap(letras)) // imprime slice "letras", longitud y capacidad

	/*
	   AL DECLARAR UN SLICE VACÍO O PERMITE ACCEDER A LOS INDICES
	   	letras[0] = a
	   	letras[1] = b
	   	letras[2] = c
	*/

	// Agregando valoresa un Slice vacío
	letras = append(letras, "a")
	fmt.Println(letras, len(letras), cap(letras)) // imprime slice "letras", longitud y capacidad

	// creando un Slice de numeros
	numbers := make([]int, 3, 3)

	fmt.Println(numbers, len(numbers), cap(numbers)) // imprime slice "numbers", longitud y capacidad

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300

	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 400)

	fmt.Println(numbers, len(numbers), cap(numbers))
}
