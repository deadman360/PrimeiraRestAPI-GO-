package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deadman360/PrimeiraRestAPI-GO-/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, RestAPI")
}

func ExibePersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
