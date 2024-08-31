package routes

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	c.JSON(http.StatusOK, gin.H{"data": alunos})
}

func GetAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno not found"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func UpdateAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno
	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno not found"})
		return
	}
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Aluno{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
