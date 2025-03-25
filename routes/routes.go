package routes

import (
	"log"
	"net/http"
	"gosistema/controllers"
)

func LoadRouters() {
	//METODOS GET
	//	http.HandleFunc("/", controllers.LoginPage)
	//http.HandleFunc("/login", controllers.LoginPage)
	//http.HandleFunc("/welcome", controllers.WelcomePage)
	http.HandleFunc("/", controllers.Inicio)
	http.HandleFunc("/crear", controllers.Crear)
	// METODOS POST
	http.HandleFunc("/insertar", controllers.Insertar)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}


