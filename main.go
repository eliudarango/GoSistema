package main

import (
	"gosistema/controllers"
	"log"
	"net/http"
)


func main() {

	//Metodo GET 
//	http.HandleFunc("/", controllers.LoginPage)
    http.HandleFunc("/login", controllers.LoginPage)
    http.HandleFunc("/welcome", controllers.WelcomePage)
	http.HandleFunc("/", controllers.UserPage)
	http.HandleFunc("/registrarUsuario", controllers.RegisterPage)
	//Metodo POST
	http.HandleFunc("/insertarUsuario", controllers.RegisterUser)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)

}



