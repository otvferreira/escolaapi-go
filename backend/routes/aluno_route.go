package routes

import (
	"backend/config"
	"backend/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Função auxiliar para converter []uint para []string
func toStringSlice(ids []uint) []string {
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, fmt.Sprintf("%d", id))
	}
	return strIds
}

// Função auxiliar para converter []string para []uint
func toUintSlice(ids []string) ([]uint, error) {
	var uintIds []uint
	for _, id := range ids {
		var uintId uint
		_, err := fmt.Sscanf(id, "%d", &uintId)
		if err != nil {
			return nil, err
		}
		uintIds = append(uintIds, uintId)
	}
	return uintIds, nil
}

// CreateAluno cria um novo aluno
func CreateAluno(c *gin.Context) {
	var input struct {
		Nome      string `json:"nome"`
		Matricula string `json:"matricula"`
		Turmas    []uint `json:"turmas"` // IDs das turmas
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aluno := models.Aluno{
		Nome:      input.Nome,
		Matricula: input.Matricula,
		TurmaIDs:  strings.Join(toStringSlice(input.Turmas), ","),
	}

	if err := config.DB.Create(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, aluno)
}

// GetAlunos retorna todos os alunos
func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	if err := config.DB.Find(&alunos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preenchendo as entidades Turma para cada aluno
	for i := range alunos {
		turmaIDs := strings.Split(alunos[i].TurmaIDs, ",")
		uintIDs, err := toUintSlice(turmaIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter IDs de turmas"})
			return
		}

		var turmas []models.Turma
		if err := config.DB.Where("id IN ?", uintIDs).Find(&turmas).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
			return
		}

		alunos[i].TurmaIDs = "" // Limpando IDs após adicionar entidades
	}

	c.JSON(http.StatusOK, gin.H{"data": alunos})
}

// GetAluno retorna um aluno específico com suas turmas associadas
func GetAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	// Buscando turmas associadas
	turmaIDs := strings.Split(aluno.TurmaIDs, ",")
	uintIDs, err := toUintSlice(turmaIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter IDs de turmas"})
		return
	}

	var turmas []models.Turma
	if err := config.DB.Where("id IN ?", uintIDs).Find(&turmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
		return
	}

	// Enviando o aluno e suas turmas associadas como resposta
	c.JSON(http.StatusOK, gin.H{
		"aluno":  aluno,
		"turmas": turmas,
	})
}

// UpdateAluno atualiza os dados de um aluno específico
func UpdateAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	var input struct {
		Nome      string `json:"nome"`
		Matricula string `json:"matricula"`
		Turmas    []uint `json:"turmas"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualizando dados do aluno
	aluno.Nome = input.Nome
	aluno.Matricula = input.Matricula
	aluno.TurmaIDs = strings.Join(toStringSlice(input.Turmas), ",")

	if err := config.DB.Save(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeleteAluno remove um aluno específico
func DeleteAluno(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Aluno{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
