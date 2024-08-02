package main

import (
	"github.com/baimiyishu13/gopanda/handlers"
	"github.com/baimiyishu13/gopanda/logger"
	"github.com/baimiyishu13/gopanda/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	log := logger.GetLogger()

	r := chi.NewRouter()
	middleware.StaticFileServer(r, "/static", http.Dir("static"))
	r.Get("/", handlers.HomeHandler)
	r.Get("/contact", handlers.ContactHandler)
	r.Get("/login", handlers.LoginHandler)
	r.NotFound(handlers.NotFoundHandler)

	log.Info("starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Errorf("Server failed to start: %v", err)
	}
}
