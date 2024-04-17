# Golang Challenge - CRUD Application

## Descrição

Este projeto, desenvolvido como parte de um teste técnico para uma vaga de desenvolvedor Go, implementa um CRUD (Create, Read, Update, Delete) em Golang, integrado com um banco de dados PostgreSQL. 
A aplicação é desenhada para demonstrar princípios de uma arquitetura escalável e eficiente.

## Arquitetura Hexagonal

O projeto adota a **Arquitetura Hexagonal** (também conhecida como **Ports and Adapters**), que visa separar claramente a lógica de negócios do aplicativo das interações externas, como banco de dados e interfaces de usuário. Isso é feito através da definição de 'ports' e 'adapters':

- **Ports**: São interfaces dentro do domínio da aplicação. Eles representam as funções que a aplicação oferece, as quais podem ser chamadas por adaptadores externos.
- **Adapters**: Adaptam a informação externa da maneira que os ports esperam. Isso pode ser dividido em dois tipos:
  - **Adapters primários (ou de condução)**: Direcionam os dados de entrada para os ports da aplicação (por exemplo, a API REST que recebe chamadas HTTP).
  - **Adapters secundários (ou conduzidos)**: Implementam os ports definidos pela aplicação para interagir com sistemas externos, como bancos de dados.
 
Esta abordagem permite uma maior flexibilidade e testabilidade do código, tornando a aplicação facilmente adaptável a mudanças em tecnologias ou na lógica de negócios sem grandes reformulações.
![hexagonal](https://github.com/magomes-dev/go-challange/assets/11399094/53476f78-5813-4226-8701-7a9912c86af8)



### Estrutura de Diretórios
```
/go-challenge
├── cmd
│   └── main.go          # Ponto de entrada
├── internal
│   └── adapter
│       └── handler
│           |── http     # Adapter para a API HTTP
|           └── storage  # Adapter para o banco de dados
│   └── core
|       |── domain       # Modelos lógica de negócios
│       ├── ports        # Interfaces para os adapters
│       └── services     # Lógica de negócios             
├── Dockerfile
└── docker-compose.yml
```

## Tecnologias Utilizadas

- **Golang**
- **PostgreSQL**
- **Docker / Docker Compose**: Facilita o deploy e a configuração através de contêineres.
- **GORM**: ORM library for Golang
- **Air**: Live reload for Go apps

## Configuração do Projeto - Rodando localmente

### Pré-requisitos
 
- [Go](https://go.dev/doc/install)
- [Docker e Docker Compose](https://docs.docker.com/compose/install/)

### Instruções de Instalação e Execução 

Clone o Projeto
```bash
$ git clone https://github.com/magomes-dev/go-challenge.git
```
Entre no diretório do projeto  
```bash
$ cd go-challenge
```
Execute a API e o Banco de dados em container
```bash
$ docker-compose up --build
```
A API estará disponivel em: http://localhost:3000/ e o banco de dados na porta 5432:5432

## Uso da API
* **POST** /team: Adiciona um novo time.
    ##### Request
    ```bash
    curl --location 'localhost:3000/team' \
    --header 'Content-Type: application/json' \
    --data '
        {
            "name": "Grêmio FBPA",
            "country": "Brasil",
            "city": "Porto Alegre"
        }'
    ```
    ##### Response
    ```bash
    {
      "id": 1,
      "name": "Grêmio FBPA",
      "country": "Brasil",
      "city": "Porto Alegre"
    }
    ```

* **GET** /teams: Lista todos os times.
    ##### Request
    ```bash
    curl --location 'localhost:3000/teams'
    ```
    ##### Response
    ```bash
    
        [{
            "id": "1",
            "name": "Grêmio FBPA",
            "country": "Brasil",
            "city": "Porto Alegre"
        },
        {
            "id": "2",
            "name": "Sport Club Internacional",
            "country": "Brasil",
            "city": "Porto Alegre"
        }]
    ```
* **GET** /team/{id}: Recupera detalhes de um time específico.
  ##### Request
    ```bash
    curl --location 'localhost:3000/team/1'
    ```
  ##### Response
    ```bash
    {
      "id": 1,
      "name": "Grêmio FBPA",
      "country": "Brasil",
      "city": "Porto Alegre"
    }
    ```
* **PATCH** /team/{id}: Atualiza um time existente.
  ##### Request
  ```bash
    curl --location --request PATCH 'localhost:3000/team/1' \
  --header 'Content-Type: application/json' \
  --data '{
  	  "name": "Grêmio",
      "country": "Brasil",
      "city": "Porto Alegre"
  }'
  ```
  ##### Response
  ```bash
  {
    "id": 1,
    "name": "Grêmio",
    "country": "Brasil",
    "city": "Porto Alegre"
  }
  ```
* **DELETE** /team/{id}: Remove um time.
  ##### Request
  ```bash
  curl --location --request DELETE 'localhost:3000/team/2' \ --data ''
  ```
