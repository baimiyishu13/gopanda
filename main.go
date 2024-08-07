package main

import (
	handlers2 "github.com/baimiyishu13/gopanda/internal/handlers"
	"github.com/baimiyishu13/gopanda/internal/middleware/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	log := logger.GetLogger()

	r := chi.NewRouter()

	//test
	r.Get("/test", handlers2.TestHandler)

	// mian
	r.Get("/", handlers2.HomeHandler)

	log.Info("starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Errorf("failed to start server: %v", err)
		return
	}
}
