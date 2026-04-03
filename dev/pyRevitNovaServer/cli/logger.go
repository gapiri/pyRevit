package cli

import (
	"log"

	"pkg.re/essentialkaos/ek.v10/fmtc"
)

// Logger provides leveled logging for the Nova server.
type Logger struct {
	PrintDebug bool
	PrintTrace bool
}

// NewLogger creates a new Logger based on the provided Options.
func NewLogger(options *Options) *Logger {
	return &Logger{
		PrintDebug: options.Debug,
		PrintTrace: options.Trace,
	}
}

func (m *Logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func (m *Logger) Debug(args ...interface{}) {
	if m.PrintDebug {
		log.Print(args...)
	}
}

func (m *Logger) Trace(args ...interface{}) {
	if m.PrintTrace {
		log.Print(args...)
	}
}

func (m *Logger) Print(args ...interface{}) {
	fmtc.Println(args...)
}
