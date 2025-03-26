package routes

import (
	"log"
	"net/http"
	"gosistema/backend/controllers"
)

func LoadRouters() {
	//METODOS GET
	//	http.HandleFunc("/", controllers.LoginPage)
	//http.HandleFunc("/login", controllers.LoginPage)
	//http.HandleFunc("/welcome", controllers.WelcomePage)
	http.HandleFunc("/", controllers.Inicio)
	http.HandleFunc("/crear", controllers.Crear)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/borrar", controllers.Borrar)
	// METODOS POST
	http.HandleFunc("/insertar", controllers.Insertar)
	http.HandleFunc("/actualizar", controllers.Actualizar)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}


