package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/database"
	"github.com/navikt/nada-backend/pkg/graph/generated"
	"github.com/navikt/nada-backend/pkg/graph/models"
)

var ErrUnauthorized = fmt.Errorf("unauthorized")

type GCP interface {
	GetTables(ctx context.Context, projectID, datasetID string) ([]*models.BigQueryTable, error)
	GetDatasets(ctx context.Context, projectID string) ([]string, error)
}

type Resolver struct {
	repo        *database.Repo
	gcp         GCP
	gcpProjects *auth.TeamProjectsUpdater
}

func New(repo *database.Repo, gcp GCP, gcpProjects *auth.TeamProjectsUpdater) *handler.Server {
	resolver := &Resolver{
		repo:        repo,
		gcp:         gcp,
		gcpProjects: gcpProjects,
	}

	config := generated.Config{Resolvers: resolver}
	config.Directives.Authenticated = authenticate
	return handler.NewDefaultServer(generated.NewExecutableSchema(config))
}

func pagination(limit *int, offset *int) (int, int) {
	l := 15
	o := 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}
	return l, o
}

func ensureUserInGroup(ctx context.Context, group string) error {
	user := auth.GetUser(ctx)
	if user == nil || !user.Groups.Contains(group) {
		return ErrUnauthorized
	}
	return nil
}

func authenticate(ctx context.Context, obj interface{}, next graphql.Resolver, on *bool) (res interface{}, err error) {
	if auth.GetUser(ctx) == nil {
		// block calling the next resolver
		return nil, fmt.Errorf("access denied")
	}

	// or let it pass through
	return next(ctx)
}
