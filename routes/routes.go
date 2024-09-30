package routes

import (
	"log"
	"net/http"

	"github.com/deadman360/PrimeiraRestAPI-GO-/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.ExibePersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
