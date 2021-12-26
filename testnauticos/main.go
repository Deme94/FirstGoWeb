package main

import (
	"fmt"
	"html/template"
	"net/http"

	// Controllers
	"testnauticos/controllers"
	"testnauticos/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles(
		"views/adminUsers.gohtml",
		"views/login.gohtml",
		"views/components/tablaUsers.gohtml",
		"views/components/navbar.gohtml",
		"views/components/footer.gohtml",
	))
}

func main() {
	// Setup Database
	models.SetupDB()

	// Server
	mux := http.NewServeMux()

	// Serves files
	mux.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./views"))))
	// Handles /favicon.ico request
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	// Routes (VIEWS)
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/admin/usuarios", AdminUsers)
	// Routes (API REST)
	mux.HandleFunc("/validateRegister", ValidateRegister)
	mux.HandleFunc("/saveUser", SaveUser)
	mux.HandleFunc("/activateUser", ActivateUser)
	mux.HandleFunc("/deleteUser", DeleteUser)
	mux.HandleFunc("/getUser", GetUser)
	mux.HandleFunc("/loginUser", LoginUser)
	mux.HandleFunc("/logout", Logout)

	// Listen and serve
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}

// Views
func Index(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/usuarios", http.StatusSeeOther)
	}
}

func AdminUsers(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user, err := controllers.GetLoggedUser(r)
	users, _ := models.GetAllUsers()
	data := struct {
		LoggedUser models.User
		Users      []models.User
	}{
		LoggedUser: *user,
		Users:      users,
	}
	err = tpl.ExecuteTemplate(w, "adminUsers.gohtml", data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found", 404)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found", 404)
	}
}

// API REST
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.LoginUser(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.Logout(w, r)
}

func ValidateRegister(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.ValidateRegister(w, r)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.SaveUser(w, r)
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.ActivateUser(w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.DeleteUser(w, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !controllers.AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	controllers.GetUser(w, r)
}
