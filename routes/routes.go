package routes

import (
	"api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controller.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controller.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controller.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
