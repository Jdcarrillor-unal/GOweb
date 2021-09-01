package handlers

import (
	"api/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "application/index", nil)

}
