package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"Nome"`
	CPF  string `json:"CPF"`
	RG   string `json:"RG"`
}

var Alunos []Aluno
