package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngreseNumeroTeclado() {
	var numero1 int
	var err error
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
		fmt.Printf("%d x %d = %d \n", numero1, i, i*numero1)
	}

}
