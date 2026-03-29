package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pyrevitnovaserver/cli"
	"pyrevitnovaserver/store"

	"github.com/gorilla/mux"
)

// RouteEvents registers the /api/v2/events/ POST and GET endpoints.
func RouteEvents(router *mux.Router, opts *cli.Options, st *store.Store, logger *cli.Logger) {
	// POST /api/v2/events/ – store a new application event record
	router.HandleFunc("/api/v2/events/", func(w http.ResponseWriter, r *http.Request) {
		var rec store.EventRecord
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			logger.Debug(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		st.AddEvent(rec)
		logger.Print(fmt.Sprintf(
			"%s %s [%s] %q @ %s doc=%q",
			OkMessage,
			rec.TimeStamp,
			rec.EventType,
			rec.HostUserName,
			rec.RevitBuild,
			rec.DocumentName,
		))
		writeJSON(w, rec, logger)
	}).Methods("POST")

	// GET /api/v2/events/ – retrieve all stored event records
	router.HandleFunc("/api/v2/events/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, st.GetEvents(), logger)
	}).Methods("GET")
}
