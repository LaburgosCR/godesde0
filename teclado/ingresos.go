package teclado

import (
	"bufio"
	"fmt"
	"os" /*operative system*/
	"strconv"
)

/*variable + nombre + tipo*/
var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumerico() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1 : ")
	/*scanner es mi objeto*/
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}

	}
	fmt.Println("Ingrese numero 2 : ")
	/*scanner es mi objeto*/
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}

	}
	fmt.Println("Ingrese leyenda : ")
	/*scanner es mi objeto*/
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Println(leyenda, numero1*numero2)
}
