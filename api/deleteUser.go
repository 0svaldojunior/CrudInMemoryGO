package api

import (
	"net/http"

	"CrudInMemoryGo/database"
	"CrudInMemoryGo/jsons"

	"github.com/go-chi/chi/v5"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	databaseInstance := database.NewDatabase()
	ok := databaseInstance.DeleteUser(id)

	if !ok {
		jsons.Send(
			writer,
			jsons.Response{Error: "User not found"},
			http.StatusNotFound,
		)
	}

	jsons.Send(
		writer,
		jsons.Response{Data: "Delete user ID: " + id},
		http.StatusOK,
	)
}
