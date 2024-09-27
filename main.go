package main

import (
	"API/GIN/models"
	"API/GIN/routes"
)


func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Alta Barton", CPF: "000.111.222-33", RG: "4293949803"},
		{Nome: "Bettie Vargas", CPF: "111.222.333-44", RG: "1977794606"},
	}
	routes.HandleRequest()
}
