package models

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int64  `json:"id"` // reemplazar json por xml o por yaml , depende de que uses
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	createdDate time.Time
}
type Users []User

var userSchema string = `CREATE TABLE IF NOT EXISTS users(
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL UNIQUE ,
    password VARCHAR(64) NOT NULL ,
    email VARCHAR(40) NOT NULL ,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)`

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$") // expreson regular para validar un correo electronico

func NewUser(username, password string, email string) (*User, error) {
	user := &User{Username: username, Email: email}
	if err := user.Valid(); err != nil {
		return &User{}, err
	}
	err := user.SetPassword(password)
	return user, err
}
func CreateUser(username, password string, email string) (*User, error) {
	user, err := NewUser(username, password, email)
	if err != nil {
		return &User{}, err
	}
	err = user.Save()
	return user, err
}
func GetUsers() Users {
	sql := "SELECT id, username, password, email,created_date FROM users"
	users := Users{}
	rows, _ := Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)
		users = append(users, user)
	}
	return users
}
func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.createdDate)

	}
	return user
}

func GetUserByUsername(username string) *User {
	sql := "SELECT id, username,password,email,created_date FROM users WHERE username=?" /// para hacerlo vulnerable Sprintf("SELECT id, username,password,email,created_date FROM users WHERE %s=?",username)
	return GetUser(sql, username)
}
func GetUserById(id int) *User {
	sql := "SELECT id, username,password,email,created_date FROM users WHERE id=?"
	return GetUser(sql, id)

}
func Login(username, password string) (*User, error) {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &User{}, errorLogin
	}
	return user, nil
}

func ValidEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errorEmail
	}
	return nil
}
func (this *User) Save() error {
	if this.Id == 0 {
		return this.insert()
	} else {
		return this.update()
	}
}
func (this *User) insert() error {
	sql := "INSERT users SET username=?,password=?,Email=?"
	id, err := InsertData(sql, this.Username, this.Password, this.Email)
	if err != nil {
		return err
	}
	this.Id = id
	return err
}
func (this *User) update() error {
	sql := "UPDATE users SET username=?,password=?,email=?"
	_, err := Exec(sql, this.Username, this.Password, this.Email)
	return err
}
func (this *User) Delete() error {
	sql := "DELETE FROM users WHERE id=?"
	_, err := Exec(sql, this.Id)
	return err
}
func (this *User) GetcreatedDate() time.Time {
	return this.createdDate
}
func (this *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorPassword
	}

	this.Password = string(hash)
	return nil

}
func (this *User) Valid() error {
	if err := ValidEmail(this.Email); err != nil {
		return err
	}
	if err := ValidUsername(this.Username); err != nil {
		return err
	}
	return nil

}
func ValidUsername(username string) error {
	if username == "" {
		return errorUsername
	}
	if len(username) < 3 {
		return errorShortUsername
	}
	if len(username) > 30 {
		return errorLargeUsername
	}
	return nil
}

// AHora que añadimos la base de datos , debemos cambiar esto
/*
var users = make(map[int]User)

func SetDefaultUser() {
	user := User{Id: 1, Username: "Juan", Password: "12341"}
	users[user.Id] = user
}
func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}
func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New(" El Usuario no existe  friendo! ")
}
func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}
func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user // mientras no tenemos base de datos
}
func DeleteUser(id int) {
	delete(users, id)
}
*/
