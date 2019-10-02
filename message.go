// @author nikoeleison
package gologger

import (
	"fmt"
	"time"
)

// @public

// @private

// message struct
type message struct {
	now   time.Time
	level string
	file  string
	line  int
	s     string
}

// message constructor
func newMsg(now time.Time, level string, file string, line int, s string) (msg *message) {
	if line == 0 {
		file = "???"
	}

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}

	msg = &message{}
	msg.now = now
	msg.level = level
	msg.file = file
	msg.line = line
	msg.s = s

	return msg
}

// decorate message
func (msg *message) decorate() string {
	return fmt.Sprintf(
		"%s [%5s] %-30s - %s\n",
		msg.now.Format("2006-01-02 15:04:05.000"),
		msg.level,
		fmt.Sprintf("%s:%d", msg.file, msg.line),
		msg.s,
	)
}
