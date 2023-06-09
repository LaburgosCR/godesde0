package ejercicios

import "strconv"

/*esta función recibe un parametro string y devuelve 2 valores int y string*/
/*para que sea funcion publica la primer letra debe ser mayuscula*/
func ConverNum(text string) (int, string) {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "es mayor a 100"
	} else {
		return num, "es menor a 100"
	}
}
