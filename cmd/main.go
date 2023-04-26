package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/drakondarquesse/pet-api-gateway/pkg/pet"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// create a chi mux
	sm := chi.NewRouter()
	sm.Use(middleware.RequestID)
	sm.Use(middleware.Logger)
	sm.Use(middleware.Recoverer)

	l := log.New(os.Stdout, "api", log.LstdFlags)

	sm.Route("/pets", func(r chi.Router) {
		pet.MountRoutes(r)
	})

	// create a server
	s := &http.Server{
		Addr:        ":8000",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
	}

	// a goroutine to run the server so it doesn't block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create a channel to receive os.Signal
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
