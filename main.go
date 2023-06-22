package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Ingrese un numero entero de 16 dígitos: ")
	fmt.Scanln(&numero)

	if numero < 1000000000000000 || numero > 9999999999999999 {
		fmt.Print("El numero ", numero, " no tiene 16 digitos")
	}

	verificador := numero % 10

	fmt.Print("El digito verificador es ", verificador)

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
