package handlers

import (
	"github.com/baimiyishu13/gopanda/logger"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	info := struct {
		Name string
	}{
		Name: "Gopanda",
	}
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("./templates/home.gohtml")
	if err != nil {
		log.Errorf("failed to parse template: %v", err)
	}
	err = t.Execute(w, info)
	if err != nil {
		log.Errorf("failed to execute template: %v", err)
	}
}
