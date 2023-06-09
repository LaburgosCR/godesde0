package main

import (
	"fmt"
	/*"github.com/godesde0/variables"*/
	"github.com/godesde0/ejercicios"
)

// Funcion main de go
func main() {
	/*estado, texto := variables.ConviertoaText(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "Os X." {
		fmt.Println("esto no es windows, es: ", os)
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("esto es Linux")
	case "darwin":
		fmt.Println("esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
	*/
	numero, text := ejercicios.ConverNum("700")
	fmt.Println(numero)
	fmt.Println(text)
}
