package models

import "errors"

type ValidationError error

var (
	errorUsername      = ValidationError(errors.New("El nombre de usuario no pude ser vacio"))
	errorShortUsername = ValidationError(errors.New("El nombre de usuario es muy corto"))
	errorLargeUsername = ValidationError(errors.New("El nombre de usuario no  es muy largo"))

	errorEmail    = ValidationError(errors.New("Formato invalido de Email"))
	errorPassword = ValidationError(errors.New("Contraseña Invalida "))
	errorLogin    = ValidationError(errors.New("Hola mama "))
)
