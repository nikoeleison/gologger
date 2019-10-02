// @author nikoeleison
package gologger

import "fmt"

// @public

// print
func (d *Driver) Info(any ...interface{}) {
	d.produce(infolevel, fmt.Sprint(any...))
}

// printf
func (d *Driver) Infof(format string, any ...interface{}) {
	d.produce(infolevel, fmt.Sprintf(format, any...))
}

// print
func (d *Driver) Error(any ...interface{}) {
	d.produce(errorlevel, fmt.Sprint(any...))
}

// printf
func (d *Driver) Errorf(format string, any ...interface{}) {
	d.produce(errorlevel, fmt.Sprintf(format, any...))
}

// print
func (d *Driver) Debug(any ...interface{}) {
	d.produce(debuglevel, fmt.Sprint(any...))
}

// printf
func (d *Driver) Debugf(format string, any ...interface{}) {
	d.produce(debuglevel, fmt.Sprintf(format, any...))
}

// print
func (d *Driver) Panic(any ...interface{}) {
	d.produce(paniclevel, fmt.Sprint(any...))
}

// printf
func (d *Driver) Panicf(format string, any ...interface{}) {
	d.produce(paniclevel, fmt.Sprintf(format, any...))
}

// @private

var (
	infolevel  = "INFO"
	errorlevel = "ERROR"
	debuglevel = "DEBUG"
	paniclevel = "PANIC"
)
