package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função para criar uma nova atividade
func CreateAtividade(c *gin.Context) {
	var atividade models.Atividade

	if err := c.ShouldBindJSON(&atividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos. Verifique os campos obrigatórios."})
		return
	}

	if atividade.TurmaID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'turma_id' é obrigatório."})
		return
	}

	if atividade.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'valor' deve ser maior que zero."})
		return
	}

	// Verificar se a turma existe
	var turma models.Turma
	if err := config.DB.First(&turma, atividade.TurmaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma não encontrada."})
		return
	}

	// Criar a atividade
	if err := config.DB.Create(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar atividade."})
		return
	}

	c.JSON(http.StatusCreated, atividade)
}

// Função para obter todas as atividades
func GetAtividades(c *gin.Context) {
	var atividades []models.Atividade
	if err := config.DB.Find(&atividades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": atividades})
}

// Função para obter uma atividade específica
func GetAtividade(c *gin.Context) {
	id := c.Param("id")
	var atividade models.Atividade
	if err := config.DB.First(&atividade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}
	c.JSON(http.StatusOK, atividade)
}

// Função para atualizar uma atividade existente
func UpdateAtividade(c *gin.Context) {
	id := c.Param("id")
	var atividade models.Atividade

	if err := config.DB.First(&atividade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada."})
		return
	}

	var newAtividade models.Atividade
	if err := c.ShouldBindJSON(&newAtividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos. Verifique os campos obrigatórios."})
		return
	}

	// Verificar se a turma existe
	var turma models.Turma
	if err := config.DB.First(&turma, newAtividade.TurmaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma não encontrada."})
		return
	}

	atividade.TurmaID = newAtividade.TurmaID
	atividade.Valor = newAtividade.Valor
	atividade.Data = newAtividade.Data

	if err := config.DB.Save(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar atividade."})
		return
	}

	c.JSON(http.StatusOK, atividade)
}

// Função para deletar uma atividade existente
func DeleteAtividade(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Atividade{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
