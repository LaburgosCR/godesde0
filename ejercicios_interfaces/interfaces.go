package ejercicios_interfaces

import (
	"fmt"

	"github.com/godesde0/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("soy un/a %s, y estoy respirando \n", humano.Sexo())
}
