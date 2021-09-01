package test

import (
	"api/models"
	"testing"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("NO se pudo conectar a la base de datos", nil)
	}
}
