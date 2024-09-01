package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome      string `gorm:"size:100;not null" json:"nome"`
	Matricula string `gorm:"size:20;unique;not null" json:"matricula"`
	Turmas    []uint `gorm:"-" json:"turmas"` // IDs das turmas, ignorado pelo GORM
}
