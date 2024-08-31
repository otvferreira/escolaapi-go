// backend/routes/professor.go
package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfessores(c *gin.Context) {
	var professores []models.Professor
	config.DB.Find(&professores)
	c.JSON(http.StatusOK, gin.H{"data": professores})
}

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
	c.JSON(http.StatusOK, professor)
}
