package routes

import (
	"api/controller"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api/personalidades", controller.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
