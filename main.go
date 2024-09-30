package main

import (
	"fmt"

	"github.com/deadman360/PrimeiraRestAPI-GO-/models"
	"github.com/deadman360/PrimeiraRestAPI-GO-/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Siegmound Freud", Descricao: "Ele explica"},
		{Id: 2, Nome: "Paul McCartney", Descricao: "Veio safado que canta muito"},
		{Id: 3, Nome: "Ronald McDonald", Descricao: "Fez o legend sabe mto"},
	}
	fmt.Println("Iniciando API")
	routes.HandleRequest()

}
