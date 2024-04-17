# Golang Challenge - CRUD Application

## Description

This project, developed as part of a technical test for a Go developer position, implements a CRUD (Create, Read, Update, Delete) in Golang, integrated with a PostgreSQL database.
The application is designed to demonstrate principles of a scalable and efficient architecture.

## Hexagonal Architecture
The project adopts the Hexagonal Architecture (also known as Ports and Adapters), which aims to clearly separate the application's business logic from external interactions, such as the database and user interfaces. This is done by defining 'ports' and 'adapters':

Ports: Are interfaces within the application's domain. They represent the functions that the application provides, which can be called by external adapters.
Adapters: Adapt external information in the way that ports expect. This can be divided into two types:
Primary (or driving) Adapters: Direct input data to the application's ports (e.g., the REST API that receives HTTP calls).
Secondary (or driven) Adapters: Implement the ports defined by the application to interact with external systems, such as databases.
This approach allows for greater flexibility and testability of the code, making the application easily adaptable to changes in technologies or business logic without major overhauls.
![hexagonal](https://github.com/magomes-dev/go-challange/assets/11399094/53476f78-5813-4226-8701-7a9912c86af8)



### Estrutura de Diretórios
```
/go-challenge
├── cmd
│   └── main.go          # Entry point
├── internal
│   └── adapter
│       └── handler
│           |── http     # Adapter for HTTP API
|           └── storage  # Adapter for the database
│   └── core
|       |── domain       # Business logic models
│       ├── ports        # Interfaces for the adapters
│       └── services     # Business logic             
├── Dockerfile
└── docker-compose.yml
```

## Technologies Used

- **Golang**
- **PostgreSQL**
- **Docker / Docker Compose**: Facilitates deployment and configuration through containers.
- **GORM**: ORM library for Golang
- **Air**: Live reload for Go apps

## Project Setup - Running Locally

### Prerequisites
 
- [Go](https://go.dev/doc/install)
- [Docker e Docker Compose](https://docs.docker.com/compose/install/)

### Installation and Execution Instructions

Clone the Project
```bash
$ git clone https://github.com/magomes-dev/go-challenge.git
```
Navigate to the project directory
```bash
$ cd go-challenge
```
Run the API and the Database in containers
```bash
$ docker-compose up --build
```
The API will be available at: http://localhost:3000/ and the database at port 5432:5432

## API Usage
* **POST** /team: Adds a new team.
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

* **GET** /teams: Lists all teams.
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
* **GET** /team/{id}: Retrieves details of a specific team.
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
* **PATCH** /team/{id}: Updates an existing team.
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
* **DELETE** /team/{id}: Removes a team.
  ##### Request
  ```bash
  curl --location --request DELETE 'localhost:3000/team/2' \ --data ''
  ```
