package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*va a devolver un dato de tipo string, no lleva parentesis por que es el unico que lleva*/
func IngreseNumeroTeclado() string {
	var numero1 int
	var err error
	var texto string
	/*stdin -> standar input*/
	/*scanner es mi objeto*/
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero: ")

		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		/*Sprintf lo devuelve en string*/
		/*el += concatena*/
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, i*numero1)
	}

	return texto
}
