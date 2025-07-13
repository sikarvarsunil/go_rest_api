package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sikarvarsunil/go_rest_api/internal/config"
	"github.com/sikarvarsunil/go_rest_api/internal/handlers/employee"
)

func main() {
	//load config
	cfg := config.MustLoad()
	// database setup

	// storage, err := sqlite.New(cfg)

	// if err != nil {
	// 	log.Fatal((err))
	// }
	fmt.Println("StoragePath =", cfg.StoragePath)
	slog.Info("storage initialized", slog.String("env", cfg.Env))
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/employees", employee.New())
	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server started", slog.String("Address", cfg.Addr))
	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start the Server")

		}
	}()

	<-done
	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown the server", slog.String("error", err.Error()))
	}
	slog.Info("shutting down the server successfully")
}
