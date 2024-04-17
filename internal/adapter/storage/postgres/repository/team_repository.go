package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magomes-dev/go-challange/internal/adapter/storage/postgres"
	"github.com/magomes-dev/go-challange/internal/core/domain"
)

type TeamRepository struct {
	db *postgres.Dbinstance
}

func NewTeamRepository(db *postgres.Dbinstance) *TeamRepository {
	return &TeamRepository{
		db,
	}
}

func (tr *TeamRepository) CreateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error) {
	var err = tr.db.Db.Create(team)
	if err != nil {
		return nil, err.Error
	}
	return team, nil
}

func (tr *TeamRepository) ListTeams(ctx *fiber.Ctx) ([]domain.Team, error) {
	var teams []domain.Team
	result := tr.db.Db.Find(&teams)

	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}

func (tr *TeamRepository) GetTeamById(ctx *fiber.Ctx, id string) (*domain.Team, error) {
	var team domain.Team
	result := tr.db.Db.Find(&team, id)
	return &team, result.Error
}

func (tr *TeamRepository) UpdateTeam(ctx *fiber.Ctx, team *domain.Team) (*domain.Team, error) {
	result := tr.db.Db.Model(&team).Updates(team)
	return team, result.Error
}

func (tr *TeamRepository) DeleteTeam(ctx *fiber.Ctx, id string) error {
	result := tr.db.Db.Delete(&domain.Team{}, id)
	return result.Error
}
