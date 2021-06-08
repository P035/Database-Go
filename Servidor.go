// En este paquete se declarará todo lo necesario para correr el servidor: Como
// las funciones handler, funciones para iniciar el servidor, etc.
package Servidor

import (
	"fmt"
	"html/template"
	"net/http"
)

// Index sera la función handler que devolverá el html index.html
func index(Writer http.ResponseWriter, Request *http.Request) {

	HTML, err := template.ParseFiles("./templates/index.html")
	if err != nil {

		fmt.Printf("Error cargando html: %v", err)
	} else {

		err := HTML.Execute(Writer, nil)
		if err == nil {

			fmt.Println("Html envíado a: ", Request.RemoteAddr)
		} else {

			fmt.Println("Error envíando html a: ", Request.RemoteAddr)
		}
	}
}

// Login sera la función handler que devolverá el html de login.html
func login(Writer http.ResponseWriter, Request *http.Request) {

	HTML, err := template.ParseFiles("./templates/login.html")
	if err != nil {

		fmt.Printf("Error cargando html: %v", err)
	}else {

		err := HTML.Execute(Writer, nil)
		if err == nil {

			fmt.Println("Html envíado a: ", Request.RemoteAddr)
		} else {

			fmt.Println("Error envíando html a: ", Request.RemoteAddr)
		}
	}
}

// Register sera la función handler que devolverá el html de register.html
func register(Writer http.ResponseWriter, Request *http.Request) {

	if Request.Method == "POST"{

		
	}
	HTML, err := template.ParseFiles("./templates/register.html")
	if err != nil {

		fmt.Printf("Error cargando html: %v", err)
	}else {

		err := HTML.Execute(Writer, nil)
		if err == nil {

			fmt.Println("Html envíado a: ", Request.RemoteAddr)
		} else {

			fmt.Println("Error envíando html a: ", Request.RemoteAddr)
		}
	}
}


var Servidor http.Server

func Inicializar() {

	Servidor.Addr = "172.20.10.9:80"
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	err := Servidor.ListenAndServe()
	fmt.Println(err)
}
