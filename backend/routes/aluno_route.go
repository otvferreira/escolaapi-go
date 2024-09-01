package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	// Convertendo IDs de turmas para entidades Turma
	var turmas []models.Turma
	if len(input.Turmas) > 0 {
		if err := config.DB.Where("id IN ?", input.Turmas).Find(&turmas).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
			return
		}
	}

	aluno := models.Aluno{
		Nome:      input.Nome,
		Matricula: input.Matricula,
		Turmas:    input.Turmas,
	}

	if err := config.DB.Create(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, aluno)
}

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	if err := config.DB.Find(&alunos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preenchendo as entidades Turma para cada aluno
	for i := range alunos {
		var turmas []models.Turma
		if err := config.DB.Where("id IN ?", alunos[i].Turmas).Find(&turmas).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
			return
		}
		alunos[i].Turmas = nil // Limpando IDs após adicionar entidades
	}

	c.JSON(http.StatusOK, gin.H{"data": alunos})
}

func GetAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	// Buscando turmas associadas
	var turmas []models.Turma
	if err := config.DB.Where("id IN ?", aluno.Turmas).Find(&turmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
		return
	}
	aluno.Turmas = nil // Limpando IDs após adicionar entidades

	c.JSON(http.StatusOK, aluno)
}

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
	aluno.Turmas = input.Turmas

	if err := config.DB.Save(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Aluno{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
