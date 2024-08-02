package handlers

import (
	"github.com/baimiyishu13/gopanda/middleware/logger"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("./templates/login.gohtml")
	if err != nil {
		log.Errorf("failed to parse template: %v", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Errorf("failed to execute template: %v", err)
	}
}
