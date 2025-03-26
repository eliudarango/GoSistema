package controllers

import (
	"gosistema/backend/models"
	"gosistema/backend/utils"
	"log"
	"net/http"
)


// Metodo para visualizar la pagina de usuarios
func InicioBlog(w http.ResponseWriter, r *http.Request) {
	utils.IniciarConexion()
	db := utils.IniciarConexion()
	var blogs []models.Blog
	if err := db.Find(&blogs).Error; err != nil {
		log.Println("Error al obtener los blogs:", err)
		http.Error(w, "Error al obtener los blogs", http.StatusInternalServerError)
		return
	}

	plantillas.ExecuteTemplate(w, "inicio", blogs)
}

// Metodo para visualizar la pagina de registrar usuarios
func CrearBlog(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// Metodo para registrar usuarios
func InsertarBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		utils.IniciarConexion()
		db := utils.IniciarConexion()

		blog := models.Blog{
			Title: title,
			Content: content,
		}
		if err := db.Create(&blog).Error; err != nil {
			log.Println("Error al insertar el blog:", err)
			http.Error(w, "Error al insertar el blog", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// Metodo para borrar usuarios
func BorrarBlog(w http.ResponseWriter, r *http.Request) {

}

// Metodo para editar usuarios
func EditarBlog(w http.ResponseWriter, r *http.Request) {

}
