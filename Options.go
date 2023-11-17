package go_gin_graceful_shutdown

import (
	"log/slog"
	"net/http"
	"time"
)

// Options is the struct to configure the server
type Options struct {
	Port        int           // Port to listen to
	Engin       http.Handler  // Engine to use usually gin.New() or gin.Default()
	WaitTimeout time.Duration // WaitTimeout to finish all run requests
	Log         *slog.Logger  // log to use
}
