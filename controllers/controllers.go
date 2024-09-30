package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deadman360/PrimeiraRestAPI-GO-/database"
	"github.com/deadman360/PrimeiraRestAPI-GO-/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, RestAPI")
}

func ExibePersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.Db.Find(&p)
	json.NewEncoder(w).Encode(p)
}
func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.Db.First(&p, id)
	json.NewEncoder(w).Encode(p)

}
func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	database.Db.Create(&p)
	json.NewEncoder(w).Encode(p)
}
func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.Db.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.Db.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.Db.Save(&p)
	json.NewEncoder(w).Encode(p)

}
