package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 8, 9, 10, 11, 12, 54}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   // slice creado desde un vector desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]  //slice creado desde un vector
	porcion3 := arreglo[6:7] // desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) //va a tener 5 elementos y tiene una capacidad de 20
	fmt.Printf("largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := []int{}
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nlargo %d, Capacidad %d", len(nums), cap(nums))

}
