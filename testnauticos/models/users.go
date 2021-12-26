package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	Id           string
	Name         string
	Username     string
	Email        string
	Password     string
	Phone        string
	Courses      []string
	Created_at   string
	Last_visited string
	Status       bool
}

func NewUser(name string, username string, email string, password string, phone string, courses []string) *User {
	return &User{"", name, username, email, password, phone, courses, "", "", false}
}

func CreateUser(u *User) error {
	courses, _ := json.Marshal((*u).Courses)
	_, err := db.Exec(`INSERT INTO usuarios (nombre, usuario, email, password, telefono, cursos, created_at)` +
		`VALUES ('` + (*u).Name + `', '` + (*u).Username + `', '` + (*u).Email + `', '` + (*u).Password + `', ` +
		(*u).Phone + `, '` + string(courses) + `', NOW());`)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetUser(id string) (*User, error) {
	var name string
	var username string
	var email string
	var password string
	var phone string
	var courses string
	var created_at string
	var last_visited sql.NullString
	var strStatus string
	var status bool

	row := db.QueryRow(`SELECT id, nombre, usuario, email, password, telefono, cursos, created_at, last_visited, status ` +
		`FROM usuarios WHERE id='` + id + `';`)
	err := row.Scan(&id, &name, &username, &email, &password, &phone, &courses, &created_at, &last_visited, &strStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	if strStatus == "1" {
		status = true
	} else {
		status = false
	}
	var coursesSlice []string
	json.Unmarshal([]byte(courses), &coursesSlice)
	u := User{
		Id:           id,
		Name:         name,
		Username:     username,
		Email:        email,
		Password:     password,
		Phone:        phone,
		Courses:      coursesSlice,
		Created_at:   created_at,
		Last_visited: last_visited.String,
		Status:       status,
	}
	return &u, err
}

func GetUserId(email string) (string, error) {
	var id string

	row := db.QueryRow(`SELECT id FROM usuarios WHERE email='` + email + `';`)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	return id, err
}

func Login(email string, password string) (string, error) {
	var id string

	row := db.QueryRow(`SELECT id FROM usuarios WHERE email='` + email + `' AND ` +
		`password='` + password + `';`)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	} else {
		_, err = db.Exec(`UPDATE usuarios SET ` +
			`last_visited = NOW()` +
			` WHERE id = ` + id + `;`)
		if err != nil {
			panic(err)
		}
	}
	return id, err
}

func UpdateUser(id string, u *User) error {
	courses, _ := json.Marshal((*u).Courses)
	_, err := db.Exec(`UPDATE usuarios SET ` +
		`nombre = '` + (*u).Name +
		`', usuario = '` + (*u).Username +
		`', email = '` + (*u).Email +
		`', password = ` + (*u).Password +
		`, telefono = ` + (*u).Phone +
		`, cursos = '` + string(courses) +
		`' WHERE id = ` + id + `;`)
	if err != nil {
		panic(err)
	}

	return err
}

func DeleteUser(id string) bool {
	var err error
	_, err = db.Exec(`DELETE FROM usuarios WHERE id = ` + id + `;`)
	if err != nil {
		return false
	} else {
		return true
	}
}

func ActivateUser(id string) bool {
	var err error
	isActivated := GetStatus(id)
	if isActivated {
		_, err = db.Exec(`UPDATE usuarios SET status = '` + "0" + `' WHERE id='` + id + `';`)

	} else {
		_, err = db.Exec(`UPDATE usuarios SET status = ` + "1" + ` WHERE id=` + id + `;`)
	}

	if err != nil {
		panic(err)
	}

	return isActivated
}

func GetStatus(id string) bool {
	var status int
	row := db.QueryRow(`SELECT status FROM usuarios WHERE id='` + id + `';`)
	err := row.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	if status == 0 {
		return false
	} else {
		return true
	}
}

func UserExists(id string) bool {
	row := db.QueryRow(`SELECT id FROM usuarios WHERE id='` + id + `';`)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
		return false
	}

	return true
}
func UserExistsEmail(email string) bool {
	var id string

	row := db.QueryRow(`SELECT id FROM usuarios WHERE email='` + email + `';`)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
		return false
	}

	return true
}
func GetAllUsers() ([]User, error) {
	var id string
	var name string
	var username string
	var email string
	var phone string
	var courses string
	var created_at string
	var last_visited sql.NullString
	var strStatus string
	var status bool

	var users []User

	rows, err := db.Query(`SELECT id, nombre, usuario, email, telefono, cursos, created_at, last_visited, status ` +
		`FROM usuarios;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name, &username, &email, &phone, &courses, &created_at, &last_visited, &strStatus)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err)
			}
		}

		if strStatus == "1" {
			status = true
		} else {
			status = false
		}
		var coursesSlice []string
		json.Unmarshal([]byte(courses), &coursesSlice)
		users = append(users, User{
			Id:           id,
			Name:         name,
			Username:     username,
			Email:        email,
			Phone:        phone,
			Courses:      coursesSlice,
			Created_at:   created_at,
			Last_visited: last_visited.String,
			Status:       status,
		})
	}
	return users, err
}

func SetSession(id string, uuid string) {
	_, err := db.Exec(`INSERT INTO sesiones (idUsuario, uuid) VALUES('` + id + `','` + uuid + `') ON DUPLICATE KEY UPDATE ` +
		`uuid='` + uuid + `';`)
	if err != nil {
		panic(err)
	}
}

func GetUserBySession(uuid string) (*User, error) {
	var id string
	var name string
	var username string
	var email string
	var password string
	var phone string
	var courses string
	var created_at string
	var last_visited sql.NullString
	var strStatus string
	var status bool

	row := db.QueryRow(`SELECT id, nombre, usuario, email, password, telefono, cursos, created_at, last_visited, status ` +
		`FROM usuarios ` +
		`INNER JOIN sesiones ON id=idUsuario AND uuid='` + uuid + `';`)
	err := row.Scan(&id, &name, &username, &email, &password, &phone, &courses, &created_at, &last_visited, &strStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	if strStatus == "1" {
		status = true
	} else {
		status = false
	}
	var coursesSlice []string
	json.Unmarshal([]byte(courses), &coursesSlice)
	u := User{
		Id:           id,
		Name:         name,
		Username:     username,
		Email:        email,
		Password:     password,
		Phone:        phone,
		Courses:      coursesSlice,
		Created_at:   created_at,
		Last_visited: last_visited.String,
		Status:       status,
	}
	return &u, err
}

func SessionExists(uuid string) bool {
	var id string

	row := db.QueryRow(`SELECT idUsuario FROM sesiones WHERE uuid='` + uuid + `';`)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
		return false
	}

	return true
}
