package go_gin_graceful_shutdown

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Serve starts the server and listen to OS signals to init graceful shutdown
//
//	if opt.log is nil, it will use slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
//		Level: slog.LevelInfo,
//	}))
//
// Options is the struct to configure the server
func Serve(opt *Options) {

	opt = loadOptionsDefaults(opt)
	//  starts the server with the options provided
	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", opt.Port), Handler: opt.Engin}

	//  readyToClose is a channel to notify when the server is ready to closed
	readyToClose := make(chan bool)

	//  listenToOSSignalsAndInitShutdown listen to OS signals and init graceful shutdown
	//  it will wait for opt.WaitTimeout to finish all run requests
	//  if opt.WaitTimeout is 0, it will use 20 * time.Second
	go listenToOSSignalsAndInitShutdown(opt.WaitTimeout, httpServer, opt.log, readyToClose)

	opt.log.Info("Server started and ready to accept requests on port", "port", opt.Port)

	// ListenAndServe starts the server and listen to port
	err := httpServer.ListenAndServe()

	if err != nil {
		if !errors.As(err, &http.ErrServerClosed) {
			opt.log.Error(err.Error())
		}
	}

	<-readyToClose
	opt.log.Info("bye bye")
}

// loadOptionsDefaults loads the default values for the options
// if opt.Port is 0, it will use 8080
// if opt.WaitTimeout is 0, it will use 20 * time.Second
//
//	if opt.log is nil, it will use slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelInfo,
//		}))
//
// if opt.Engin is nil, it will panic
// if opt.Engin is not nil, it will use opt.Engine
// if opt.log is not nil, it will use opt.log
// if opt.Port is not nil, it will use opt.Port
// if opt.WaitTimeout is not nil, it will use opt.WaitTimeout
func loadOptionsDefaults(opt *Options) *Options {
	if opt.Engin == nil {
		panic("gin engine is required")
	}
	if opt.Port == 0 {
		opt.Port = 8080
	}

	if opt.WaitTimeout == 0 {
		opt.WaitTimeout = 20 * time.Second
	}

	if opt.log == nil {
		opt.log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return opt
}
