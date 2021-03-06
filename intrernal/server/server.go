package server

import (
	"context"
	"net/http"
	"time"
)

type Bootstrap struct {
	httpServer *http.Server
}


func (s *Bootstrap) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Bootstrap) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

