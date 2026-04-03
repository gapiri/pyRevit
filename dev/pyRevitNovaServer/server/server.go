package server

import (
	"fmt"
	"net/http"

	"pyrevitnovaserver/cli"
	"pyrevitnovaserver/store"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

var serverID uuid.UUID

// OkMessage is the formatted success indicator printed to the console.
var OkMessage = "[ {g}OK{!} ]"

// Start initialises the router, registers all routes, and begins listening.
func Start(opts *cli.Options, st *store.Store, logger *cli.Logger) {
	serverID = uuid.Must(uuid.NewV4())

	router := mux.NewRouter().StrictSlash(true)

	RouteScripts(router, opts, st, logger)
	RouteEvents(router, opts, st, logger)
	RouteStatus(router, opts, st, logger)

	logger.Print(fmt.Sprintf("Nova server listening on port %d...", opts.Port))
	addr := fmt.Sprintf(":%d", opts.Port)

	if opts.Https {
		logger.Fatal(
			http.ListenAndServeTLS(
				addr,
				fmt.Sprintf("%s.crt", opts.ExeName),
				fmt.Sprintf("%s.key", opts.ExeName),
				router,
			))
	} else {
		logger.Fatal(http.ListenAndServe(addr, router))
	}
}
