package server

import (
	"encoding/json"
	"net/http"

	"pyrevitnovaserver/cli"
)

// writeJSON serialises data as JSON and writes it to the response writer.
func writeJSON(w http.ResponseWriter, data interface{}, logger *cli.Logger) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, writeErr := w.Write(jsonData); writeErr != nil {
		logger.Debug(writeErr)
	}
}
