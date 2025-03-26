package controllers

import (
	"gosistema/backend/models"
	"gosistema/backend/utils"
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("frontend/templates/*"))

// Metodo para visualizar la pagina de usuarios
func Inicio(w http.ResponseWriter, r *http.Request) {
	utils.IniciarConexion()
	db := utils.IniciarConexion()
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Println("Error al obtener los empleados:", err)
		http.Error(w, "Error al obtener los empleados", http.StatusInternalServerError)
		return
	}

	plantillas.ExecuteTemplate(w, "inicio", users)
}

// Metodo para visualizar la pagina de registrar usuarios
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// Metodo para registrar usuarios
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		utils.IniciarConexion()
		db := utils.IniciarConexion()

		user := models.User{
			Name: name,
			Email: email,
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error al insertar el empleado:", err)
			http.Error(w, "Error al insertar el empleado", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// Metodo para borrar usuarios
func Borrar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()

	if err := db.Delete(&models.User{}, id).Error; err != nil {
		log.Println("Error al borrar el usuario:", err)
		http.Error(w, "Error al borrar el usuario", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Metodo para visualizar la pagina de editar usuario
func Editar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		log.Println("Error al encontrar el usuario:", err)
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	plantillas.ExecuteTemplate(w, "editar", user)
}

// Metodo para actualizar los datos de un usuario
func Actualizar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db := utils.IniciarConexion()

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		log.Println("Error al encontrar el usuario:", err)
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		user.Name = name
		user.Email = email
		if err := db.Save(&user).Error; err != nil {
			log.Println("Error al actualizar el usuario:", err)
			http.Error(w, "Error al actualizar el usuario", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
