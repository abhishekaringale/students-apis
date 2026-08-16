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

	"github.com/abhishekaringale/students-api/internal/config"
	"github.com/abhishekaringale/students-api/internal/http/handlers/student"
	"github.com/abhishekaringale/students-api/storage/sqlite"
)

func main() {
	fmt.Println("welcome to go")
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))


	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	// Print that the server is starting BEFORE calling the blocking ListenAndServe
	slog.Info("server started", slog.String("address", cfg.Addr))

	// Create a channel to listen for OS signals (like Ctrl+C)
	done := make(chan os.Signal, 1)

	// Register the channel to receive Interrupt (SIGINT) and Kill (SIGTERM) signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine (background thread) so it doesn't block the main thread
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server")
		}
	}()

	// Block (pause) the main thread here until we receive an OS signal on the 'done' channel (like ctrl + c)
	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server..", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
