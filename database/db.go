package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func DbConnect() {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(conexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao Conectar com o banco de dados")
	}
}
