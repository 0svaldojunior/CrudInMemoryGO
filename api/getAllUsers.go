package api

import (
	"net/http"

	"CrudInMemoryGo/database"
	"CrudInMemoryGo/jsons"
)

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	databaseInstance := database.NewDatabase()
	users := databaseInstance.GetAllUsers()

	jsons.Send(writer, jsons.Response{Data: users}, http.StatusOK)
}
