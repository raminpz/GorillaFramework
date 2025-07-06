package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	// Define the root route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenido al servideo Gorilla!\n"))
	})

	// Define a dynamic route with a variable
	r.HandleFunc("/saludo/{nombre}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nombre := vars["nombre"]
		w.Write([]byte("Hola, " + nombre + "!"))
	})

	//

}
