package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/magomes-dev/go-challange/internal/core/domain"
	"github.com/magomes-dev/go-challange/internal/core/port"
)

// CategoryHandler represents the HTTP handler for category-related requests
type TeamHandler struct {
	teamservice port.TeamService
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewTeamHandler(teamservice port.TeamService) *TeamHandler {
	return &TeamHandler{
		teamservice,
	}
}

func (th *TeamHandler) Home(c *fiber.Ctx) error {
	return c.SendString("golang challange @magomes-dev!")
}

func (th *TeamHandler) CreateTeam(c *fiber.Ctx) error {
	req := new(TeamRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	team := domain.Team{
		Name:    req.Name,
		Country: req.Country,
		City:    req.City,
	}

	_, err := th.teamservice.CreateTeam(c, &team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	return c.Status(fiber.StatusCreated).JSON(toTeamResponse(team))
}

func (th *TeamHandler) ListTeams(ctx *fiber.Ctx) error {
	teams, err := th.teamservice.ListTeams(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(toTeamsResponseSlice(teams))
}

func (th *TeamHandler) GetTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	team, err := th.teamservice.GetTeamById(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	if team.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ID not found"})
	}

	return c.Status(fiber.StatusOK).JSON(toTeamResponse(*team))
}

func (th *TeamHandler) UpdateTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(TeamRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot parse id"})
	}

	team := domain.Team{
		ID:      num,
		Name:    req.Name,
		Country: req.Country,
		City:    req.City,
	}

	updatedTeam, err := th.teamservice.UpdateTeam(c, &team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	response := toTeamResponse(*updatedTeam)
	return c.JSON(response)
}

func (th *TeamHandler) DeleteTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	err := th.teamservice.DeleteTeam(c, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.SendStatus(204)
}
