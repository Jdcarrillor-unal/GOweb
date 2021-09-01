package test

import (
	"api/models"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	beforeTest()
	result := m.Run()
	afterTest()
	os.Exit(result)
}

func beforeTest() {
	models.CreateConnection()
	models.CreateTables()
	createDeafultUser()
}
func afterTest() {
	models.CloseConnection()
}
func createDeafultUser() {
	sql := fmt.Sprintf("INSERT users SET id='%d',username='%s',password='%s', email='%s',created_date='%s'", id, username, passwordHash, email, createdDate)
	_, err := models.Exec(sql)
	if err != nil {
		panic(err)
	}
	user = &models.User{Id: id, Username: username, Password: password, Email: email}
}
