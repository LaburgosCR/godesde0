package main

import (
	"github.com/godesde0/webserver"
	/*"github.com/godesde0/models"*/)

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
	/*
		numero, text := ejercicios.ConverNum("500")
		fmt.Println(numero)
		fmt.Println(text)
	*/
	/*
		teclado.IngresoNumerico()
	*/
	/*fmt.Println(ejercicios.IngreseNumeroTeclado())*/
	/*files.SumaTabla()*/
	/*files.LeoArchivo()*/
	/*mapas.MostrarMapas()*/
	/*users.AlraUsuario()*/
	/*
		Pedro := new(models.Hombre)
		e.HumanosRespirando(Pedro)
		Maria := new(models.Mujer)
		e.HumanosRespirando(Maria)
	*/
	/*
		canal1 := make(chan bool)
		go goroutines.MiNombreLento("Luis Diego", canal1)
		defer func() {
			<-canal1
		}()
		fmt.Println("estoy aqui")
	*/
	webserver.MiWebServer()
}
