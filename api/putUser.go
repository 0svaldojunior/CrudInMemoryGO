package api

import (
	"encoding/json"
	"net/http"

	"CrudInMemoryGo/database"
	"CrudInMemoryGo/jsons"
	"CrudInMemoryGo/log"

	"github.com/go-chi/chi/v5"
)

func PutUser(writer http.ResponseWriter, request *http.Request) {
	var body database.User
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		jsons.Send(
			writer,
			jsons.Response{Error: "Invalid request body"},
			http.StatusUnprocessableEntity,
		)
	}

	id := chi.URLParam(request, "id")
	databaseInstance := database.NewDatabase()
	ok := databaseInstance.UpdateUser(id, body)

	if !ok {
		jsons.Send(
			writer,
			jsons.Response{Error: "Someting error"},
			http.StatusBadRequest,
		)
	}

	log.Info("Success PUT user with ID: " + id)
}
