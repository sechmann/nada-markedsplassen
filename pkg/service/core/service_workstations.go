package core

import (
	"context"
	"fmt"

	sluglib "github.com/gosimple/slug"
	"github.com/navikt/nada-backend/pkg/errs"
	"github.com/navikt/nada-backend/pkg/service"
)

var _ service.WorkstationsService = (*workstationService)(nil)

type workstationService struct {
	projectID string

	workstationAPI    service.WorkstationsAPI
	serviceAccountAPI service.ServiceAccountAPI
}

func (s *workstationService) StartWorkstation(ctx context.Context, user *service.User) error {
	const op errs.Op = "workstationService.StartWorkstation"

	slug := NormalizedEmail(user.Email)

	err := s.workstationAPI.StartWorkstation(ctx, &service.WorkstationStartOpts{
		Slug:       slug,
		ConfigName: slug,
	})
	if err != nil {
		return errs.E(op, err)
	}

	return nil
}

func (s *workstationService) StopWorkstation(ctx context.Context, user *service.User) error {
	const op errs.Op = "workstationService.StopWorkstation"

	slug := NormalizedEmail(user.Email)

	err := s.workstationAPI.StopWorkstation(ctx, &service.WorkstationStopOpts{
		Slug:       slug,
		ConfigName: slug,
	})
	if err != nil {
		return errs.E(op, err)
	}

	return nil
}

func (s *workstationService) EnsureWorkstation(ctx context.Context, user *service.User, input *service.WorkstationInput) (*service.WorkstationOutput, error) {
	const op errs.Op = "workstationService.EnsureWorkstation"

	slug := NormalizedEmail(user.Email)

	sa, err := s.serviceAccountAPI.EnsureServiceAccount(ctx, &service.ServiceAccountRequest{
		ProjectID:   s.projectID,
		AccountID:   slug,
		DisplayName: slug,
		Description: fmt.Sprintf("Workstation service account for %s (%s)", user.Name, user.Email),
	})
	if err != nil {
		return nil, errs.E(op, fmt.Errorf("ensuring service account for %s: %w", user.Email, err))
	}

	// FIXME: Need to grant the correct roles for the user to able to access the created workstation

	c, w, err := s.workstationAPI.EnsureWorkstationWithConfig(ctx, &service.EnsureWorkstationOpts{
		Workstation: service.WorkstationOpts{
			Slug:        slug,
			ConfigName:  slug,
			DisplayName: displayName(user),
			Labels:      service.DefaultWorkstationLabels(slug),
		},
		Config: service.WorkstationConfigOpts{
			Slug:                slug,
			DisplayName:         displayName(user),
			ServiceAccountEmail: sa.Email,
			SubjectEmail:        user.Email,
			Labels:              service.DefaultWorkstationLabels(slug),
			MachineType:         input.MachineType,
			ContainerImage:      input.ContainerImage,
		},
	})
	if err != nil {
		return nil, errs.E(op, fmt.Errorf("ensuring workstation for %s: %w", user.Email, err))
	}

	return &service.WorkstationOutput{
		Slug:        w.Slug,
		DisplayName: w.DisplayName,
		Reconciling: w.Reconciling,
		CreateTime:  w.CreateTime,
		UpdateTime:  w.UpdateTime,
		StartTime:   w.StartTime,
		State:       w.State,
		Config: &service.WorkstationConfigOutput{
			CreateTime:     c.CreateTime,
			UpdateTime:     c.UpdateTime,
			IdleTimeout:    c.IdleTimeout,
			RunningTimeout: c.RunningTimeout,
			MachineType:    c.MachineType,
			Image:          c.Image,
			Env:            c.Env,
		},
	}, nil
}

func (s *workstationService) GetWorkstation(ctx context.Context, user *service.User) (*service.WorkstationOutput, error) {
	const op errs.Op = "workstationService.GetWorkstation"

	slug := NormalizedEmail(user.Email)

	c, err := s.workstationAPI.GetWorkstationConfig(ctx, &service.WorkstationConfigGetOpts{
		Slug: slug,
	})
	if err != nil {
		return nil, errs.E(op, err)
	}

	w, err := s.workstationAPI.GetWorkstation(ctx, &service.WorkstationGetOpts{
		Slug:       slug,
		ConfigName: slug,
	})
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &service.WorkstationOutput{
		Slug:        w.Slug,
		DisplayName: w.DisplayName,
		Reconciling: w.Reconciling,
		CreateTime:  w.CreateTime,
		UpdateTime:  w.UpdateTime,
		StartTime:   w.StartTime,
		State:       w.State,
		Config: &service.WorkstationConfigOutput{
			CreateTime:     c.CreateTime,
			UpdateTime:     c.UpdateTime,
			IdleTimeout:    c.IdleTimeout,
			RunningTimeout: c.RunningTimeout,
			MachineType:    c.MachineType,
			Image:          c.Image,
			Env:            c.Env,
		},
	}, nil
}

func (s *workstationService) DeleteWorkstation(ctx context.Context, user *service.User) error {
	const op errs.Op = "workstationService.DeleteWorkstation"

	err := s.workstationAPI.DeleteWorkstationConfig(ctx, &service.WorkstationConfigDeleteOpts{
		Slug: NormalizedEmail(user.Email),
	})
	if err != nil {
		return errs.E(op, err)
	}

	return nil
}

func NormalizedEmail(email string) string {
	sluglib.MaxLength = 63
	sluglib.CustomSub = map[string]string{
		"_": "-",
		"@": "-at-",
	}

	return sluglib.Make(email)
}

func displayName(user *service.User) string {
	return fmt.Sprintf("%s (%s)", user.Name, user.Email)
}

func NewWorkstationService(projectID string, s service.ServiceAccountAPI, w service.WorkstationsAPI) *workstationService {
	return &workstationService{
		projectID:         projectID,
		workstationAPI:    w,
		serviceAccountAPI: s,
	}
}
