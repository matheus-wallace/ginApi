package database

import (
	"ginApi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComOBancoDeDados() {
	stringDeConexao := "host=172.21.0.2 user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o bando de dados", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
