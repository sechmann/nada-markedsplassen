package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/graph/generated"
	"github.com/navikt/nada-backend/pkg/graph/models"
)

// AccessRequest is the resolver for the accessRequest field.
func (r *accessResolver) AccessRequest(ctx context.Context, obj *models.Access) (*models.AccessRequest, error) {
	if obj.AccessRequestID == nil {
		return nil, nil
	}
	return r.repo.GetAccessRequest(ctx, *obj.AccessRequestID)
}

// GrantAccessToDataset is the resolver for the grantAccessToDataset field.
func (r *mutationResolver) GrantAccessToDataset(ctx context.Context, input models.NewGrant) (*models.Access, error) {
	if input.Expires != nil && input.Expires.Before(time.Now()) {
		return nil, fmt.Errorf("Datoen tilgangen skal utløpe må være fram i tid.")
	}

	user := auth.GetUser(ctx)
	subj := user.Email
	if input.Subject != nil {
		subj = *input.Subject
	}
	ds, err := r.repo.GetDataset(ctx, input.DatasetID)
	if err != nil {
		return nil, err
	}

	dp, err := r.repo.GetDataproduct(ctx, ds.DataproductID)
	if err != nil {
		return nil, err
	}
	if err := isAllowedToGrantAccess(ctx, r.repo, dp, ds.ID, subj, user); err != nil {
		return nil, err
	}

	if ds.Pii == "sensitive" && subj == "all-users@nav.no" {
		return nil, fmt.Errorf("Datasett som inneholder personopplysninger kan ikke gjøres tilgjengelig for alle interne brukere (all-users@nav.no).")
	}

	bq, err := r.repo.GetBigqueryDatasource(ctx, ds.ID)
	if err != nil {
		return nil, err
	}

	subjType := models.SubjectTypeUser
	if input.SubjectType != nil {
		subjType = *input.SubjectType
	}

	subjWithType := subjType.String() + ":" + subj

	if err := r.accessMgr.Grant(ctx, bq.ProjectID, bq.Dataset, bq.Table, subjWithType); err != nil {
		return nil, err
	}

	return r.repo.GrantAccessToDataset(ctx, input.DatasetID, input.Expires, subjWithType, user.Email)
}

// RevokeAccessToDataset is the resolver for the revokeAccessToDataset field.
func (r *mutationResolver) RevokeAccessToDataset(ctx context.Context, id uuid.UUID) (bool, error) {
	access, err := r.repo.GetAccessToDataset(ctx, id)
	if err != nil {
		return false, err
	}

	ds, err := r.repo.GetDataset(ctx, access.DatasetID)
	if err != nil {
		return false, err
	}

	dp, err := r.repo.GetDataproduct(ctx, ds.DataproductID)
	if err != nil {
		return false, err
	}

	bq, err := r.repo.GetBigqueryDatasource(ctx, access.DatasetID)
	if err != nil {
		return false, err
	}

	user := auth.GetUser(ctx)
	if !user.GoogleGroups.Contains(dp.Owner.Group) && !strings.EqualFold("user:"+user.Email, access.Subject) {
		return false, ErrUnauthorized
	}

	if err := r.accessMgr.Revoke(ctx, bq.ProjectID, bq.Dataset, bq.Table, access.Subject); err != nil {
		return false, err
	}
	return true, r.repo.RevokeAccessToDataset(ctx, id)
}

// CreateAccessRequest is the resolver for the createAccessRequest field.
func (r *mutationResolver) CreateAccessRequest(ctx context.Context, input models.NewAccessRequest) (*models.AccessRequest, error) {
	user := auth.GetUser(ctx)
	subj := user.Email
	if input.Subject != nil {
		subj = *input.Subject
	}

	owner := "user:" + user.Email
	if input.Owner != nil {
		owner = "group:" + *input.Owner
	}

	subjType := models.SubjectTypeUser
	if input.SubjectType != nil {
		subjType = *input.SubjectType
	}

	subjWithType := subjType.String() + ":" + subj

	var pollyID uuid.NullUUID
	if input.Polly != nil {
		dbPolly, err := r.repo.CreatePollyDocumentation(ctx, *input.Polly)
		if err != nil {
			return nil, err
		}

		pollyID = uuid.NullUUID{UUID: dbPolly.ID, Valid: true}
	}

	ar, err := r.repo.CreateAccessRequestForDataset(ctx, input.DatasetID, pollyID, subjWithType, owner, input.Expires)
	if err != nil {
		return nil, err
	}
	r.SendNewAccessRequestSlackNotification(ctx, ar)
	return ar, nil
}

