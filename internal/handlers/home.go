package handlers

import (
	"fmt"
	"github.com/baimiyishu13/gopanda/internal/middleware/logger"
	"html/template"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("internal", "templates", "home.gohtml")
	fmt.Println(tplPath)
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Errorf("Failed to parse template: %v", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Errorf("Failed to execute template: %v", err)
		return
	}
}
