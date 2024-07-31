package main

import (
	"fmt"
	"github.com/baimiyishu13/gopanda/logger"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"os"
)

//type User struct {
//	Name string
//}

func main() {
	fmt.Println("exp main")
	log := logger.GetLogger()

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Errorf("failed to parse template: %v", err)
	}

	user := struct {
		Name string
	}{
		Name: "Gopanda",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Errorf("failed to execute template: %v", err)
	}

	// chi
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, user)
	})
	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		// info.html
		http.ServeFile(w, r, "info.html")
	})
	http.ListenAndServe(":8888", r)
}
