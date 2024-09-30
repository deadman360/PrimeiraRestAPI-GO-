package main

import (
	"fmt"

	"github.com/deadman360/PrimeiraRestAPI-GO-/database"
	"github.com/deadman360/PrimeiraRestAPI-GO-/routes"
)

func main() {

	database.DbConnect()
	fmt.Println("Iniciando API")
	routes.HandleRequest()

}
