package models

import "gorm.io/gorm"

type Atividade struct {
	gorm.Model
	TurmaID uint    `json:"turma_id"`
	Valor   float64 `json:"valor"`
	Data    string  `json:"data"`
	Notas   []Nota  `gorm:"foreignKey:AtividadeID" json:"notas"`
}
