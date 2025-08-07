package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fesbarbosa/CSVreader/handlers"
	"github.com/gin-gonic/gin"
)

const (
	port      = 8080
	uploadDir = "./uploads"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Iniciando aplicação de leitura de CSV...")
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("Falha ao criar diretório de upload: %v", err)
	}

	absUploadDir, err := filepath.Abs(uploadDir)
	if err != nil {
		log.Fatalf("Falha ao obter caminho absoluto para uploads: %v", err)
	}
	log.Printf("Arquivos serão enviados para: %s", absUploadDir)

	router := gin.Default()
	csvHandler := handlers.NewCSVHandler(absUploadDir)
	csvHandler.RegisterRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API de processamento de CSV disponível. Use o endpoint /leitura/teste para enviar arquivos CSV.",
		})
	})

	serverAddr := ":8080"
	log.Printf("Servidor iniciado em http://localhost%s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}
}
