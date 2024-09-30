package database

import (
	"API/GIN/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConectaComDB() *gorm.DB{
	strConexao := "host=localhost user=postgres password=postgres dbname=api_golang_gin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConexao))
	if err != nil{
		log.Panic("Erro ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{}) //o gorm cria uma tabela no BD de acordo com os parametros da struct alem de outros parametros criados automaticamente
	return DB
}