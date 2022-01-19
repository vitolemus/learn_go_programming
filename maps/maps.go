package main

import "fmt"

func main() {

	// Creando un mapa con la función MAKE
	animals := make(map[string]string)
	animals["cat"] = "gato"
	animals["dog"] = "perro"
	animals["monkey"] = "mono"

	fmt.Println(animals)

	// Creando un mapa LITERAL
	fruits := map[string]string{
		"apple":     "manzana",
		"banana":    "cambur",
		"pineapple": "piña",
	}

	fmt.Println(fruits)

	// Consultar un valor dentro de un mapa
	fmt.Println(fruits["banana"])

	// Consultar un valor que NO EXISTE dentro de un mapa (Devuelve vacío)
	fmt.Println(fruits["melon"])

	consulta, ok := animals["donkey"]
	fmt.Println(consulta, ok) // Devuelve vacío y false

	if ok == false {
		animals["donkey"] = "burro"
	}

	fmt.Println(animals) // agrega donkey al map
	/*
		if consulta, ok := animals["donkey"]; !ok{
			animals["donkey"] = "burro"
		}
	*/

	cars := map[int]string{
		0: "chevrolet",
		1: "ford",
		2: "chrysler",
	}

	fmt.Println(cars)

}
