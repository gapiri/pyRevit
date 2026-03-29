package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
)

// Options holds all parsed command-line options for the Nova server.
type Options struct {
	ExeName string `json:"exe_name"`
	Version string `json:"version"`
	Opts    *docopt.Opts
	Port    int  `json:"server_port"`
	Https   bool `json:"https"`
	Debug   bool `json:"debug_mode"`
	Trace   bool `json:"trace_mode"`
}

func getExeName() string {
	return strings.TrimSuffix(
		filepath.Base(os.Args[0]),
		filepath.Ext(os.Args[0]),
	)
}

// NewOptions parses command-line arguments and returns an Options struct.
func NewOptions() *Options {
	argv := os.Args[1:]

	parser := &docopt.Parser{
		HelpHandler: printHelpAndExit,
	}

	opts, parseErr := parser.ParseArgs(help, argv, version)
	if parseErr != nil {
		fmt.Fprintln(os.Stderr, parseErr)
		os.Exit(1)
	}

	port, _ := opts.Int("--port")
	https, _ := opts.Bool("--https")
	debug, _ := opts.Bool("--debug")
	trace, _ := opts.Bool("--trace")

	return &Options{
		ExeName: getExeName(),
		Version: version,
		Opts:    &opts,
		Port:    port,
		Https:   https,
		Debug:   debug,
		Trace:   trace,
	}
}
