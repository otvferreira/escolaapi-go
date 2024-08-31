package models

import "gorm.io/gorm"

type Nota struct {
	gorm.Model
	AlunoID     uint    `json:"aluno_id"`
	AtividadeID uint    `json:"atividade_id"`
	Nota        float64 `json:"nota"`
}
