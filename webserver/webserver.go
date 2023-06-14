package webserver

import "net/http"

/*Handlefunc es un manejador que recibe la peticion*/
/*cuando desde main llame a la funcion Mi webserver*/
func MiWebServer() {
	http.HandleFunc("/", home)        //manejador que recibe la peticion, ejecute la funcion home
	http.ListenAndServe(":3000", nil) //le asignamos el puerto que vamos a usar, escuchar y servir el puerto
}

/*a donde tiene que enviar esa funcion*/
/*w y r son variables*/
func home(w http.ResponseWriter, r *http.Request) { //esta funcion recibe 2 parametros, que son 2 interfaces de net http
	//w -> writer r ->request
	http.ServeFile(w, r, "./webserver/index.html")
	// la funcion serfile permite servir un archivo en este caso el html al web
}
