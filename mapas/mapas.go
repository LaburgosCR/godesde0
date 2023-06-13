package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 36,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}
	/*para eliminar elementos de un mapa se usa el delete*/
	/*el primer elemento es el map y el segundo la clave*/
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
