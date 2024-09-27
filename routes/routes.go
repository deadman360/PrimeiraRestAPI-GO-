package routes

import (
	"net/http"

	"github.com/deadman360/PrimeiraRestAPI-GO-/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.ExibePersonalidades)
}
