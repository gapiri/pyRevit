package server

import (
	"encoding/json"
	"net/http"

	"pyrevitnovaserver/cli"
	"pyrevitnovaserver/store"

	"github.com/gorilla/mux"
)

// ServerStatus is the response body for the health-check endpoint.
// It follows the RFC Health Check Response Format for HTTP APIs.
// See https://inadarei.github.io/rfc-healthcheck/
type ServerStatus struct {
	Status    string         `json:"status"`
	Version   string         `json:"version"`
	ServiceId string         `json:"serviceid"`
	Stats     map[string]int `json:"stats"`
}

// RouteStatus registers the /api/v2/status GET endpoint.
func RouteStatus(router *mux.Router, opts *cli.Options, st *store.Store, logger *cli.Logger) {
	router.HandleFunc("/api/v2/status", func(w http.ResponseWriter, r *http.Request) {
		resp := ServerStatus{
			Status:    "pass",
			Version:   opts.Version,
			ServiceId: serverID.String(),
			Stats:     st.Stats(),
		}
		jsonData, err := json.Marshal(resp)
		if err != nil {
			logger.Debug(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/health+json")
		if _, writeErr := w.Write(jsonData); writeErr != nil {
			logger.Debug(writeErr)
		}
	}).Methods("GET")
}
