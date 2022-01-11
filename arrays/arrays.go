package main

import "fmt"

func main() {
	var numeros [5]int
	//numeros =[10, 20, 30, 40, 50]
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)

	letras := [3]string{
		"perro",
		"gato",
		"burro",
	}

	fmt.Println(letras)

	colores := [...]string{
		"amarillo",
		"azul",
		"rojo",
		"blanco",
		"negro",
	}

	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "Euro",
		1: "Dolar",
		3: "Bolívar",
		5: "Pesos",
	}

	fmt.Println(monedas, len(monedas))

	subArray := monedas[:3]

	fmt.Println(subArray, len(subArray))

}
