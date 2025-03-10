package maths

import "fmt"

func Par(numero int) string {

	fmt.Println("El numero a comprobar es ", numero)
	if numero%2 == 0 {
		return "Par"
	} else {
		return "Impar"
	}

}

func Divisible(numero int, divisor int) string {
	fmt.Println("El numero a comprobar ", numero)
	if numero%divisor == 0 {
		return fmt.Sprint("es divisible entre", divisor)
	} else {
		return fmt.Sprint("no es divisible entre", divisor)
	}

}

// Entorno
