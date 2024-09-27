package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deadman360/PrimeiraRestAPI-GO-/models"
	"github.com/deadman360/PrimeiraRestAPI-GO-/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Siegmound Freud", Descricao: "Ele explica"},
		{Nome: "Paul McCartney", Descricao: "Veio safado que canta muito"},
		{Nome: "Ronald McDonald", Descricao: "Fez o legend sabe mto"},
	}
	fmt.Println("Iniciando API")
	routes.HandleRequest()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
