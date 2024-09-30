package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model //adicione este 'tipo' do GORM para ser criar essa struct no banco de dados
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno
