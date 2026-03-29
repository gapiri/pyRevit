package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pyrevitnovaserver/cli"
	"pyrevitnovaserver/store"

	"github.com/gorilla/mux"
)

// RouteScripts registers the /api/v2/scripts/ POST and GET endpoints.
func RouteScripts(router *mux.Router, opts *cli.Options, st *store.Store, logger *cli.Logger) {
	// POST /api/v2/scripts/ – store a new script execution record
	router.HandleFunc("/api/v2/scripts/", func(w http.ResponseWriter, r *http.Request) {
		var rec store.ScriptRecord
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			logger.Debug(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		st.AddScript(rec)
		logger.Print(fmt.Sprintf(
			"%s %s %q %s [%s.%s] code=%d",
			OkMessage,
			rec.TimeStamp,
			rec.UserName,
			rec.RevitBuild,
			rec.ExtensionName,
			rec.CommandName,
			rec.ResultCode,
		))
		writeJSON(w, rec, logger)
	}).Methods("POST")

	// GET /api/v2/scripts/ – retrieve all stored script execution records
	router.HandleFunc("/api/v2/scripts/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, st.GetScripts(), logger)
	}).Methods("GET")
}
