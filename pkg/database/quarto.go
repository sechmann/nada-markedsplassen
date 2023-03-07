package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/navikt/nada-backend/pkg/database/gensql"
	"github.com/navikt/nada-backend/pkg/graph/models"
)

func (r *Repo) CreateQuartoStory(ctx context.Context, creator string,
	newQuartoStory models.NewQuartoStory,
) (models.QuartoStory, error) {
	quartoSQL, err := r.querier.CreateQuartoStory(ctx, gensql.CreateQuartoStoryParams{
		Name:             newQuartoStory.Name,
		Creator:          creator,
		Description:      newQuartoStory.Description,
		Keywords:         newQuartoStory.Keywords,
		TeamkatalogenUrl: ptrToNullString(newQuartoStory.TeamkatalogenURL),
		ProductAreaID:    ptrToNullString(newQuartoStory.ProductAreaID),
		TeamID:           ptrToNullString(newQuartoStory.TeamID),
		OwnerGroup:       newQuartoStory.Group,
	})

	return *quartoSQLToGraphql(quartoSQL), err
}

func (r *Repo) GetQuartoStory(ctx context.Context, id uuid.UUID) (*models.QuartoStory, error) {
	quartoSQL, err := r.querier.GetQuartoStory(ctx, id)

	return quartoSQLToGraphql(quartoSQL), err
}

func (r *Repo) GetQuartoStories(ctx context.Context) ([]*models.QuartoStory, error) {
	quartoSQLs, err := r.querier.GetQuartoStories(ctx)
	if err != nil {
		return nil, err
	}

	quartoGraphqls := make([]*models.QuartoStory, len(quartoSQLs))
	for idx, quarto := range quartoSQLs {
		quartoGraphqls[idx] = quartoSQLToGraphql(quarto)
	}

	return quartoGraphqls, err
}

func quartoSQLToGraphql(quarto gensql.QuartoStory) *models.QuartoStory {
	return &models.QuartoStory{
		ID:            quarto.ID,
		Name:          quarto.Name,
		Creator:       quarto.Creator,
		Created:       quarto.Created,
		LastModified:  &quarto.LastModified,
		Keywords:      quarto.Keywords,
		ProductAreaID: nullStringToPtr(quarto.ProductAreaID),
		TeamID:        nullStringToPtr(quarto.TeamID),
		Description:   quarto.Description,
		Group:         quarto.Group,
	}
}
