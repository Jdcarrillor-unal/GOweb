package handlers

import (
	"api/models"
	"api/utils"
	"fmt"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	fmt.Println(r.Method)
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		if user, err := models.CreateUser(username, password, email); err != nil {
			errorMessage := err.Error()
			context["Error"] = errorMessage
		} else {
			utils.SetSession(user, w)
			context["Exito"] = "Creado correctamente "
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
	}

	utils.RenderTemplate(w, "users/new", context)
}
func Login(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if user, err := models.Login(username, password); err != nil {
			context["Error"] = err.Error()
		} else {
			utils.SetSession(user, w)
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}

	}
	utils.RenderTemplate(w, "users/login", context)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteCookie(w, r)
	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}
func EditUser(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	user := utils.GetUser(r)
	context["User"] = user

	utils.RenderTemplate(w, "users/edit", context)

}
func Home(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	if r.Method == "GET" {
		fmt.Println(r.Method)
		user := utils.GetUser(r)
		data, _ := models.GetlastDataByUserID(user.Id)
		context["Data"] = data
		context["User"] = user
		fmt.Println("context:", context)
		fmt.Println("user: ", user)
		fmt.Println("data", data)
		utils.RenderTemplate(w, "users/home", context)
	}
	// fmt.Println(data.Fecha_hora)
	// fmt.Println(data.Temperatura)
	// fmt.Println(data.Humedad)
	// fmt.Println(data.Led1)

	if r.Method == "POST" {

		init_day := r.FormValue("dia_init")
		init_houre := r.FormValue("hora_init")
		end_day := r.FormValue("dia_fin")
		end_houre := r.FormValue("hora_fin")
		begin := init_day + " " + init_houre
		end := end_day + " " + end_houre
		fmt.Println(begin, end)
		if datos, err := models.GetDataBydate(begin, end); err != nil {
			fmt.Println(err)
		} else {
			context["Data"] = datos
			fmt.Println(datos)
			utils.RenderTemplate(w, "users/home", context)
		}
	}

}
