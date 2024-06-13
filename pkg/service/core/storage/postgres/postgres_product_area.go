package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/navikt/nada-backend/pkg/database"
	"github.com/navikt/nada-backend/pkg/service"
)

var _ service.ProductAreaStorage = &productAreaStorage{}

type productAreaStorage struct {
	db *database.Repo
}

func (s *productAreaStorage) GetDashboard(ctx context.Context, id string) (*service.Dashboard, error) {
	dashboard, err := s.db.Querier.GetDashboard(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get dashboard: %w", err)
	}

	return &service.Dashboard{
		ID:  dashboard.ID,
		Url: dashboard.Url,
	}, nil
}

func (s *productAreaStorage) GetProductArea(ctx context.Context, paID string) (*service.ProductArea, error) {
	pa, err := s.db.Querier.GetProductArea(ctx, uuid.MustParse(paID))
	if err != nil {
		return nil, err
	}

	teams, err := s.db.Querier.GetTeamsInProductArea(ctx, uuid.NullUUID{
		UUID:  pa.ID,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}

	paTeams := make([]*service.Team, 0)
	for _, team := range teams {
		paTeams = append(paTeams, &service.Team{
			TeamkatalogenTeam: &service.TeamkatalogenTeam{
				ID:            team.ID.String(),
				Name:          team.Name.String,
				ProductAreaID: team.ProductAreaID.UUID.String(),
			},
		})
	}

	areaType := ""
	if pa.AreaType.Valid {
		areaType = pa.AreaType.String
	}

	return &service.ProductArea{
		TeamkatalogenProductArea: &service.TeamkatalogenProductArea{
			ID:       pa.ID.String(),
			Name:     pa.Name.String,
			AreaType: areaType,
		},
		Teams: paTeams,
	}, nil
}

func (s *productAreaStorage) GetProductAreas(ctx context.Context) ([]*service.ProductArea, error) {
	pas, err := s.db.Querier.GetProductAreas(ctx)
	if err != nil {
		return nil, err
	}

	// FIXME: not optimal, but unsure how else to do this
	teams, err := s.db.Querier.GetAllTeams(ctx)
	if err != nil {
		return nil, err
	}

	productAreas := make([]*service.ProductArea, 0)
	for _, pa := range pas {
		paTeams := make([]*service.Team, 0)
		for _, team := range teams {
			if team.ProductAreaID.Valid && team.ProductAreaID.UUID == pa.ID {
				paTeams = append(paTeams, &service.Team{
					TeamkatalogenTeam: &service.TeamkatalogenTeam{
						ID:            team.ID.String(),
						Name:          team.Name.String,
						ProductAreaID: team.ProductAreaID.UUID.String(),
					},
				})
			}
		}
		areaType := ""
		if pa.AreaType.Valid {
			areaType = pa.AreaType.String
		}
		productAreas = append(productAreas, &service.ProductArea{
			TeamkatalogenProductArea: &service.TeamkatalogenProductArea{
				ID:       pa.ID.String(),
				Name:     pa.Name.String,
				AreaType: areaType,
			},
			Teams:        paTeams,
			DashboardURL: "",
		})
	}

	return productAreas, nil
}

func NewProductAreaStorage(db *database.Repo) *productAreaStorage {
	return &productAreaStorage{db: db}
}
