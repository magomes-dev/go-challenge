package http

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, teamHandler *TeamHandler) {
	app.Get("/", teamHandler.Home)
	app.Get("/teams", teamHandler.ListTeams)
	app.Get("/team/:id", teamHandler.GetTeam)
	app.Post("/team", teamHandler.CreateTeam)
	app.Patch("/team/:id", teamHandler.UpdateTeam)
	app.Delete("/team/:id", teamHandler.DeleteTeam)
}
