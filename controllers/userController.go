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
func Inicio(w http.ResponseWriter, r *http.Request) {
	utils.IniciarConexion()
	db := utils.IniciarConexion()
	var usuarios []model.Usuario
	if err := db.Find(&usuarios).Error; err != nil {
		log.Println("Error al obtener los empleados:", err)
		http.Error(w, "Error al obtener los empleados", http.StatusInternalServerError)
		return
	}

	plantillas.ExecuteTemplate(w, "inicio", usuarios)
}
//Metodo para visualizar la pagina de registrar usuarios
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

//Metodo para registrar usuarios
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		utils.IniciarConexion()
		db := utils.IniciarConexion()

		user := model.Usuario{
			Nombre:  nombre,
			Correo: correo,
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error al insertar el empleado:", err)
			http.Error(w, "Error al insertar el empleado", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}