package handlers

import (
	"api/utils"
	"net/http"
)

type customHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// aquí va nuestra logíca
		if !utils.IsAuthenticated(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}
		function(w, r)
	})

}
func MiddlewateTow(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
