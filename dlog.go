package dlog

import (
	"log"
)

var debug bool = false

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
