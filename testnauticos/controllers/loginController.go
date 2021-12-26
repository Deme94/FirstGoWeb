package controllers

import (
	"io"
	"net/http"
	"testnauticos/models"

	"github.com/google/uuid"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	id, err := models.Login(r.FormValue("email"), r.FormValue("password"))

	if err != nil {
		// Login error
		io.WriteString(w, "error")
	} else {
		// Login successful
		// Create session
		sID := uuid.NewString()
		c := &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		http.SetCookie(w, c)
		models.SetSession(id, c.Value)

		// Return response to client which will redirect the user to URL: /admin/usuarios
		io.WriteString(w, "")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	// ToDo: Delete session from DB
	// Remove cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AlreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	sessionExists := models.SessionExists(c.Value)
	if !sessionExists {
		// ToDo: Delete cookie
		return false
	}
	return true
}

func GetLoggedUser(r *http.Request) (*models.User, error) {
	c, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}
	user, err := models.GetUserBySession(c.Value)
	if err != nil {
		return nil, err
	}
	return user, err
}
