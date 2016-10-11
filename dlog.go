package dlog

import (
	"fmt"
	"io"
	"log"
)

type DebugLogger struct {
	*log.Logger
	debugEnabled bool
}

func New(out io.Writer, prefix string, flag int) *DebugLogger {
	return &DebugLogger{log.New(out, prefix, flag), false}
}

func (d *DebugLogger) EnableDebug()         { d.debugEnabled = true }
func (d *DebugLogger) DisableDebug()        { d.debugEnabled = false }
func (d *DebugLogger) IsDebugEnabled() bool { return d.debugEnabled }

func (d *DebugLogger) Debug(v ...interface{}) {
	if d.debugEnabled {
		d.Output(2, "[DEBUG] "+fmt.Sprint(v...))
	}
}

func (d *DebugLogger) Debugln(v ...interface{}) {
	if d.debugEnabled {
		d.Output(2, "[DEBUG] "+fmt.Sprintln(v...))
	}
}

func (d *DebugLogger) Debugf(format string, v ...interface{}) {
	if d.debugEnabled {
		d.Output(2, "[DEBUG] "+fmt.Sprintf(format, v...))
	}
}

func (d *DebugLogger) Enter(fn string) {
	d.Debugln("ENTER: ", fn)
}
