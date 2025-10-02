package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/client"
	"github.com/YourCurseSheyme/go_homework_2025/hw2/server"
	"github.com/YourCurseSheyme/go_homework_2025/hw2/server/config"
)

func main() {
	cfg := config.CreateConfig()
	app := server.NewHttpServer(cfg)
	errCh := app.Start()
	fmt.Printf("http: listening on %s\n", app.Addr())

	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		client.Demo()
		shCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		if err := app.Stop(shCtx); err != nil {
			fmt.Printf("http: server stop error after demo: %v\n", err)
		} else {
			fmt.Printf("http: server stop successfully after demo\n")
		}
	}()

	select {
	case <-sigCtx.Done():
		shCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		if err := app.Stop(shCtx); err != nil {
			fmt.Printf("http: stop error signal: %v\n", err)
		} else {
			fmt.Printf("http: stop success signal\n")
		}
	case err := <-errCh:
		if err != nil {
			_ = fmt.Errorf("http: listen %v\n", err)
		}
	}

	fmt.Println("http: server stopped")
}
