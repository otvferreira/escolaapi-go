package routes

import (
	"backend/config"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validarTotalPontos(turmaID uint) (float64, error) {
	var total float64

	// Adicionando log para depuração antes da consulta
	log.Printf("Calculando o total de pontos para turma_id: %d", turmaID)

	// Usando o método Pluck para extrair o total de pontos
	err := config.DB.Model(&models.Atividade{}).Where("turma_id = ?", turmaID).Select("SUM(valor)").Row().Scan(&total)
	if err != nil {
		log.Printf("Erro ao calcular o total de pontos: %v", err) // Log do erro
		return 0, err
	}

	log.Printf("Total de pontos para turma_id %d: %f", turmaID, total) // Log do resultado

	return total, nil
}

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

	// Validar total de pontos
	total, err := validarTotalPontos(atividade.TurmaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular o total de pontos."})
		return
	}

	if total+atividade.Valor > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de pontos ultrapassa 100 pontos."})
		return
	}

	if err := config.DB.Create(&atividade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar atividade."})
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

	// Validar total de pontos
	total, err := validarTotalPontos(newAtividade.TurmaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular o total de pontos."})
		return
	}

	// Subtrair o valor antigo da atividade antes de adicionar o novo
	total = total - atividade.Valor + newAtividade.Valor

	if total > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total de pontos ultrapassa 100 pontos."})
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

func DeleteAtividade(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Atividade{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atividade não encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
