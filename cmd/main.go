package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/magomes-dev/go-challange/internal/adapter/handler/http"
	"github.com/magomes-dev/go-challange/internal/adapter/storage/postgres"
	"github.com/magomes-dev/go-challange/internal/adapter/storage/postgres/repository"
	"github.com/magomes-dev/go-challange/internal/core/service"
)

func main() {

	postgres.ConnectDb()

	app := fiber.New()

	teamRepo := repository.NewTeamRepository(&postgres.DB)
	teamService := service.NewTeamService(teamRepo)
	teamHandler := http.NewTeamHandler(teamService)

	http.RegisterRoutes(app, teamHandler)

	app.Listen(":3000")
}
