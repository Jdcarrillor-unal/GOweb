package main

// EndPoint
import (
	"log"
	"net/http"

	"api/handlers"
	"api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//middleware
	editHandler := handlers.Authentication(handlers.EditUser)
	homeHandler := handlers.Authentication(handlers.Home)

	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.Index).Methods("GET")
	mux.HandleFunc("/users/login", handlers.Login).Methods("POST", "GET")
	mux.HandleFunc("/users/new", handlers.NewUser).Methods("POST", "GET")
	mux.HandleFunc("/users/logout", handlers.Logout).Methods("GET")
	mux.Handle("/users/edit", editHandler).Methods("GET")
	mux.Handle("/users/home", homeHandler).Methods("POST", "GET")
	// EndPoints ---> rutas del servidor ////
	mux.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	mux.HandleFunc("/api/v2/datos/fecha/{inicio}&{fin}", handlers.GetDataByTime).Methods("GET")
	mux.HandleFunc("/api/v2/datos/limite/{limit:[0-9]+}", handlers.GetDataByNum).Methods("GET")
	// mux.HandleFunc("/api/v2/datos/temp/{limit:[0-9]+}", handlers.GetTempByNum).Methods("GET")
	// mux.HandleFunc("/api/v2/datos/humedad/{limit:[0-9]+}", handlers.GetHumedadByNum).Methods("GET")
	// mux.HandleFunc("/api/v2/datos/pwm/{limit:[0-9]+}", handlers.GetpwmByNum).Methods("GET")

	// Servir archivos estaticos
	assets := http.FileServer(http.Dir("assets"))
	statics := http.StripPrefix("/assets/", assets)
	mux.PathPrefix("/assets/").Handler(statics)

	log.Println("EL servidor esta a la escucha en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	models.CloseConnection()
}
