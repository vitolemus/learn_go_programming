package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(expression string) int {
	values := strings.Split(expression, "+") // separamos los caracteres por el simbolo +

	var result int

	for _, value := range values { // inciamos el proceso de conversion y suma de numeros

		num, error := strconv.Atoi(value) // convertimos los caracteres a numeros enteros

		if error != nil {
			fmt.Println("Ingrese un número entero..!")                  // manejamos el error de ingresar un caracter que no sea un numero entero
			fmt.Println("Por el momento solo podemos realizar SUMA..!") // O que no sea la función suma
		} else {
			result += num // realizamos la suma
		}

	}

	return result // devolvemos el resultado de la suma

}

// func resta(a, b int)int {
// 	return a-b

// }

// func mult(a, b int)int {
// 	return a*b

// }

// func div(a, b int)int {
// 	return a+b

// }

func main() {

	fmt.Println("*****CALCULADORA*****")
	fmt.Println("*Enter your operation, please*")

	var result int        // Guardamos el resultado de la operación
	var expression string // Guardamos la operación a realizar

	fmt.Print("=>")
	fmt.Scanln(&expression) // Recibimos por teclado la operación a realizar

	result = add(expression) // llamamos a la función sumar
	fmt.Println(result)      // lImprimos el rsultado de la operación

}
