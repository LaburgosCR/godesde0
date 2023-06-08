package main

import (
	"fmt"
	"github.com/godesde0/variables"
)

// Funcion main de go
func main() {
	estado, texto := variables.ConviertoaText(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}
