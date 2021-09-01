package config

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
}

var database *DatabaseConfig

func init() { // la primera funcion que se ejecuta cuando el paquete se mada a llamar
	os.Setenv("USERNAME", "")
	os.Setenv("PASSWORD", "")
	os.Setenv("HOST", "")
	os.Setenv("PORT", "")
	os.Setenv("DATABASE", "")
	database = &DatabaseConfig{}
	database.username = os.Getenv("USERNAME")
	database.password = os.Getenv("PASSWORD")
	database.host = os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	database.port = port
	database.database = os.Getenv("DATABASE")
}
func GetUrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", this.username, this.password, this.host, this.port, this.database)
}

func DirTemplate() string {
	return "templates/**/*.html"
}
func DirTempalteError() string {
	return "templates/error.html"
}
