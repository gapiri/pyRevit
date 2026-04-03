package cli

import (
	"fmt"
	"os"
)

const version string = "1.0"
const help string = `pyRevit Nova - In-memory data server for pyRevit

Usage:
	pyrevit-novaserver --port=<port> [--https] [--debug] [--trace]

Options:
	-h --help        show this screen
	-V --version     show version
	--port=<port>    server port number to listen on
	--https          secure connection, expects ./pyrevit-novaserver.key and ./pyrevit-novaserver.crt
	--debug          print debug info
	--trace          print trace info e.g. full json payloads

Examples:
	pyrevit-novaserver --port=8090
	pyrevit-novaserver --port=8090 --debug
	pyrevit-novaserver --port=8090 --https --trace`

var printHelpAndExit = func(err error, docoptMessage string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	} else {
		fmt.Println(docoptMessage)
		os.Exit(0)
	}
}
