// backend/main.go
package main

import (
	"backend/config" // Importando o pacote config
	"backend/models" // Importando o pacote models
	"backend/routes" // Importando o pacote routes
	"log"
	"os"

	"github.com/gin-contrib/cors" // Importando o middleware CORS
	"github.com/gin-gonic/gin"    // Importando o framework Gin
	"gopkg.in/yaml.v2"
)

func main() {

	// Carregar o arquivo de configuração
	configFile := "config/config.yaml"
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de configuração: %v", err)
	}
	defer file.Close()

	// Analisar o arquivo de configuração
	var configData map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&configData); err != nil {
		log.Fatalf("Erro ao analisar o arquivo de configuração: %v", err)
	}

	// Inicializando o Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Conectando ao banco de dados
	config.ConnectDatabase()

	// Auto-migração dos modelos para o banco de dados
	config.DB.AutoMigrate(&models.Professor{}, &models.Turma{}, &models.Aluno{}, &models.Atividade{}, &models.Nota{})

	r.GET("/professores", routes.GetProfessores)
	r.GET("/professores/buscar/:id", routes.GetProfessor)
	r.POST("/professores/adicionar", routes.CreateProfessor)
	r.PUT("/professores/alterar/:id", routes.UpdateProfessor)
	r.DELETE("/professores/deletar/:id", routes.DeleteProfessor)

	r.GET("/alunos", routes.GetAlunos)
	r.GET("/alunos/buscar/:id", routes.GetAluno)
	r.POST("/alunos", routes.CreateAluno)
	r.PUT("/alunos/alterar/:id", routes.UpdateAluno)
	r.DELETE("/alunos/deletar/:id", routes.DeleteAluno)

	r.GET("/atividades", routes.GetAtividades)
	r.GET("/atividades/buscar/:id", routes.GetAtividade)
	r.POST("/atividades", routes.CreateAtividade)
	r.PUT("/atividades/alterar/:id", routes.UpdateAtividade)
	r.DELETE("/atividades/deletar/:id", routes.DeleteAtividade)

	r.GET("/notas", routes.GetNotas)
	r.GET("/notas/buscar/:id", routes.GetNota)
	r.POST("/notas", routes.CreateNota)
	r.PUT("/notas/alterar/:id", routes.UpdateNota)
	r.DELETE("/notas/deletar/:id", routes.DeleteNota)

	r.GET("/turmas", routes.GetTurmas)
	r.GET("/turmas/buscar/:id", routes.GetTurma)
	r.POST("/turmas", routes.CreateTurma)
	r.PUT("/turmas/alterar/:id", routes.UpdateTurma)
	r.DELETE("/turmas/deletar/:id", routes.DeleteTurma)

	// Iniciando o servidor
	r.Run()
}
