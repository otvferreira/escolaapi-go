package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome      string `gorm:"size:100;not null" json:"nome"`
	Matricula string `gorm:"size:50;not null" json:"matricula"`
	TurmaIDs  string `gorm:"size:255" json:"turmas"` // Armazena IDs das turmas como uma string
}
