package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/navikt/nada-backend/pkg/graph/generated"
	"github.com/navikt/nada-backend/pkg/graph/models"
	stripmd "github.com/writeas/go-strip-markdown/v2"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, q *models.SearchQueryOld, options *models.SearchQuery) ([]*models.SearchResultRow, error) {
	if q == nil {
		q = &models.SearchQueryOld{}
	}
	if options == nil {
		options = &models.SearchQuery{
			Text:   q.Text,
			Limit:  q.Limit,
			Offset: q.Offset,
			Types: []models.SearchType{
				models.SearchTypeDataproduct,
			},
		}

		if q.Keyword != nil {
			options.Keywords = []string{*q.Keyword}
		}
		if q.Group != nil {
			options.Groups = []string{*q.Group}
		}
	}
	return r.repo.Search(ctx, options)
}

// Excerpt is the resolver for the excerpt field.
func (r *searchResultRowResolver) Excerpt(ctx context.Context, obj *models.SearchResultRow) (string, error) {
	return stripmd.Strip(obj.Excerpt), nil
}

// SearchResultRow returns generated.SearchResultRowResolver implementation.
func (r *Resolver) SearchResultRow() generated.SearchResultRowResolver {
	return &searchResultRowResolver{r}
}

type searchResultRowResolver struct{ *Resolver }
