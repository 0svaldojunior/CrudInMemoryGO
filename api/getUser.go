package api

import (
	"net/http"

	"CrudInMemoryGo/database"
	"CrudInMemoryGo/jsons"

	"github.com/go-chi/chi/v5"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	databaseInstance := database.NewDatabase()
	user, ok := databaseInstance.GetUser(id)

	if !ok {
		jsons.Send(
			writer,
			jsons.Response{Error: "User not found"},
			http.StatusNotFound,
		)
	}

	jsons.Send(writer, jsons.Response{Data: user}, http.StatusOK)
}
