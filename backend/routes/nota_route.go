package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Obter alunos por turma
func GetAlunosByTurma(c *gin.Context) {
	turmaID := c.Param("id")
	var alunos []models.Aluno
	if err := config.DB.Where("turma_id = ?", turmaID).Find(&alunos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": alunos})
}

func CreateNota(c *gin.Context) {
	var nota models.Nota
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se a atividade existe e buscar seu valor máximo permitido
	var atividade models.Atividade
	if err := config.DB.First(&atividade, nota.AtividadeID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade não encontrada"})
		return
	}

	// Verificar se a nota excede o valor máximo da atividade
	if nota.Nota > atividade.Valor {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nota não pode ser maior que o valor máximo da atividade"})
		return
	}

	// Calcular o total de notas existentes para a atividade
	var totalNotas float64
	config.DB.Model(&models.Nota{}).Where("atividade_id = ?", nota.AtividadeID).Select("sum(nota)").Scan(&totalNotas)

	// Verificar se o total de notas ultrapassa o máximo permitido (100 pontos)
	if totalNotas+nota.Nota > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de notas para esta atividade excede o limite de 100 pontos"})
		return
	}

	// Criar a nota
	if err := config.DB.Create(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nota)
}

// Função para atualizar uma nota existente
func UpdateNota(c *gin.Context) {
	id := c.Param("id")
	var nota models.Nota
	if err := config.DB.First(&nota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se a atividade existe e buscar seu valor máximo permitido
	var atividade models.Atividade
	if err := config.DB.First(&atividade, nota.AtividadeID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Atividade não encontrada"})
		return
	}

	// Verificar se a nota excede o valor máximo da atividade
	if nota.Nota > atividade.Valor {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nota não pode ser maior que o valor máximo da atividade"})
		return
	}

	// Calcular o total de notas existentes para a atividade, excluindo a nota atual
	var totalNotas float64
	config.DB.Model(&models.Nota{}).Where("atividade_id = ? AND id != ?", nota.AtividadeID, id).Select("sum(nota)").Scan(&totalNotas)

	// Verificar se o total de notas ultrapassa o máximo permitido (100 pontos)
	if totalNotas+nota.Nota > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de notas para esta atividade excede o limite de 100 pontos"})
		return
	}

	// Atualizar a nota
	if err := config.DB.Save(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nota)
}

// Outras funções não precisam de alterações
func GetNotas(c *gin.Context) {
	var notas []models.Nota
	if err := config.DB.Find(&notas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": notas})
}

func GetNota(c *gin.Context) {
	id := c.Param("id")
	var nota models.Nota
	if err := config.DB.First(&nota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
		return
	}
	c.JSON(http.StatusOK, nota)
}

func DeleteNota(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Nota{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota não encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetNotasByTurmaEAtividade(c *gin.Context) {
	turmaID := c.Param("turma_id")
	atividadeID := c.Param("atividade_id")

	var notas []models.Nota
	if err := config.DB.Where("turma_id = ? AND atividade_id = ?", turmaID, atividadeID).Find(&notas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": notas})
}
