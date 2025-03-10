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
//7- ..O sea, el ejercicio 7 sería hacer el ejercicio 1 pero con una función recursiva

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	coeficiente := 1
	fmt.Println(lista)
	valorLimite := 150

	for _, i := range lista {
		if coeficiente < valorLimite {
			total := maths.Multiplicador(i, coeficiente)
			fmt.Println("El resultado es", total)
			coeficiente = total
		} else {
			fmt.Println("Has superado el valor limite ", valorLimite)
		}

	}
}

/*func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	mayor := 0
	fmt.Println(lista)
	for _, i := range lista {
		if mayor < i {
			mayor = i
			fmt.Println("Ahora el nuevo valor mas alto es", mayor)
		} else {
			fmt.Println("El valor ", i, "de la posición X", "no es mayor que el numero ", mayor)
		}
	}
}*/

/*func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	coeficiente := 1
	fmt.Println(lista)

	for _, i := range lista {
		total := maths.Sumador(i, coeficiente)
		fmt.Println("El resultado es", total)
		coeficiente = total
	}
}*/

/* func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	coeficiente := 1
	fmt.Println(lista)

	for _, i := range lista {
		total := maths.Multiplicador(i, coeficiente)
		fmt.Println("El resultado es", total)
		coeficiente = total
	}
}*/

/*func main() {

	lista := []int{2, 5, 6, 8, 9}
	fmt.Println(lista)

	for _, i := range lista {
		resultado := maths.Divisible(i, 4)
		fmt.Println("El resultado es", resultado)
	}



}*/

/*func main (){
	resultado := maths.Par(8)
	fmt.Println("El resultado es", resultado)
}*/
