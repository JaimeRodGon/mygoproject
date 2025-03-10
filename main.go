package main

import (
	"MyProjects/ProjectOne/maths"
	"fmt"
)

//1- Dada una lista de 5 valores, devuelve el resultado de la multiplicacion de toda la lista.
//2- Ejercicio 1 pero en sumatorio.
//3- Dada una lista quiero saber el valor mas alto.
//4- dada una lista desordenada ordena de menor a mayor.
//5- Repetir los ejercicios pero con los datos pedidos por pantalla en vez de declarados en la funcion.
//6- Ejercicio 1 pero devolver error si la multiplicacion es mayor de X. Enfasis en ERROR.
//7- ..

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
