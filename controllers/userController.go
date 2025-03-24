package controllers


import(
	"gosistema/model"
	"gosistema/utils"
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

//Metodo para visualizar la pagina de usuarios
func UserPage(w http.ResponseWriter, r *http.Request) {
	utils.IniciarConexion()
	db := utils.IniciarConexion()
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		log.Println("Error al obtener los empleados:", err)
		http.Error(w, "Error al obtener los empleados", http.StatusInternalServerError)
		return
	}

	plantillas.ExecuteTemplate(w, "inicio", users)
}
//Metodo para visualizar la pagina de registrar usuarios
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

//Metodo para registrar usuarios
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		utils.IniciarConexion()
		db := utils.IniciarConexion()

		user := model.User{
			Name:  name,
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