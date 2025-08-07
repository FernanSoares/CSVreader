package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fesbarbosa/CSVreader/utils"
	"github.com/gin-gonic/gin"
)

type CSVHandler struct {
	uploadDir string
}

func NewCSVHandler(uploadDir string) *CSVHandler {
	return &CSVHandler{
		uploadDir: uploadDir,
	}
}

func (h *CSVHandler) UploadAndProcessCSV(c *gin.Context) {
	log.Println("Recebida requisição de upload de arquivo")

	file, header, err := c.Request.FormFile("csvFile")
	if err != nil {
		log.Printf("Erro ao recuperar arquivo do formulário: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "erro",
			"mensagem": "Erro ao recuperar arquivo do formulário",
		})
		return
	}
	defer file.Close()

	log.Printf("Arquivo enviado: %s, tamanho: %d bytes, MIME: %s",
		header.Filename, header.Size, header.Header.Get("Content-Type"))

	if header.Size == 0 {
		log.Println("Erro: Arquivo enviado está vazio")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "erro",
			"mensagem": "Arquivo enviado está vazio",
		})
		return
	}

	filePath, err := utils.SaveUploadedFile(file, header.Filename, h.uploadDir)
	if err != nil {
		log.Printf("Erro ao salvar arquivo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "erro",
			"mensagem": fmt.Sprintf("Erro ao salvar arquivo: %v", err),
		})
		return
	}

	savedFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Erro ao abrir arquivo salvo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "erro",
			"mensagem": fmt.Sprintf("Erro ao abrir arquivo salvo: %v", err),
		})
		return
	}
	defer savedFile.Close()

	separator := ','
	products, err := utils.ProcessCSV(savedFile, separator)
	if err != nil {
		log.Printf("Erro ao processar arquivo CSV: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "erro",
			"mensagem": fmt.Sprintf("Erro ao processar arquivo CSV: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":             "sucesso",
		"mensagem":           "Arquivo processado com sucesso",
		"linhas_processadas": len(products),
	})
}

func (h *CSVHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/leitura/teste", h.UploadAndProcessCSV)
}
