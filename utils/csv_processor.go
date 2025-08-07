package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fesbarbosa/CSVreader/models"
)

func ProcessCSV(file io.Reader, separator rune) ([]models.Product, error) {
	log.Println("Iniciando processamento do CSV")

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		log.Println("Erro: Arquivo vazio")
		return nil, fmt.Errorf("arquivo está vazio")
	}

	log.Println("Pulando linha de cabeçalho")

	var products []models.Product
	lineNumber := 1

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		log.Printf("Processando linha %d: %s", lineNumber, line)

		fields := strings.Split(line, string(separator))

		if len(fields) < 13 {
			log.Printf("Erro na linha %d: campos insuficientes (esperado 13, obteve %d)", lineNumber, len(fields))
			continue
		}

		price, err := strconv.ParseFloat(fields[5], 64)
		if err != nil {
			log.Printf("Erro ao converter preço na linha %d: %v", lineNumber, err)
			price = 0.0
		}

		stock, err := strconv.Atoi(fields[7])
		if err != nil {
			log.Printf("Erro ao converter estoque na linha %d: %v", lineNumber, err)
			stock = 0
		}

		product := models.Product{
			ID:           fields[0],
			Name:         fields[1],
			Description:  fields[2],
			Brand:        fields[3],
			Category:     fields[4],
			Price:        price,
			Currency:     fields[6],
			Stock:        stock,
			EAN:          fields[8],
			Color:        fields[9],
			Size:         fields[10],
			Availability: fields[11],
			InternalID:   fields[12],
		}

		products = append(products, product)
		log.Printf("Produto processado com sucesso: %s", product.Name)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Erro ao ler CSV: %v", err)
		return products, fmt.Errorf("erro ao ler CSV: %v", err)
	}

	log.Printf("Processamento CSV concluído. Processados %d produtos", len(products))
	return products, nil
}

func SaveUploadedFile(fileData io.Reader, fileName, uploadDir string) (string, error) {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, 0755)
		if err != nil {
			log.Printf("Erro ao criar diretório de upload: %v", err)
			return "", fmt.Errorf("erro ao criar diretório de upload: %v", err)
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	uniqueFileName := fmt.Sprintf("%s_%s", timestamp, fileName)
	filePath := filepath.Join(uploadDir, uniqueFileName)

	log.Printf("Salvando arquivo em: %s", filePath)

	outputFile, err := os.Create(filePath)
	if err != nil {
		log.Printf("Erro ao criar arquivo: %v", err)
		return "", fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer outputFile.Close()

	bytesWritten, err := io.Copy(outputFile, fileData)
	if err != nil {
		log.Printf("Erro ao copiar dados do arquivo: %v", err)
		return "", fmt.Errorf("erro ao copiar dados do arquivo: %w", err)
	}

	log.Printf("Arquivo salvo com sucesso %s (%d bytes escritos)", filePath, bytesWritten)
	return filePath, nil
}
