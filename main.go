package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"zjuici.com/tablegpt/eg-webhook/api"
	"zjuici.com/tablegpt/eg-webhook/config"
	db "zjuici.com/tablegpt/eg-webhook/storage"
)

func main() {

	// Load configuration
	config.LoadConfig()
	// Initialize the database
	db.Init()

	sessionStore := db.NewKernelSessionStore(db.GetDB())

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", api.ListKernels(sessionStore))
	r.Get("/{id}", api.GetSession(sessionStore))
	r.Delete("/", api.DeleteKernels(sessionStore))
	r.Post("/{id}", api.SaveKernel(sessionStore))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic(err)
	}
}
