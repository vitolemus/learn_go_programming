package main

import "fmt"

func main() {
	//Slices
	numeros := []int{10, 20, 30}

	fmt.Println(numeros, len(numeros))

	//agregar datos al slice

	numeros = append(numeros, 40, 50)

	fmt.Println(numeros, len(numeros))

	// Obtener un Sub Slice

	subSlice := numeros[:3]
	fmt.Println(subSlice, len(subSlice))

	numeros[0] = 100

	fmt.Println(numeros)

	//Capacidad de los slices

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Longitud: %v, Capacidad: %v, Dirección de Memoria: %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Longitud: %v, Capacidad: %v, Dirección de Memoria: %p \n", len(meses), cap(meses), meses)

}
