package dlog

import (
	"io"
	"log"
)

var debug bool = false

type DebugLogger struct {
	*log.Logger
	debugEnabled bool
}

func New(out io.Writer, prefix string, flag int, debugEnabled bool) *DebugLogger {
	return &DebugLogger{log.New(out, prefix, flag), debugEnabled}
}

func (d *DebugLogger) Enable()  { d.debugEnabled = true }
func (d *DebugLogger) Disable() { d.debugEnabled = false }

func (d *DebugLogger) Debug(args ...interface{}) {
	if d.debugEnabled {
		d.Print("[DEBUG]", args)
	}
}

func (d *DebugLogger) Enter(fn string) {
	d.Debug("ENTER:", fn)
}

// TODO: Integrate debug control into the logger instead of having to have ugly if statements directly in the code
// TODO: Extend the logger even further to also have ready generic functionality to log function ENTRY/EXIT
func SetDebug(flag bool) {
	debug = flag
}

func Debug(args ...interface{}) {
	if debug {
		log.Println("[DEBUG]", args)
	}
}

func Enter(fn string) {
	Debug("ENTER:", fn)
}