func (r *mutationResolver) SendNewAccessRequestSlackNotification(ctx context.Context, ar *models.AccessRequest) {
	ds, err := r.repo.GetDataset(ctx, ar.DatasetID)
	if err != nil {
		r.log.Warn("Access request created but failed to fetch dataset during sending slack notification", err)
		return
	}

	dp, err := r.repo.GetDataproduct(ctx, ds.DataproductID)
	if err != nil {
		r.log.Warn("Access request created but failed to fetch dataproduct during sending slack notification", err)
		return
	}

	if dp.Owner.TeamContact == nil || *dp.Owner.TeamContact == "" {
		r.log.Info("Access request created but skip slack message because teamcontact is empty")
		return
	}

	if IsEmail(*dp.Owner.TeamContact) {
		r.log.Info("Access request created but skip slack message because teamcontact is email")
		return
	}

	err = r.slack.NewAccessRequest(*dp.Owner.TeamContact, ar)
	if err != nil {
		r.log.Warn("Access request created, failed to send slack message", err)
	}
}

func IsEmail(contact string) bool {
	matched, err := regexp.Match(`(?:[a-z0-9!#$%&'*+/=?^_`+
		"`"+
		`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`+
		"`"+
		`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@
	  (?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])
	  ?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]
	  :(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`, []byte(contact))
	if err != nil {
		fmt.Println("Error with teamcontact")
		return false
	}
	return matched
}

// UpdateAccessRequest is the resolver for the updateAccessRequest field.
func (r *mutationResolver) UpdateAccessRequest(ctx context.Context, input models.UpdateAccessRequest) (*models.AccessRequest, error) {
	var pollyID uuid.NullUUID
	if input.Polly != nil {
		if input.Polly.ID != nil {
			// Keep existing polly
			pollyID = uuid.NullUUID{UUID: *input.Polly.ID, Valid: true}
		} else {
			dbPolly, err := r.repo.CreatePollyDocumentation(ctx, *input.Polly)
			if err != nil {
				return nil, err
			}
			pollyID = uuid.NullUUID{UUID: dbPolly.ID, Valid: true}
		}
	}

	return r.repo.UpdateAccessRequest(ctx, input.ID, pollyID, input.Owner, input.Expires)
}

// DeleteAccessRequest is the resolver for the deleteAccessRequest field.
func (r *mutationResolver) DeleteAccessRequest(ctx context.Context, id uuid.UUID) (bool, error) {
	accessRequest, err := r.repo.GetAccessRequest(ctx, id)
	if err != nil {
		return false, err
	}

	splits := strings.Split(accessRequest.Owner, ":")
	if len(splits) != 2 {
		return false, fmt.Errorf("%v is not a valid owner format (cannot split on :)", accessRequest.Owner)
	}
	owner := splits[1]

	if err := ensureOwner(ctx, owner); err != nil {
		return false, err
	}

	if err := r.repo.DeleteAccessRequest(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

// ApproveAccessRequest is the resolver for the approveAccessRequest field.
func (r *mutationResolver) ApproveAccessRequest(ctx context.Context, id uuid.UUID) (bool, error) {
	ar, err := r.repo.GetAccessRequest(ctx, id)
	if err != nil {
		return false, err
	}

	ds, err := r.repo.GetDataset(ctx, ar.DatasetID)
	if err != nil {
		return false, err
	}

	dp, err := r.repo.GetDataproduct(ctx, ds.DataproductID)
	if err != nil {
		return false, err
	}

	if err := ensureUserInGroup(ctx, dp.Owner.Group); err != nil {
		return false, err
	}

	user := auth.GetUser(ctx)
	if err := r.repo.ApproveAccessRequest(ctx, id, user.Email); err != nil {
		return false, err
	}

	return true, nil
}

// DenyAccessRequest is the resolver for the denyAccessRequest field.
func (r *mutationResolver) DenyAccessRequest(ctx context.Context, id uuid.UUID, reason *string) (bool, error) {
	ar, err := r.repo.GetAccessRequest(ctx, id)
	if err != nil {
		return false, err
	}

	ds, err := r.repo.GetDataset(ctx, ar.DatasetID)
	if err != nil {
		return false, err
	}

	dp, err := r.repo.GetDataproduct(ctx, ds.DataproductID)
	if err != nil {
		return false, err
	}

	if err := ensureUserInGroup(ctx, dp.Owner.Group); err != nil {
		return false, err
	}

	user := auth.GetUser(ctx)
	if err := r.repo.DenyAccessRequest(ctx, id, user.Email, reason); err != nil {
		return false, err
	}

	return true, nil
}

// AccessRequest is the resolver for the accessRequest field.
func (r *queryResolver) AccessRequest(ctx context.Context, id uuid.UUID) (*models.AccessRequest, error) {
	return r.repo.GetAccessRequest(ctx, id)
}

// Access returns generated.AccessResolver implementation.
func (r *Resolver) Access() generated.AccessResolver { return &accessResolver{r} }

type accessResolver struct{ *Resolver }
