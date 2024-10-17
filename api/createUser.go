package api

import (
	"encoding/json"
	"net/http"

	"CrudInMemoryGo/database"
	"CrudInMemoryGo/jsons"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var body database.User
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		jsons.Send(
			writer,
			jsons.Response{Error: "Invalid request body"},
			http.StatusUnprocessableEntity,
		)
	}

	databaseInstance := database.NewDatabase()
	id := databaseInstance.CreateUser(body)

	jsons.Send(writer, jsons.Response{Data: id}, http.StatusCreated)
}
