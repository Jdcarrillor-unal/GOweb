package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int                 `json:"status"`
	Data        interface{}         `json:"data"`
	Message     string              `json:"message"`
	contentType string              /// privado
	writer      http.ResponseWriter /// al estar en minuscual son privados
}

func CreateDeafaultMessage(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}
func SendNotFound(w http.ResponseWriter) {
	response := CreateDeafaultMessage(w)
	response.NotFound("Resource Not Found")
	response.Send()
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDeafaultMessage(w)
	response.Data = data
	response.Send()
}
func (this *Response) NotFound(message string) {
	this.Status = http.StatusNotFound
	this.Message = message
}
func (this *Response) Send() {
	this.writer.Header().Set("Content-Type", this.contentType)
	this.writer.WriteHeader(this.Status)
	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output)+"\n")
}
func SendNoContent(w http.ResponseWriter) {
	response := CreateDeafaultMessage(w)
	response.NoContent()
	response.Send()
}
func (this *Response) NoContent() {
	this.Status = http.StatusNoContent
	this.Message = "No content mi amigo "
}
func SendUnprocessedEntity(w http.ResponseWriter) {
	response := CreateDeafaultMessage(w)
	response.UnprocessedEntity(w)
	response.Send()
}
func (this *Response) UnprocessedEntity(w http.ResponseWriter) {
	this.Status = http.StatusUnprocessableEntity
	this.Message = "UnprocessableEntity friendo! "
}
