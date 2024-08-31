package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProfessor(c *gin.Context) {
	var professor models.Professor
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&professor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, professor)
}

func GetProfessores(c *gin.Context) {
	var professores []models.Professor
	if err := config.DB.Find(&professores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": professores})
}

func GetProfessor(c *gin.Context) {
	id := c.Param("id")
	var professor models.Professor
	if err := config.DB.First(&professor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
		return
	}
	c.JSON(http.StatusOK, professor)
}

func UpdateProfessor(c *gin.Context) {
	id := c.Param("id")
	var professor models.Professor
	if err := config.DB.First(&professor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
		return
	}
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&professor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, professor)
}

func DeleteProfessor(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Professor{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
