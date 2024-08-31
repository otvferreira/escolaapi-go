package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNota(c *gin.Context) {
	var nota models.Nota
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nota)
}

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
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota not found"})
		return
	}
	c.JSON(http.StatusOK, nota)
}

func UpdateNota(c *gin.Context) {
	id := c.Param("id")
	var nota models.Nota
	if err := config.DB.First(&nota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota not found"})
		return
	}
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&nota).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nota)
}

func DeleteNota(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Nota{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
