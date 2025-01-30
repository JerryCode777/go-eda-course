package exercises

import (
	"fmt"
	"math/rand"
)

// programa que declara variables de diferentes tipos y las imprima
func Declaravar() {
	var a int
	var b float32
	var c float64
	var d string
	var e bool
	var array [10]int

	a = 12
	b = 12.32
	c = 23.2343
	d = "jerry"
	e = true

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(3)
	}

	slice := []int{1, 2, 3, 4}

	fmt.Println(a, b, c, d, e, array, slice)

	//declaremos la estructura
	structura := Notes{
		Id:    3,
		Name:  "jery",
		Esono: true,
	}

	fmt.Println(structura)
}
