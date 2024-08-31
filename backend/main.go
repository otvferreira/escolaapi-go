// backend/main.go
package main

import (
	"backend/config" // Importando o pacote config
	"backend/models" // Importando o pacote models
	"backend/routes" // Importando o pacote routes

	"github.com/gin-gonic/gin" // Importando o framework Gin
)

func main() {
	// Inicializando o Gin
	r := gin.Default()

	// Conectando ao banco de dados
	config.ConnectDatabase()

	// Auto-migração dos modelos para o banco de dados
	config.DB.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{}, &models.Nota{})

	// Definindo as rotas
	r.GET("/professores", routes.GetProfessores)
	r.POST("/professores", routes.CreateProfessor)

	// Iniciando o servidor
	r.Run()
}
