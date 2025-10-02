package server

import (
	"context"
	"net/http"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/server/config"
	endpoints2 "github.com/YourCurseSheyme/go_homework_2025/hw2/server/endpoints"
)

type HttpServer struct {
	server *http.Server
	mux    *http.ServeMux
	config config.Config
}

func NewHttpServer(config config.Config) *HttpServer {
	mux := http.NewServeMux()
	mux.Handle("/version", endpoints2.VersionEndpoint())
	mux.Handle("/decode", endpoints2.DecodeEndpoint())
	mux.Handle("/hard-op", endpoints2.HardOpEndpoint(config.HardMinSleep, config.HardMaxSleep, config.HardErrPct))

	server := &http.Server{
		Addr:         config.Addr,
		Handler:      mux,
		ReadTimeout:  config.IOTimeout,
		WriteTimeout: config.IOTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
	return &HttpServer{
		server: server,
		mux:    mux,
		config: config,
	}
}

func (h *HttpServer) Start() <-chan error {
	errCh := make(chan error, 1)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			errCh <- err
		}
		close(errCh)
	}()
	return errCh
}

func (h *HttpServer) Stop(parent context.Context) error {
	ctx, cancel := context.WithTimeout(parent, h.config.ShutdownTimeout)
	defer cancel()
	return h.server.Shutdown(ctx)
}

func (h *HttpServer) Addr() string { return h.server.Addr }
