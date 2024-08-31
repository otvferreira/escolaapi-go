package models

import "gorm.io/gorm"

type Turma struct {
	gorm.Model
	Nome        string      `gorm:"size:100;not null" json:"nome"`
	Semestre    string      `gorm:"size:20" json:"semestre"`
	Ano         int         `json:"ano"`
	ProfessorID uint        `json:"professor_id"`
	Atividades  []Atividade `gorm:"foreignKey:TurmaID" json:"atividades"`
}
