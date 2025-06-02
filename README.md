# mbalab1

Projeto em Go para consulta de clima por CEP, utilizando Echo para API HTTP.

## Descrição

Este projeto expõe uma API REST que, ao receber um CEP, retorna informações meteorológicas da localidade correspondente. Ele integra dois serviços externos:

- [ViaCEP](https://viacep.com.br/) para consulta de dados do CEP.
- [wttr.in](https://wttr.in/) para consulta de informações do clima.

## Estrutura do Projeto

- `cmd/`: Ponto de entrada da aplicação.
- `config/`: Carregamento de variáveis de ambiente.
- `internal/`: Lógica interna, incluindo casos de uso, handlers, infraestrutura HTTP e web.
- `api/`: Exemplos de requisições HTTP.
- `.env`: Variáveis de ambiente.
- `Dockerfile`: Containerização da aplicação.
# mbalab1

Projeto em Go para consulta de clima por CEP, utilizando Echo para API HTTP.

## Descrição

Este projeto expõe uma API REST que, ao receber um CEP, retorna informações meteorológicas da localidade correspondente. Ele integra dois serviços externos:

- [ViaCEP](https://viacep.com.br/) para consulta de dados do CEP.
- [wttr.in](https://wttr.in/) para consulta de informações do clima.

## Servidor Online para Testes

Você pode testar a API sem instalar nada, utilizando o servidor online:

```
https://lab-1-91072667969.us-central1.run.app/weather/{cep}
```

## Estrutura do Projeto

- `cmd/`: Ponto de entrada da aplicação.
- `config/`: Carregamento de variáveis de ambiente.
- `internal/`: Lógica interna, incluindo casos de uso, handlers, infraestrutura HTTP e web.
- `api/`: Exemplos de requisições HTTP.
- `.env`: Variáveis de ambiente.
- `Dockerfile`: Containerização da aplicação.

## Variáveis de Ambiente

Configure o arquivo `.env` na raiz do projeto:

```
CEP_SERVICE_PATH=https://viacep.com.br/ws/%s/json/
WEATHER_SERVICE_PATH=https://wttr.in/%s?format=j1
WEB_PORT=8080
```

## Como Executar

### Localmente

1. Instale o Go (versão 1.24+).
2. Instale as dependências:

   ```sh
   go mod tidy
   ```

3. Execute a aplicação:

   ```sh
   go run ./cmd/main.go
   ```

A API estará disponível em `http://localhost:8080`.

### Com Docker

1. Construa a imagem:

   ```sh
   docker build -t mbalab1 .
   ```

2. Rode o container:

   ```sh
   docker run -p 8080:8080 --env-file .env mbalab1
   ```

## Endpoints

- `GET /weather/:cep`  
  Retorna informações do clima para o CEP informado.

  **Exemplo:**
  ```
  GET /weather/18050-605
  ```

  **Resposta:**
  ```json
  {
    "data": {
      "temp_c": "25.00",
      "temp_f": "77.00",
      "temp_k": "298.15"
    }
  }
  ```

- `GET /health`  
  Checagem de saúde da API.

## Testes

Execute os testes com:

```sh
go test ./internal/...
```

## Licença

MIT
## Variáveis de Ambiente

Configure o arquivo `.env` na raiz do projeto:

```
CEP_SERVICE_PATH=https://viacep.com.br/ws/%s/json/
WEATHER_SERVICE_PATH=https://wttr.in/%s?format=j1
WEB_PORT=8080
```

## Como Executar

### Localmente

1. Instale o Go (versão 1.24+).
2. Instale as dependências:

   ```sh
   go mod download
   ```

3. Execute a aplicação:

   ```sh
   go run ./cmd/main.go
   ```

A API estará disponível em `http://localhost:8080`.

### Com Docker

1. Construa a imagem:

   ```sh
   docker build -t mbalab1 .
   ```

2. Rode o container:

   ```sh
   docker run -p 8080:8080 --env-file .env mbalab1
   ```

## Endpoints

- `GET /weather/:cep`  
  Retorna informações do clima para o CEP informado.

  **Exemplo:**
  ```
  GET /weather/18050-605
  ```

  **Resposta:**
  ```json
  {
    "data": {
      "temp_c": "25.00",
      "temp_f": "77.00",
      "temp_k": "298.15"
    }
  }
  ```

- `GET /health`  
  Checagem de saúde da API.

## Testes

Execute os testes com:

```sh
go test ./internal/...
```

## Licença

MIT
