package main

import (
	"fmt"

	"pyrevitnovaserver/cli"
	"pyrevitnovaserver/server"
	"pyrevitnovaserver/store"
)

func main() {
	// parse command-line arguments
	options := cli.NewOptions()

	logger := cli.NewLogger(options)
	logger.Print(fmt.Sprintf("Starting pyRevit Nova Server v%s", options.Version))
	logger.Trace(options)

	// create in-memory store
	st := store.NewStore()

	// start HTTP server
	server.Start(options, st, logger)
}
