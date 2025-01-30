package main

import (
	"cap1/exercises"
	"fmt"
)

func main() {
	fmt.Println("Ejercicio 1 :")
	exercises.Declaravar()
	fmt.Println()

	fmt.Println("Ejercicio 2 :")
	booleano := exercises.Par(2)
	fmt.Println("Es par? :", booleano)
	fmt.Println()

	fmt.Println("Ejercicio 3 :")
	exercises.Fibonacci(19)
	fmt.Println()

	jerry := exercises.Notes{
		Id:    4,
		Name:  "and",
		Esono: false,
	}
	fmt.Println(jerry)
	fmt.Println(jerry.Id)
	fmt.Println(jerry.Name)
}
