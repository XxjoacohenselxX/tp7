package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Ingrese un numero entero de 16 dígitos: ")
	fmt.Scanln(numero)

	verificador := numero % 10
	acumulador := 0
	for i := 0; i < 15; i++ {
		//Extraigo el ultimo dígito de la derecha
		digito := numero % 10

		if i%2 == 0 {
			//posicion par
			//TODO actualizar el acumulador
			digito += acumulador

		} else {
			//posicion impar
			//TODO actualizar el acumulador
			digito += acumulador
		}

		//Elimino el ultimo dígito de la derecha
		numero = numero / 10
	}

	resultado := (acumulador * 9) % 10

	if resultado == verificador {

	} else {

	}

}
