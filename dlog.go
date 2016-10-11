package dlog

import (
	"io"
	"log"
)

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
