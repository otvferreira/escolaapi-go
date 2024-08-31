package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTurma(c *gin.Context) {
	var turma models.Turma
	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&turma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, turma)
}

func GetTurmas(c *gin.Context) {
	var turmas []models.Turma
	if err := config.DB.Find(&turmas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": turmas})
}

func GetTurma(c *gin.Context) {
	id := c.Param("id")
	var turma models.Turma
	if err := config.DB.First(&turma, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma not found"})
		return
	}
	c.JSON(http.StatusOK, turma)
}

func UpdateTurma(c *gin.Context) {
	id := c.Param("id")
	var turma models.Turma
	if err := config.DB.First(&turma, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma not found"})
		return
	}
	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&turma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, turma)
}

func DeleteTurma(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Turma{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Turma not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
