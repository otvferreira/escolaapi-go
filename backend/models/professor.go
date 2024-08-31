// backend/models/professor.go
package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Nome   string  `gorm:"size:100;not null" json:"nome"`
	Email  string  `gorm:"size:100;unique;not null" json:"email"`
	CPF    string  `gorm:"size:11;unique;not null" json:"cpf"`
	Turmas []Turma `gorm:"foreignKey:ProfessorID" json:"turmas"`
}
