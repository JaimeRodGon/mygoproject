package main

import (
	"MyProjects/ProjectOne/maths"
	"fmt"
)

func main() {

	lista := []int{2, 5, 6, 8, 9}
	fmt.Println(lista)

	for _, i := range lista {
		resultado := maths.Divisible(i, 4)
		fmt.Println("El resultado es", resultado)
	}

	//resultado := maths.Par(8)
	//fmt.Println("El resultado es", resultado)

}
