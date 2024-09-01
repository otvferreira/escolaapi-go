package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome      string  `json:"nome"`
	Matricula string  `json:"matricula"`
	TurmaIDs  string  `json:"turmas"` // IDs das turmas como string separada por vírgulas
	Turmas    []Turma `gorm:"many2many:aluno_turmas;" json:"-"`
}
