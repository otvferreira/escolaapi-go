package routes

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateAluno(c *gin.Context) {
	var input struct {
		Nome      string `json:"nome"`
		Matricula string `json:"matricula"`
		Turmas    string `json:"turmas"` // IDs das turmas como string separada por vírgulas
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aluno := models.Aluno{
		Nome:      input.Nome,
		Matricula: input.Matricula,
		TurmaIDs:  input.Turmas, // Salva a string diretamente
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

	for i := range alunos {
		var turmas []models.Turma
		if err := config.DB.Where("id IN ?", strings.Split(alunos[i].TurmaIDs, ",")).Find(&turmas).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
			return
		}
		// Adiciona a lista de turmas à estrutura de resposta
		alunos[i].Turmas = turmas
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

	var turmas []models.Turma
	if err := config.DB.Where("id IN ?", strings.Split(aluno.TurmaIDs, ",")).Find(&turmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar turmas"})
		return
	}
	// Adiciona a lista de turmas à estrutura de resposta
	aluno.Turmas = turmas

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
		Turmas    string `json:"turmas"` // IDs das turmas como string separada por vírgulas
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualizando os dados do aluno
	aluno.Nome = input.Nome
	aluno.Matricula = input.Matricula
	aluno.TurmaIDs = input.Turmas // Atualiza TurmaIDs

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
