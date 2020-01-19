package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/medis/go-expressive/config"
	expressive "github.com/medis/go-expressive/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
}

func run() error {
	// Load config.
	config := config.Load()

	// Configure gorilla router.
	r := mux.NewRouter()
	expressive.RegisterRoutes(r, config.Routes)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         ":80",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return err

	case sig := <-shutdown:
		log.Printf("%v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("Graceful shutdown did not complete in %v : %v", 30*time.Second, err)
			err = srv.Close()
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			log.Println("could not stop server gracefully")
			return err
		}
	}

	return nil
}
