package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome      string  `gorm:"size:100;not null" json:"nome"`
	Matricula string  `gorm:"size:20;unique;not null" json:"matricula"`
	Turmas    []Turma `gorm:"many2many:aluno_turmas" json:"turmas"`
}
