package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função para validar o total de pontos de uma turma
func validarTotalPontos(turmaID uint) (float64, error) {
	var total float64
	if err := config.DB.Model(&models.Atividade{}).Where("turma_id = ?", turmaID).Select("SUM(valor)").Row().Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func CreateAtividade(c *gin.Context) {
	var atividade models.Atividade
	if err := c.ShouldBindJSON(&atividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se a turma existe
	var turma models.Turma
	if err := config.DB.First(&turma, atividade.TurmaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma não encontrada"})
		return
	}

	// Validar total de pontos
	total, err := validarTotalPontos(atividade.TurmaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if total+atividade.Valor > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de pontos ultrapassa 100 pontos"})
		return
	}

	if err := config.DB.Create(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, atividade)
}

func GetAtividades(c *gin.Context) {
	var atividades []models.Atividade
	if err := config.DB.Find(&atividades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": atividades})
}

func GetAtividade(c *gin.Context) {
	id := c.Param("id")
	var atividade models.Atividade
	if err := config.DB.First(&atividade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade not found"})
		return
	}
	c.JSON(http.StatusOK, atividade)
}

func UpdateAtividade(c *gin.Context) {
	id := c.Param("id")
	var atividade models.Atividade
	if err := config.DB.First(&atividade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&atividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se a turma existe
	var turma models.Turma
	if err := config.DB.First(&turma, atividade.TurmaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma não encontrada"})
		return
	}

	// Validar total de pontos
	total, err := validarTotalPontos(atividade.TurmaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Subtrair o valor antigo da atividade para ajustar a validação
	var oldAtividade models.Atividade
	if err := config.DB.First(&oldAtividade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}
	total -= oldAtividade.Valor

	if total+atividade.Valor > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de pontos ultrapassa 100 pontos"})
		return
	}

	if err := config.DB.Save(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, atividade)
}

func DeleteAtividade(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Atividade{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
