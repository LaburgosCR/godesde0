package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/godesde0/ejercicios"
	/*bufio captura datos por teclado y archivos*/)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	/*va a tener un texto que lo va a sacar de ejercicios.IngreseNumeroTeclado */
	var texto string = ejercicios.IngreseNumeroTeclado()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		/*el return sin nada lo que hace es salir del metodo*/
		return
	}
	/*Fprintln lo va a grabar en archivo que lo creó el Create, y va a grabar el texto*/
	fmt.Fprintln(archivo, texto)
	/*todo archivo que se abre debe quedar cerrado*/
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.IngreseNumeroTeclado()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("error durante el WriteString" + err.Error())
		return false
	}

	arch.Close()
	return true
}

/*para lectura de archivo*/

/*func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))
}
*/

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	/*el scanner.scan devuelve un booleano, si estubo bien o mal*/
	/*por cada lectura de linea correcta vamos hacer el proceso*/
	for scanner.Scan() {
		/*me va a devolver cada linea del archivo*/
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
