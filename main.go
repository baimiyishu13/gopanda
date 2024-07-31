package main

import (
	"github.com/baimiyishu13/gopanda/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	// index.html
	http.ServeFile(w, r, "index.html")
}

func main() {
	// 初始化 logrus
	log := logger.GetLogger()
	// chi 路由器
	r := chi.NewRouter()
	r.Get("/", handlerfunc)

	// 启动服务器
	log.Info("starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Errorf("Server failed to start: %v", err)
	}
}
