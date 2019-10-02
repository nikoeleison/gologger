// @author nikoeleison
package gologger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// @public

// driver struct
// pool section
// writer section
type Driver struct {
	pool chan *message
	kill chan bool
	sm   sync.Mutex
	swg  sync.WaitGroup

	path     string
	prefix   string
	suffix   string
	filepath string
	file     *os.File
	writer   io.Writer
}

// driver constructor
// dispatch pool
// rotate writer file
func New(path string, prefix string) (d *Driver) {
	d = &Driver{}
	d.pool = make(chan *message, 100)
	d.kill = make(chan bool, 1)
	d.dispatch()

	d.path = path
	d.prefix = prefix
	d.rotate(
		time.Now().Format(suffixformat),
	)

	return
}

// kill pool
// consume the rest of message pool
func (d *Driver) Kill() {
	d.sm.Lock()
	defer d.sm.Unlock()

	d.kill <- true

	for len(d.pool) > 0 {
		msg := <-d.pool
		d.consume(msg)
	}

	d.swg.Wait()
}

// @private

var (
	suffixformat = "2006-01-02"
)

// produce new message
// deliver message to consumer
func (d *Driver) produce(level string, s string) {
	d.sm.Lock()
	defer d.sm.Unlock()

	_, file, line, _ := runtime.Caller(2)

	d.swg.Add(1)
	go func() {
		msg := newMsg(
			time.Now(),
			level,
			file,
			line,
			s,
		)
		d.pool <- msg
	}()
}

// consume message
// rotate writer file
// release message deliver
func (d *Driver) consume(msg *message) {
	d.rotate(
		msg.now.Format(suffixformat),
	)
	fmt.Fprint(d.writer, msg.decorate())
	d.swg.Done()
}

// rotate writer file
// skip rotate if driver suffix equal to now suffix
// mkdir if not exist
// touch if not exist
func (d *Driver) rotate(suffix string) (err error) {
	if d.suffix == suffix {
		return nil
	}

	d.suffix = suffix

	d.filepath = fmt.Sprintf(
		"%s.log.%s",
		d.prefix,
		d.suffix,
	)
	d.filepath = filepath.Join(d.path, d.filepath)

	err = os.MkdirAll(
		d.path,
		os.ModePerm,
	)
	if err != nil && !os.IsExist(err) {
		return
	}

	d.file, err = os.OpenFile(
		d.filepath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		return
	}

	d.writer = io.MultiWriter(os.Stdout, d.file)

	return nil
}

// pool dispatcher
// consume message
// break infinite loop
func (d *Driver) dispatch() {
	go func() {
		for {
			select {
			case msg := <-d.pool:
				d.consume(msg)
			case <-d.kill:
				break
			default:
			}
		}
	}()
}
