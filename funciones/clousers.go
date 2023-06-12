package funciones

import "fmt"

/*funcion Tabla que recibe valor de tipo int y va a devolver una func anonima de tipo entero*/
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClousure() {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i < 10; i++ {
		fmt.Println(MiTabla())
	}
}
