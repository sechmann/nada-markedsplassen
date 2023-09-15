package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/graph/generated"
	"github.com/navikt/nada-backend/pkg/graph/models"
)

// UserInfo is the resolver for the userInfo field.
func (r *queryResolver) UserInfo(ctx context.Context) (*models.UserInfo, error) {
	user := auth.GetUser(ctx)
	groups := []*models.Group{}
	for _, g := range user.GoogleGroups {
		groups = append(groups, &models.Group{
			Name:  g.Name,
			Email: g.Email,
		})
	}

	return &models.UserInfo{
		Name:            user.Name,
		Email:           user.Email,
		Groups:          groups,
		LoginExpiration: user.Expiry,
	}, nil
}

// GoogleGroups is the resolver for the googleGroups field.
func (r *userInfoResolver) GoogleGroups(ctx context.Context, obj *models.UserInfo) ([]*models.Group, error) {
	return obj.Groups, nil
}

// AllGoogleGroups is the resolver for the allGoogleGroups field.
func (r *userInfoResolver) AllGoogleGroups(ctx context.Context, obj *models.UserInfo) ([]*models.Group, error) {
	user := auth.GetUser(ctx)
	groups := []*models.Group{}
	for _, g := range user.AllGoogleGroups {
		groups = append(groups, &models.Group{
			Name:  g.Name,
			Email: g.Email,
		})
	}
	return groups, nil
}

// AzureGroups is the resolver for the azureGroups field.
func (r *userInfoResolver) AzureGroups(ctx context.Context, obj *models.UserInfo) ([]*models.Group, error) {
	user := auth.GetUser(ctx)

	groups := []*models.Group{}
	for _, g := range user.AzureGroups {
		groups = append(groups, &models.Group{
			Name:  g.Name,
			Email: g.Email,
		})
	}

	return groups, nil
}

// GCPProjects is the resolver for the gcpProjects field.
func (r *userInfoResolver) GCPProjects(ctx context.Context, obj *models.UserInfo) ([]*models.GCPProject, error) {
	user := auth.GetUser(ctx)
	ret := []*models.GCPProject{}

	for _, grp := range user.GoogleGroups {
		proj, ok := r.gcpProjects.Get(grp.Email)
		if !ok {
			continue
		}

		ret = append(ret, &models.GCPProject{
			ID: proj,
			Group: &models.Group{
				Name:  grp.Name,
				Email: grp.Email,
			},
		})
	}

	return ret, nil
}

// NadaTokens is the resolver for the nadaTokens field.
func (r *userInfoResolver) NadaTokens(ctx context.Context, obj *models.UserInfo) ([]*models.NadaToken, error) {
	teams := teamNamesFromGroups(obj.Groups)
	return r.repo.GetNadaTokens(ctx, teams)
}

// Dataproducts is the resolver for the dataproducts field.
func (r *userInfoResolver) Dataproducts(ctx context.Context, obj *models.UserInfo) ([]*models.Dataproduct, error) {
	user := auth.GetUser(ctx)
	return r.repo.GetDataproductsByGroups(ctx, user.GoogleGroups.Emails())
}

// Accessable is the resolver for the accessable field.
func (r *userInfoResolver) Accessable(ctx context.Context, obj *models.UserInfo) ([]*models.Dataproduct, error) {
	user := auth.GetUser(ctx)
	return r.repo.GetDataproductsByUserAccess(ctx, "user:"+user.Email)
}

// Stories is the resolver for the stories field.
func (r *userInfoResolver) Stories(ctx context.Context, obj *models.UserInfo) ([]*models.GraphStory, error) {
	user := auth.GetUser(ctx)

	stories, err := r.repo.GetStoriesByGroups(ctx, user.GoogleGroups.Emails())
	if err != nil {
		return nil, err
	}

	gqlStories := make([]*models.GraphStory, len(stories))
	for i, s := range stories {
		gqlStories[i], err = storyFromDB(s)
		if err != nil {
			return nil, err
		}
	}
	return gqlStories, nil
}

// QuartoStories is the resolver for the quartoStories field.
func (r *userInfoResolver) QuartoStories(ctx context.Context, obj *models.UserInfo) ([]*models.QuartoStory, error) {
	user := auth.GetUser(ctx)
	return r.repo.GetQuartoStoriesByGroups(ctx, user.GoogleGroups.Emails())
}

// InsightProducts is the resolver for the insightProducts field.
func (r *userInfoResolver) InsightProducts(ctx context.Context, obj *models.UserInfo) ([]*models.InsightProduct, error) {
	user := auth.GetUser(ctx)
	return r.repo.GetInsightProductsByGroups(ctx, user.GoogleGroups.Emails())
}

// AccessRequests is the resolver for the accessRequests field.
func (r *userInfoResolver) AccessRequests(ctx context.Context, obj *models.UserInfo) ([]*models.AccessRequest, error) {
	user := auth.GetUser(ctx)

	groups := []string{"user:" + strings.ToLower(user.Email)}
	for _, g := range user.GoogleGroups {
		groups = append(groups, "group:"+strings.ToLower(g.Email))
	}

	return r.repo.ListAccessRequestsForOwner(ctx, groups)
}

// UserInfo returns generated.UserInfoResolver implementation.
func (r *Resolver) UserInfo() generated.UserInfoResolver { return &userInfoResolver{r} }

type userInfoResolver struct{ *Resolver }
