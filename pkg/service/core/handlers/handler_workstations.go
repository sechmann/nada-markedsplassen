package handlers

import (
	"context"
	"net/http"

	"github.com/navikt/nada-backend/pkg/auth"
	"github.com/navikt/nada-backend/pkg/errs"
	"github.com/navikt/nada-backend/pkg/service"
)

type WorkstationsHandler struct {
	service service.WorkstationsService
}

func (h *WorkstationsHandler) CreateWorkstation(ctx context.Context, r *http.Request, input *service.WorkstationInput) (*service.WorkstationOutput, error) {
	const op errs.Op = "WorkstationsHandler.CreateWorkstation"

	user := auth.GetUser(ctx)
	if user == nil {
		return nil, errs.E(errs.Unauthenticated, op, errs.Str("no user in context"))
	}

	workstation, err := h.service.EnsureWorkstation(ctx, user, input)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return workstation, nil
}

func (h *WorkstationsHandler) GetWorkstation(ctx context.Context, _ *http.Request, _ any) (*service.WorkstationOutput, error) {
	const op errs.Op = "WorkstationsHandler.GetWorkstation"

	user := auth.GetUser(ctx)
	if user == nil {
		return nil, errs.E(errs.Unauthenticated, op, errs.Str("no user in context"))
	}

	workstation, err := h.service.GetWorkstation(ctx, user)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return workstation, nil
}

func NewWorkstationsHandler(service service.WorkstationsService) *WorkstationsHandler {
	return &WorkstationsHandler{
		service: service,
	}
}
