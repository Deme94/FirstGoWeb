package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	// Models
	"testnauticos/models"
)

func ValidateRegister(w http.ResponseWriter, r *http.Request) {
	exists := models.UserExistsEmail(r.FormValue("email"))

	if exists {
		io.WriteString(w, "Error")
	} else {
		io.WriteString(w, "")
	}
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := models.NewUser(
		r.FormValue("name"),
		r.FormValue("user"),
		r.FormValue("email"),
		r.FormValue("psswd"),
		r.FormValue("phone"),
		r.Form["courses"],
	)
	var err error
	// Si no existe lo registramos como nuevo usuario
	if !models.UserExistsEmail(r.FormValue("email")) {
		err = models.CreateUser(u)
	} else { // Si existe entonces editamos el usuario
		id, _ := models.GetUserId(u.Email)
		err = models.UpdateUser(id, u)
	}
	if err != nil {
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, "")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	deleted := models.DeleteUser(r.FormValue("id"))
	if deleted {
		io.WriteString(w, "")
	} else {
		io.WriteString(w, "error")
	}

}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	isActivated := models.ActivateUser(r.FormValue("id"))
	if isActivated {
		io.WriteString(w, "1")
	} else {
		io.WriteString(w, "")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	u, _ := models.GetUser(r.FormValue("id"))
	userJson, err := json.Marshal(*u)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(userJson))
}
