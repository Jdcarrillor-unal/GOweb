package handlers

import (
	"fmt"
	"net/http"

	// "log"
	//	"encodign/xml" en caso de quere usar xml
	/// https://github.com/go-yaml/yaml  otro formato a usar
	"api/models"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	if err := user.Valid(); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	user.SetPassword(user.Password)
	if err := user.Save(); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	models.SendData(w, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	request := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(request); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	if err := user.Valid(); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	user.Username = request.Username
	user.Email = request.Email
	user.SetPassword(request.Password)
	if err := user.Save(); err != nil {
		models.SendUnprocessedEntity(w)
		return
	}
	models.SendData(w, user)

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		user.Delete()
		models.SendNoContent(w)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"]) // string
	user := models.GetUserById(userId)
	if user.Id == 0 {
		return *user, errors.New(" Este Usuario no existe ")
	}
	return *user, nil
}
func GetDataByTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inicio := vars["inicio"]
	fin := vars["fin"]
	datos, err := models.GetDataBydate(inicio, fin)
	if err != nil {
		panic(err)
	}
	fmt.Println(datos)
	fmt.Println(w)
	models.SendData(w, datos)
}
func GetDataByNum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit, _ := strconv.Atoi(vars["limit"])
	datos, err := models.GetDataByNum(limit)
	if err != nil {
		panic(err)
	}
	fmt.Println(datos)
	models.SendData(w, datos)
}
func GetTempByNum(w http.ResponseWriter, r *http.Request) {

}

/*
Old FUNC GetUser

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := models.GetDeafaultMessage()
	userId, _ := strconv.Atoi(vars["id"]) // string
	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound(err.Error())
	} else {
		response.Data = user
	}
	w.Header().Set("Content-Type", "application/json") // text/xml && con yaml no se cambia el Header
	output, _ := json.Marshal(&response)               // serializar el objeto
	//	output, _ := xml.Marshal(&user) // serializar el objeto
	//	output, _ := yaml.Marshal(&user) // serializar el objeto
	fmt.Fprintf(w, string(output)+"\n")
}
*/
