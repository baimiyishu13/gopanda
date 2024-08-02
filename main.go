package main

import (
	handler "github.com/baimiyishu13/gopanda/cmd/k8s-pod-info-exporter"
	"github.com/baimiyishu13/gopanda/handlers"
	"github.com/baimiyishu13/gopanda/middleware"
	"github.com/baimiyishu13/gopanda/middleware/logger"
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

	// k8s-pod-info-exporter
	r.Get("/k8s-pod-info-exporter", handlers.K8sPodInfoExporter)
	r.Post("/api/upload", handler.UploadHandler)               // 文件上传的路由
	r.Get("/api/download/{filename}", handler.DownloadHandler) // 文件下载的路由
	r.NotFound(handlers.NotFoundHandler)

	log.Info("starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Errorf("Server failed to start: %v", err)
	}
}
