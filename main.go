package main

import (
	"GorillaFramework/route"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", map[string]interface{}{
			"Title":   "Gorilla Framework",
			"Heading": "Welcome to Gorilla Framework",
			"Message": "Bienvenido al curso de Go con Gorilla Framework",
		})
	})
	route.SetupRoutes(r)

	r.Use(LogginMiddleware)

	log.Println("Servidor Gorilla ejecutando en el puerto :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("template/" + tmpl)
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}

}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(start)
		log.Printf("Request: %s %s, Latency: %v, Status: %s\n", r.Method, r.URL.Path, latency, http.StatusText(http.StatusOK))
	})
}
