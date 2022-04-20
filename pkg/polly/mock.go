package polly

import (
	"context"

	"github.com/google/uuid"
	"github.com/navikt/nada-backend/pkg/graph/models"
)

type PollyMock struct {
	url string
}

func NewMock(url string) *PollyMock {
	return &PollyMock{
		url: url,
	}
}

func (m *PollyMock) SearchPolly(ctx context.Context, q string) ([]*models.PollyResult, error) {
	ret := []*models.PollyResult{}

	ret = append(ret, &models.PollyResult{
		ID:   uuid.MustParse("28570031-e2b3-4110-8864-41ab279e2e0c"),
		Name: "Behandling",
	})
	return ret, nil
}
