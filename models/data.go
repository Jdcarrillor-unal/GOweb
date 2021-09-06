package models

import (
	"fmt"
	"time"
)

type Data struct {
	Id          int       `json:"id"`
	Temperatura float64   `json:"temperatura"`
	Humedad     float64   `json:"humedad"`
	Led1        string    `json:"boton"`
	Pwm         float64   `json:"pwn"`
	Fecha_hora  time.Time `json:"fecha_hora"`
}
type Datos []Data

var dataschema string = `CREATE TABLE IF NOT EXISTS DATA (
	ID int(11) PRIMARY KEY  NOT NULL AUTO_INCREMENT ,
	user_id INT(6) UNSIGNED DEFAULT 1,
	temperatura FLOAT NOT NULL DEFAULT 0,
	humedad FLOAT NOT NULL DEFAULT 0,
	led1 VARCHAR(8) NOT NULL DEFAULT 0 ,
	pwm FLOAT NOT NULL,
	fecha_hora DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREING KEY (user_id) REFERENCES users(id) 	 
) `

func GetData(sql string, conditional interface{}) *Data {
	data := &Data{}
	rows, err := Query(sql, conditional)
	if err != nil {
		println(err)
		return data
	}
	for rows.Next() {
		rows.Scan(&data.Id, &data.Temperatura, &data.Humedad, &data.Led1, &data.Pwm, &data.Fecha_hora)
	}
	return data
}
func GetDatos(sql string) (Datos, error) {
	var datos Datos
	rows, error := Query(sql)
	if error != nil {
		panic(error)
	}
	for rows.Next() {
		data := Data{}
		rows.Scan(&data.Id, &data.Temperatura, &data.Humedad, &data.Led1, &data.Pwm, &data.Fecha_hora)
		datos = append(datos, data)
	}
	return datos, error
}
func GetDatos2(sql string, conditional interface{}) (Datos, error) {
	datos := Datos{}
	rows, error := Query(sql, conditional)
	if error != nil {
		panic(error)
	}
	for rows.Next() {
		data := Data{}
		rows.Scan(&data.Id, &data.Temperatura, &data.Humedad, &data.Led1, &data.Pwm, &data.Fecha_hora)
		datos = append(datos, data)
	}
	return datos, error
}
func GetlastDataByUserID(id int64) (Datos, error) {
	sql := "SELECT ID,temperatura,humedad,led1,pwm,fecha_hora  FROM DATA WHERE user_id =?  ORDER BY id DESC  LIMIT 1"
	return GetDatos2(sql, id)
}
func GetDataBydate(inicio string, fin string) (Datos, error) {

	sql := fmt.Sprintf("SELECT ID,temperatura,humedad,led1,pwm,fecha_hora FROM DATA WHERE fecha_hora BETWEEN '%s' AND '%s' ORDER BY ID DESC", inicio, fin)

	return GetDatos(sql)
}

func GetDataByNum(id int) (Datos, error) {

	sql := "SELECT ID,temperatura,humedad,led1,pwm,fecha_hora FROM DATA ORDER BY ID DESC LIMIT ? "

	return GetDatos2(sql, id)
}
