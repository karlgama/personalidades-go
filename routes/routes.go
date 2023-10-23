package routes

import (
	"log"
	"net/http"
	"personalidades/controllers"
	"personalidades/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.Edit).Methods("Put")
	r.HandleFunc("/api/personalidades", controllers.Create).Methods("Post")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
