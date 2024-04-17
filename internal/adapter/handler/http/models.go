package http

import "github.com/magomes-dev/go-challange/internal/core/domain"

// ProductRequest é usado para criar e atualizar produtos
type TeamRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
}

// ProductResponse representa como os dados do produto são enviados na resposta
type TeamResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
}

// Mapeia o domínio do produto para a resposta do produto
func toTeamResponse(t domain.Team) TeamResponse {
	return TeamResponse{
		ID:      t.ID,
		Name:    t.Name,
		Country: t.Country,
		City:    t.City,
	}
}

// Mapeia uma lista de produtos de domínio para uma lista de respostas de produtos
func toTeamsResponseSlice(teams []domain.Team) []TeamResponse {
	var resp []TeamResponse
	for _, t := range teams {
		resp = append(resp, toTeamResponse(t))
	}
	return resp
}
