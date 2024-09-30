package controllers

import (
	"API/GIN/database"
	"API/GIN/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context){
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "Qual vai ser "+ nome +" ???",
	})
}

func CriaNovoAluno(c *gin.Context){

	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}