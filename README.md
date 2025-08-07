# Serviço de Leitura de CSV

Este projeto implementa um serviço para ler arquivos CSV linha por linha em Go, com foco no processamento de arquivos que serão usados na atualização do Elasticsearch. Ele utiliza o framework Gin para uma API REST eficiente e segue uma arquitetura simples e direta para manter uma separação clara de responsabilidades.

## Funcionalidades

- API REST com framework Gin para upload e processamento de arquivos CSV
- Análise de arquivos CSV linha por linha
- Detecção e ignorar a linha de cabeçalho
- Mapeamento de dados CSV para objetos estruturados em Go
- Tratamento abrangente de erros com respostas JSON estruturadas
- Logs detalhados e middleware de recuperação

## Arquitetura do Projeto

O projeto segue uma arquitetura simples e intuitiva utilizando o framework Gin:

```
CSVreader/
├── models/      # Definições de estruturas de dados
├── handlers/    # Manipuladores HTTP usando Gin
├── utils/       # Funções utilitárias e processamento de CSV
├── uploads/     # Diretório para arquivos enviados
├── main.go      # Ponto de entrada da aplicação com configuração do Gin
└── README.md    # Documentação
```

## Como Funciona a Leitura de CSV

1. O serviço recebe um arquivo CSV via requisição HTTP POST para `/leitura/teste`
2. O arquivo é salvo localmente usando `os.Create` e `io.Copy`
3. O serviço abre o arquivo com `os.Open`
4. Ele lê linha por linha usando `bufio.NewScanner`
5. A primeira linha (cabeçalho) é ignorada
6. Para cada linha subsequente:
   - A linha é dividida em campos usando `strings.Split`
   - Os dados são mapeados para uma struct `Product`
   - A struct é registrada em log para fins de depuração

## Estrutura da Struct Product

```go
type Product struct {
	ID           string
	Name         string
	Description  string
	Brand        string
	Category     string
	Price        float64
	Currency     string
	Stock        int
	EAN          string
	Color        string
	Size         string
	Availability string
	InternalID   string
}
```

## Exemplo de Formato CSV

O serviço espera um arquivo CSV com as seguintes colunas:
```
Index,Name,Description,Brand,Category,Price,Currency,Stock,EAN,Color,Size,Availability,Internal ID
1,Thermostat,Descrição simples,Marca XYZ,Eletrônicos,74,USD,139,8619793560985,Preto,Médio,backorder,38
2,Ultra Speaker,Outra descrição,Marca ABC,Áudio,510,USD,351,3057216124300,Branco,Pequeno,backorder,27
```

## Como Executar o Serviço

1. Instale as dependências (caso necessário):
   ```
   go mod tidy
   ```
2. Construa e inicie a API:
   ```
   go run main.go
   ```
3. A API Gin será disponibilizada em `http://localhost:8080`
4. O endpoint GET `/` retorna um objeto JSON com informações sobre a API
5. Use o endpoint POST `/leitura/teste` para enviar arquivos CSV para processamento

## Testando o Endpoint

### Usando curl

```bash
curl -X POST -F "csvFile=@/caminho/para/seu/arquivo.csv" http://localhost:8080/leitura/teste
```

### Usando Postman

1. Crie uma nova requisição POST para `http://localhost:8080/leitura/teste`
2. Na aba "Body", selecione "form-data"
3. Adicione uma chave chamada "csvFile" e defina seu tipo como "File"
4. Selecione seu arquivo CSV
5. Envie a requisição

## Tratamento de Erros

O serviço trata vários cenários de erro:
- Arquivos vazios
- Linhas CSV mal formatadas
- Erros de acesso ao arquivo
- Tipos de dados inválidos nos campos CSV

Cada erro é devidamente registrado em log e uma resposta de erro descritiva é enviada ao cliente.
