package graceful

import (
	go_gin_graceful_shutdown "github.com/joaziz/go-gin-graceful-shutdown"
	"net/http"
	"time"
)

type Graceful struct {
	router   http.Handler
	duration time.Duration
}

func (receiver *Graceful) ListenAndServe(addr string) {
	go_gin_graceful_shutdown.Serve(&go_gin_graceful_shutdown.Options{
		Engin:       receiver.router,
		WaitTimeout: receiver.duration,
		Addr:        addr,
	})
}

func New(router http.Handler, duration time.Duration) *Graceful {
	return &Graceful{router, duration}
}
