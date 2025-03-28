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
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", r))
}
