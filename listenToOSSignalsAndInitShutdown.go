package go_gin_graceful_shutdown

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// listenToOSSignalsAndInitShutdown listen to OS signals and init graceful shutdown
// it will wait for timeout to finish all run requests
// if timeout is 0, it will use 20 * time.Second
func listenToOSSignalsAndInitShutdown(timeout time.Duration, httpServer *http.Server, log *slog.Logger, closed chan bool) {
	// identify the signals to listen to
	// SIGINT: interrupt signal, sent from CTRL+C or KILL
	// SIGTERM: termination signal, sent from kill
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, os.Kill)
	log.Info("server waiting for signal to shutdown")
	// wait for the signal
	s := <-sig
	close(sig)
	log.Info("server received signal", "signal", s.String())
	log.Info("server waiting to finish all run requests or reach timeout of ", "timeout", timeout.String())
	// create a context with timeout
	// if timeout is 0, it will use 20 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// shutdown the server
	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Error("server close with error ", "error", err.Error())
	}
	// 	notify that the server is ready to close
	closed <- true

}
