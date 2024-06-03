# Desafio: Client-Server-API

## Descrição

Neste desafio, você criará dois sistemas em Go: `client.go` e `server.go`. O objetivo é aplicar conceitos de webserver
HTTP, contextos, banco de dados e manipulação de arquivos.

### Requisitos

- `client.go`:
    - Realiza uma requisição HTTP para `server.go` solicitando a cotação do dólar.
    - Recebe o valor atual do câmbio (campo "bid" do JSON) do `server.go`.
    - Salva a cotação em um arquivo "cotacao.txt" no formato: `Dólar: {valor}`.
    - Usa o package `context` com um timeout máximo de 300ms para receber o resultado do `server.go`.

- `server.go`:
    - Consome a API [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) para obter a cotação de Dólar
      para Real.
    - Retorna o resultado no formato JSON para o cliente.
    - Usa o package `context`:
        - Timeout máximo de 200ms para chamar a API de cotação do dólar.
        - Timeout máximo de 10ms para persistir os dados no banco de dados SQLite.
    - Registra no banco de dados SQLite cada cotação recebida.
    - Disponibiliza um endpoint `/cotacao` na porta 8080.

### Pré-requisitos

- Go 1.18+ instalado

### Clonar o Repositório

```bash
git clone https://github.com/souluanf/client-server-challenge-go.git
cd client-server-challenge-go
```

### Executar a Aplicação

Instale as dependências:

```bash
go mod download
```

#### Executar o Servidor

No diretório raiz do projeto, execute:

```bash
go run cmd/server/main.go
```

#### Executar o Cliente

Em outro terminal, no diretório raiz do projeto, execute:

```bash
go run cmd/client/main.go
```

### Testar o Endpoint

Você pode testar o endpoint do servidor usando `curl`:

```bash
curl --location 'localhost:8080/cotacao'
```

### Exemplo de Saída

```json
{
  "bid": "5.2462"
}
```

## Dados

Os valores das cotações serão adicionados na pasta `data`:

- `data/quotation.db`
- `data/quotation.txt`