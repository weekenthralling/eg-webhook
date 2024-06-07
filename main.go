package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"zjuici.com/tablegpt/eg-webhook/api"
	"zjuici.com/tablegpt/eg-webhook/config"
	"zjuici.com/tablegpt/eg-webhook/storage"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()
	// Initialize the database
	db := storage.Init(cfg)

	sessionStore := storage.NewKernelSessionStore(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", api.ListKernels(sessionStore))
	r.Get("/{id}", api.GetSession(sessionStore))
	r.Delete("/", api.DeleteKernels(sessionStore))
	r.Post("/{id}", api.SaveKernel(sessionStore))

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Panic(err)
	}
}
