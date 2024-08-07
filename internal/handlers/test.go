package handlers

import (
	"github.com/baimiyishu13/gopanda/internal/middleware/logger"
	"html/template"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("./internal/templates/test/test.gohtml")
	if err != nil {
		log.Errorf("failed to parse template: %v", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Errorf("failed to execute template: %v", err)
	}
}
