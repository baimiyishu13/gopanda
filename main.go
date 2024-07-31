package main

import (
	"github.com/baimiyishu13/gopanda/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// index.html
	http.ServeFile(w, r, "./static/index.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// contact.html
	http.ServeFile(w, r, "./static/contact.html")
}

// 404
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// 404.html
	http.ServeFile(w, r, "./static/404.html")

}

func main() {
	// 初始化 logrus
	log := logger.GetLogger()
	//chi 路由器
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	// 404
	r.NotFound(notFoundHandler)

	// 启动服务器
	log.Info("starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Errorf("Server failed to start: %v", err)
	}
}
