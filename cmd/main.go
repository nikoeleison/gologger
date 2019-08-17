//nikoeleison
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/nikoeleison/gologger"
)

//>>>public

//>>>private

var (
	sm = &sync.Mutex{}
	i  = 0
)

//setup logger
//setup handler
//setup server
//setup graceful shutdown
func main() {
	logger := gologger.New(
		"./log",
		"benchmark_apache",
	)
	defer logger.Kill()

	handler := http.NewServeMux()
	handler.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		sm.Lock()
		defer sm.Unlock()

		i++

		logger.Info("print ", i)
		logger.Infof("print:%d", i)

		logger.Error("print ", i)
		logger.Errorf("print:%d", i)

		logger.Debug("print ", i)
		logger.Debugf("print:%d", i)

		logger.Panic("print ", i)
		logger.Panicf("print:%d", i)

		w.Write([]byte(fmt.Sprintf("print %d\n", i)))
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	go func() {
		kill := make(chan os.Signal, 1)
		signal.Notify(kill, os.Interrupt, os.Kill)

		<-kill

		logger.Info("caught signal interrupt")

		close(kill)
		signal.Stop(kill)

		server.SetKeepAlivesEnabled(false)
		err := server.Shutdown(context.Background())
		if err != nil {
			logger.Errorf("unable to stop server, reason: %v", err)
		}
	}()

	logger.Infof("server started on %s", server.Addr)
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		logger.Errorf("unable to start server, reason: %v", err)
	} else {
		logger.Info("server stopped!")
	}
}
