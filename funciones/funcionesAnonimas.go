package funciones

import "fmt"

/*funciones que pueden ser asignadas a variables, funciones que se pueden enviar por parametros*/
func Calculos() {

	var numero3 int = 32
	var numero4 int = 234
	/*esta funcion recibe un numero1 y numero2 y va a devolver un entero*/
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}
	fmt.Println(calculo(10, 25))
}
