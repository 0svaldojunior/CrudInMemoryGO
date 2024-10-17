package jsons

import (
	"CrudInMemoryGo/log"
	"encoding/json"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Send(writer http.ResponseWriter, response Response, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(response)
	if err != nil {
		log.Error("Failed to marshal response data", err)
		Send(writer, Response{Error: "Internal server error"}, http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(statusCode)
	if _, err := writer.Write(data); err != nil {
		log.Error("Failed to write json response data to client", err)
		return
	}
}
