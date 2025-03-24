package controllers

import(
	"fmt"
	"net/http"
	"text/template"
)

func SignupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
	 username := r.FormValue("username")
	 password := r.FormValue("password")
   

	 fmt.Printf("New user signup: Username - %s, Password - %s\n", username, password)
   
	 http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	 return
	}
   
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
	 http.Error(w, err.Error(), http.StatusInternalServerError)
	 return
	}
	tmpl.Execute(w, nil)
   }