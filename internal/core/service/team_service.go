package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magomes-dev/go-challange/internal/core/domain"
	"github.com/magomes-dev/go-challange/internal/core/port"
)

type TeamService struct {
	repository port.TeamRepository
}

func NewTeamService(repository port.TeamRepository) *TeamService {
	return &TeamService{
		repository,
	}
}

func (ts *TeamService) CreateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error) {
	team, err := ts.repository.CreateTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (ts *TeamService) GetTeamById(ctx *fiber.Ctx, id string) (*domain.Team, error) {
	var team *domain.Team
	team, err := ts.repository.GetTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (ts *TeamService) ListTeams(ctx *fiber.Ctx) ([]domain.Team, error) {
	var teams []domain.Team
	teams, err := ts.repository.ListTeams(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (ts *TeamService) UpdateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error) {
	_, err := ts.repository.UpdateTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (ts *TeamService) DeleteTeam(ctx *fiber.Ctx, id string) error {
	return ts.repository.DeleteTeam(ctx, id)
}
