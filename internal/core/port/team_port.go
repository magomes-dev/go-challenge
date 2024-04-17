package port

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magomes-dev/go-challange/internal/core/domain"
)

type TeamRepository interface {
	CreateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error)
	ListTeams(ctx *fiber.Ctx) ([]domain.Team, error)
	GetTeamById(ctx *fiber.Ctx, id string) (*domain.Team, error)
	UpdateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error)
	DeleteTeam(ctx *fiber.Ctx, id string) error
}

type TeamService interface {
	CreateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error)
	ListTeams(ctx *fiber.Ctx) ([]domain.Team, error)
	GetTeamById(ctx *fiber.Ctx, id string) (*domain.Team, error)
	UpdateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error)
	DeleteTeam(ctx *fiber.Ctx, id string) error
}
