package controllers


import(
	"fmt"
	"net/http"
	"text/template"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")


        if username == "admin" && password == "admin" {
            http.Redirect(w, r, "/welcome", http.StatusSeeOther)
            return
        }

        fmt.Fprintf(w, "Invalid credentials. Please try again.")
        return
    }

    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}


func WelcomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, you have successfully logged in!")
}