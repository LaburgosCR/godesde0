package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("primer mensaje")
	defer fmt.Println("mensaje final")

	fmt.Println("ultimo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
