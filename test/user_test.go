package test

import (
	"api/models"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var user *models.User

const (
	id           = 1
	username     = "eduardo_gpg"
	password     = "password"
	passwordHash = "$2a$10$09jehQ.wFwxAqFEjcuCWzOUFSZNZxwtKf7rpF3jDGVR9e4znJUDn2"
	email        = "correo@correo.com"
	createdDate  = "2021-08-18"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if !equalUser(user) || err != nil {
		t.Error("No es posible crear objeto")
	}
}
func TestSave(t *testing.T) {
	user, _ := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("NO es posible crear el usuario", err)
	}
}
func TestPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60 {
		t.Error("No es posible cifrar el password")
	}
}

// func TestLogin(t *testing.T) {
// 	if valid := models.Login(username, password); !valid {
// 		t.Error("No es posible reliazr el login")
// 	}
// }
func TestValidEmail(t *testing.T) {
	if err := models.ValidEmail(email); err != nil {
		t.Error("Validación errorena del Email")
	}
}
func TestInvalidEmail(t *testing.T) {
	if err := models.ValidEmail("asdajsaldsa"); err == nil {
		t.Error("La expresion regular NO  funciona con correos invalidos")
	}
}

// func TestNologin(t *testing.T) {
// 	if valid := models.Login(randomUsername(), password); valid {
// 		t.Error("Es posible reliazr el login con parametros incorrectos")
// 	}
// }
func TestCreatUser(t *testing.T) {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto", err)
	}
}
func TestUsernameLenght(t *testing.T) {
	newUsername := username
	for i := 0; i < 10; i++ {
		newUsername += newUsername
	}
	_, err := models.NewUser(newUsername, password, email)
	if err == nil {
		t.Error("Es posible generar un Usuario con un Username muy grande")
	}
}
func TestUniqueUser(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err == nil {
		t.Error("Es posible insertar registros con username duplicados")
	}
}
func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'users.username'", username)
	if err.Error() != message {
		t.Error("Es posible tener usernames duplicados v2")
	}
}
func TestGetUser(t *testing.T) {
	user := models.GetUserById(id)

	if !equalUser(user) || !equalsCreatedDate(user.GetcreatedDate()) {
		t.Error("No es posible obtener el usuario ")
	}
}
func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener los usuarios ")
	}
}
func equalUser(user *models.User) bool {
	if user.Username == username && user.Email == email {
		return true
	}
	return false
}
func equalsCreatedDate(date time.Time) bool {
	t, _ := time.Parse("2001-01-02", createdDate)
	return t == date
}
func TestDeleteUser(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("No es posible eliminar al usuario")
	}
}

func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}
