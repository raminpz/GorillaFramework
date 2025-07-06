package main

import (
	"GorillaFramework/route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	r := mux.NewRouter()

	route.SetupRoutes(r)

	r.Use(LogginMiddleware)

	log.Println("Servidor Gorilla ejecutando en el puerto :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
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
